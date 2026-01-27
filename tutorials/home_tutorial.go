package tutorials

// GetHomeSteps returns the full mapping of items for the interactive home tutorial.
func GetHomeSteps() []map[string]interface{} { // Changed from string to interface{}
    return []map[string]interface{}{
		{
			"target_key":  "Create",
			"title":       "Create",
			"instruction": "This is your starting point. Tap the center icon to build a custom itinerary or use AI to plan your night. Now, spin the wheel to see more!",
		},
		{
			"target_key":  "My Nights",
			"title":       "My Nights",
			"instruction": "All your upcoming plans and past adventures are stored here for quick access.",
		},
		{
			"target_key":  "Board",
			"title":       "Community Board",
			"instruction": "Post updates and see what's happening in your local social scene.",
		},
		{
			"target_key":  "Alerts",
			"title":       "Alerts",
			"instruction": "Check here for invites and real-time updates on your active plans.",
		},
		{
			"target_key":  "Connections",
			"title":       "Connections",
			"instruction": "Find friends nearby and manage your core group of co-hosts.",
		},
		{
			"target_key":  "Profile",
			"title":       "Your Profile",
			"instruction": "Update your social vibe and badges to help our AI find better matches for you.",
		},
		{
			"target_key":  "Explore",
			"title":       "Explore",
			"instruction": "Discover Nights out that other local hosts have planned.",
		},
		{
			"target_key":  "settings_key",
			"title":       "Menu & Support",
			"instruction": "Access Contact Us, Logout, and Account Deletion here.",
		},
	}
}