package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	// ✅ Ensure this matches your directory structure
	"ono-backend/tutorials"
)

var client *firestore.Client
var ctx = context.Background()

// RevenueCatPayload defines the structure for incoming payment notifications
type RevenueCatPayload struct {
	Event struct {
		AppUserID      string `json:"app_user_id"`
		ProductID      string `json:"product_id"`
		Type           string `json:"type"`
		ExpirationAtMS int64  `json:"expiration_at_ms"`
	} `json:"event"`
}

func main() {
	// 1. Initialize Firestore
	projectID := "one-night-out-eb2ed"
	var sa option.ClientOption

	// Railway Environment Variable Check
	creds := os.Getenv("FIREBASE_CREDENTIALS")
	if creds != "" {
		sa = option.WithCredentialsJSON([]byte(creds))
		log.Println("🔑 Using Firebase Credentials from Environment Variables")
	} else {
		sa = option.WithServiceAccountFile("one-night-out-eb2ed-firebase-adminsdk-fbsvc-d62d276844.json")
		log.Println("📁 Using local Firebase JSON file")
	}

	var err error
	client, err = firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		log.Fatalf("❌ Firestore Client Error: %v", err)
	}
	defer client.Close()

	log.Println("🚀 ONO Backend connected to Firestore (Cloud Mode)")

	r := gin.Default()

	// 🟢 PERMANENT FIX: Handle path mismatches automatically
	// This prevents 404 errors if Flutter adds a slash or slightly differs in path formatting.
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	// --- SECTION 1: TUTORIAL ENDPOINTS ---

	r.GET("/tutorials/:screen", func(c *gin.Context) {
		screen := c.Param("screen")
		userID := c.Query("userId")

		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
			return
		}

		doc, err := client.Collection("users").Doc(userID).
			Collection("tutorials").Doc(screen).Get(ctx)

		if err == nil && doc.Exists() {
			data := doc.Data()
			if finished, ok := data["finished"].(bool); ok && finished {
				c.JSON(http.StatusOK, []interface{}{})
				return
			}
		}

		steps := getTutorialSteps(screen)
		c.JSON(http.StatusOK, steps)
	})

	r.POST("/tutorials/complete", func(c *gin.Context) {
		var req struct {
			UserID string `json:"userId"`
			Screen string `json:"screen"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		_, err := client.Collection("users").Doc(req.UserID).
			Collection("tutorials").Doc(req.Screen).
			Set(ctx, map[string]interface{}{
				"finished":  true,
				"updatedAt": firestore.ServerTimestamp,
			}, firestore.MergeAll)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save progress"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// --- SECTION 2: REVENUECAT WEBHOOK ---
	r.POST("/webhooks/revenuecat", func(c *gin.Context) {
		var payload RevenueCatPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}

		log.Printf("💳 Webhook Received: Type=%s, User=%s, Product=%s",
			payload.Event.Type, payload.Event.AppUserID, payload.Event.ProductID)

		// Handle successful purchases or renewals
		if payload.Event.Type == "INITIAL_PURCHASE" || payload.Event.Type == "RENEWAL" {
			expiry := time.Unix(0, payload.Event.ExpirationAtMS*int64(time.Millisecond))

			// 🟢 PERMANENT FIX: Match IDs exactly to your RevenueCat "Published" list
			tier := "premium"
			if payload.Event.ProductID == "premium_tier_25:pt25" {
				tier = "local_gem"
			}

			// 🟢 PERMANENT FIX: Use 'Set' with 'MergeAll' instead of 'Update'
			// This ensures the document is created if it doesn't exist, preventing "NOT_FOUND" errors.
			_, err := client.Collection("businesses").Doc(payload.Event.AppUserID).Set(ctx, map[string]interface{}{
				"subscriptionTier":   tier,
				"subscriptionExpiry": expiry,
				"isRegistered":       true,
			}, firestore.MergeAll)

			if err != nil {
				log.Printf("❌ Webhook Database Error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database update failed"})
				return
			}
		}

		// Handle expired subscriptions
		if payload.Event.Type == "EXPIRATION" {
			_, err := client.Collection("businesses").Doc(payload.Event.AppUserID).Update(ctx, []firestore.Update{
				{Path: "subscriptionTier", Value: "free"},
				{Path: "subscriptionExpiry", Value: nil},
			})
			if err != nil {
				log.Printf("❌ Expiration Update Error: %v", err)
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": "processed"})
	})

	// --- SECTION 3: TRANSACTIONAL SLOT RESERVATION ---
	r.POST("/subscriptions/reserve-slot", func(c *gin.Context) {
		var req struct {
			BusinessID string `json:"businessId"`
			City       string `json:"city"`
			State      string `json:"state"`
			Tag        string `json:"tag"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Perform check inside a transaction to ensure accuracy
		err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
			query := client.Collection("businesses").
				Where("city", "==", req.City).
				Where("state", "==", req.State).
				Where("subscriptionTier", "in", []string{"premium", "local_gem"})

			docs, err := tx.Documents(query).GetAll()
			if err != nil {
				return err
			}

			count := 0
			for _, doc := range docs {
				data := doc.Data()
				if tags, ok := data["searchTags"].([]interface{}); ok {
					for _, t := range tags {
						if t.(string) == req.Tag {
							count++
						}
					}
				}
			}

			if count >= 10 {
				return fmt.Errorf("SLOT_FULL")
			}
			return nil
		})

		if err != nil {
			if err.Error() == "SLOT_FULL" {
				// Use 409 Conflict so Flutter knows exactly why it failed
				c.JSON(http.StatusConflict, gin.H{"error": "Niche capacity reached"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "available"})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Go backend is active"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("📡 Server starting on port %s", port)
	r.Run("0.0.0.0:" + port)
}

func getTutorialSteps(screen string) []map[string]interface{} {
	switch screen {
	case "home": return tutorials.GetHomeSteps()
	case "builder": return tutorials.GetBuilderSteps()
	case "business": return tutorials.GetBusinessSteps()
	case "profile": return tutorials.GetProfileSteps()
	case "explore": return tutorials.GetExploreSteps()
	case "wallet": return tutorials.GetWalletSteps()
	case "planner": return tutorials.GetPlannerSteps()
	case "message_board": return tutorials.GetBoardSteps()
	default: return []map[string]interface{}{}
	}
}