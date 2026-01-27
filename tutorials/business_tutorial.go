package tutorials

func GetBusinessSteps() []map[string]interface{} { // Changed from string to interface{}
    return []map[string]interface{}{
		{
			"target_key":  "explore_key",
			"title":       "Discover Places",
			"instruction": "Tap here to find the best spots in your city.",
		},
		{
			"target_key":  "create_key",
			"title":       "Plan a Night",
			"instruction": "Ready to go out? Start your journey here.",
		},
	}
}