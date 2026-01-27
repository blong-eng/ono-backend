package tutorials

func GetExploreSteps() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"target_key":  "magic_search",
			"title":       "Local Discovery",
			"instruction": "Tap here to see all active nights in your city. We prioritize results based on your social archetype so you find the perfect vibe immediately.",
		},
		{
			"target_key":  "location_date",
			"title":       "Change Your View",
			"instruction": "Traveling? Update the city, state, or search for specific dates to plan your future adventures.",
		},
		{
			"target_key":  "vibe_selector",
			"title":       "Specific Vibes",
			"instruction": "Narrow your results to specific categories like 'Foodie Adventures' or 'Hidden Gems' with one tap.",
		},
		{
			"target_key":  "advanced_filters",
			"title":       "Deep Refinement",
			"instruction": "Expand this to filter by distance, participant types (Singles/Couples), and financial obligations like who pays the bill.",
		},
	}
}