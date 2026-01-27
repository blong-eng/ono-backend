package tutorials

func GetBuilderSteps() []map[string]interface{} {
	return []map[string]interface{}{
		{"target_key": "location_key",
		"title": "Destination",
		"instruction": "Where are you heading? Enter the city and state."},

		{"target_key": "date_key",
		"title": "Date Selection",
		"instruction": "Pick the night you want to head out."},

		{"target_key": "party_size_key",
		"title": "Host Party", "instruction":
		"How many people are arriving with you?"},

		{"target_key": "guest_invite_key",
		"title": "Guest List",
		"instruction": "How many total guests do you want to join the night?"},

		{"target_key": "cohosts_key",
		"title": "Co-Hosting",
		"instruction": "Select connections to help you manage the night."},

		{"target_key": "criteria_key",
		"title": "Motivations",
		"instruction": "What's the goal for the night? Exploring, meeting people, or just fun?"},

		{"target_key": "participants_key",
		"title": "Participant Type",
		"instruction": "Define if this is for singles, couples, or everyone."},

		{"target_key": "specific_invites_key",
		"title": "Specific Invites",
		"instruction": "Invite your existing connections directly to this plan."},

		{"target_key": "age_range_key",
		"title": "Preferred Age",
		"instruction": "Set the age range for potential guests joining via the app."},

		{"target_key": "financial_key",
		"title": "Financial Vibe",
		"instruction": "Set expectations for how the tab will be handled."},

		{"target_key": "magic_button_key",
		"title": "Instant Magic",
		"instruction": "Take a Chance! Let the AI build a 3-stop night for you instantly based on your archetypes!"},

		{"target_key": "plan_manual_key",
		"title": "Plan Manually",
		"instruction": "Prefer total control? Choose your own venues step-by-step."},
	}
}