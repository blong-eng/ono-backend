package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	// ✅ Import your local tutorials package
	"ono-backend/tutorials"
)

var client *firestore.Client
var ctx = context.Background()

func main() {
	// 1. Initialize Firestore
	projectID := "one-night-out-eb2ed"

	// Access the key you added to .gitignore
sa := option.WithServiceAccountFile("one-night-out-eb2ed-firebase-adminsdk-fbsvc-d62d276844.json")
	var err error
	client, err = firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		log.Fatalf("❌ Firestore Client Error: %v", err)
	}
	defer client.Close()

	log.Println("🚀 ONO Backend connected to Firestore (Cloud Mode)")

	r := gin.Default()

	// 2. Endpoint: Get Tutorial Steps
	r.GET("/tutorials/:screen", func(c *gin.Context) {
		screen := c.Param("screen")
		userID := c.Query("userId")

		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
			return
		}

		// Check Firestore for existing completion record
		doc, err := client.Collection("users").Doc(userID).
			Collection("tutorials").Doc(screen).Get(ctx)

		if err == nil && doc.Exists() {
			data := doc.Data()
			if finished, ok := data["finished"].(bool); ok && finished {
				log.Printf("⏩ User %s already finished %s. Returning empty.", userID, screen)
				c.JSON(http.StatusOK, []interface{}{})
				return
			}
		}

		// ✅ Updated to handle the interface{} return type
		steps := getTutorialSteps(screen)
		c.JSON(http.StatusOK, steps)
	})

	// 3. Endpoint: Mark Tutorial Complete
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
			})

		if err != nil {
			log.Printf("❌ Firestore Write Error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save progress"})
			return
		}

		log.Printf("✅ Tutorial '%s' marked finished for user %s", req.Screen, req.UserID)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Go backend is active"})
	})

	r.Run("0.0.0.0:8080")
}

// ✅ Updated return type to []map[string]interface{} to resolve compile errors
func getTutorialSteps(screen string) []map[string]interface{} {
	switch screen {
	case "home":
		return tutorials.GetHomeSteps()
	case "builder":
		return tutorials.GetBuilderSteps()
	case "business":
		return tutorials.GetBusinessSteps()
	case "profile":
		return tutorials.GetProfileSteps() // ✅ Added Profile
	case "explore":
		return tutorials.GetExploreSteps() // ✅ Added Explore
	case "wallet":
		return tutorials.GetWalletSteps() // ✅ Added Wallet
	case "planner":
		return tutorials.GetPlannerSteps() // ✅ Added Planner
	case "message_board":
		return tutorials.GetBoardSteps() // ✅ Added Message Board
	default:
		return []map[string]interface{}{}
	}
}