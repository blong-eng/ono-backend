 package services

 import (
 	"context"
 	"time"
 	"cloud.google.com/go/firestore"
 )

 // MarkTutorialComplete writes the "true" flag to Firestore.
 // This is what finally creates the 'user_tutorials' collection.
 func MarkTutorialComplete(ctx context.Context, client *firestore.Client, userId string, screenKey string) error {
 	// Set the flag for the specific screen (e.g., 'home': true)
 	_, err := client.Collection("user_tutorials").Doc(userId).Set(ctx, map[string]interface{}{
 		screenKey:    true,
 		"last_updated": time.Now(),
 	}, firestore.MergeAll) // MergeAll ensures we don't overwrite other tutorial flags
 	return err
 }

 // IsTutorialComplete checks if the user has already finished this tutorial.
 func IsTutorialComplete(ctx context.Context, client *firestore.Client, userId string, screenKey string) bool {
 	doc, err := client.Collection("user_tutorials").Doc(userId).Get(ctx)
 	if err != nil {
 		return false // If error or no doc, assume they need the tutorial
 	}

 	data := doc.Data()
 	if complete, ok := data[screenKey].(bool); ok {
 		return complete
 	}
 	return false
 }