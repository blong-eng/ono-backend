package tutorials

func GetProfileSteps() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"target_key":  "avatar_display",
			"title":       "Your Bifurcated Avatar",
			"instruction": "This is your social signature. The left side represents your Primary Archetype, and the right side represents your Secondary. Colors and emojis reflect your unique social energy at a glance!",
		},
		{
			"target_key":  "header_section",
			"title":       "Your Identity",
			"instruction": "This area displays your basic info.",
		},
		{
			"target_key":  "reputation_section",
			"title":       "Social Reputation",
			"instruction": "Your activity and positive interactions build your reputation. High-reputation users are often prioritized in the 'Explore' results.",
		},
		{
			"target_key":  "archetype_section",
			"title":       "Social Archetypes",
			"instruction": "Your Archetypes are the 'engine' of the app. They determine which venues the AI generator picks for you and how you are matched with others.",
		},

        {
                "target_key":  "badges_section",
                "title":       "Your Achievements",
                "instruction": "Badges represent your milestones and verified social traits. They add unique flavor to your profile and contribute to your overall social reputation!",
        },

		{
			"target_key":  "vibe_section",
			"title":       "AI Social Vibe",
			"instruction": "This is an AI-crafted summary of your social personality. It evolves dynamically as you update your badges or retake the archetype quiz.",
		},
		{
			"target_key":  "edit_button",
			"title":       "Customize & Update",
			"instruction": "Need a change? Tap here to update your vitals, select new badges, or retake the quiz to refresh your vibe and avatar!",
		},
	}
}