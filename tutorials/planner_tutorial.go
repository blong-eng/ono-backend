package tutorials

// GetPlannerSteps returns the instructional sequence for the manual Night Out Planner.
func GetPlannerSteps() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"target_key":  "distance_control",
			"title":       "Search Radius",
			"instruction": "Toggle the switch and use the slider to adjust how far you want to search from your current or previous location.",
		},
		{
			"target_key":  "search_bar",
			"title":       "Specific Search",
			"instruction": "Have a specific place in mind? Looking for a specific activity? Use the search bar to find any venue by name or interest.",
		},
		{
			"target_key":  "category_filters",
			"title":       "Refining Choices",
			"instruction": "Narrow down results with these category chips to find exactly the type of vibe you're looking for.",
		},
		{
			"target_key":  "nav_buttons",
			"title":       "The Flow",
			"instruction": "Use 'Continue' to move through the three phases: Meeting, Dining, and Activities.",
		},
	}
}