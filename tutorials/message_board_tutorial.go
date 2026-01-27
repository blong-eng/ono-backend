package tutorials

func GetBoardSteps() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"target_key":  "tribe_tabs",
			"title":       "Your Tribes",
			"instruction": "We've created separate boards for your Primary and Secondary archetypes. Talk to people who share your social energy or join 'The Mix' for everyone!",
		},
		{
			"target_key":  "starter_card",
			"title":       "AI Conversation Starters",
			"instruction": "Not sure what to say? Tap these AI-generated prompts to kickstart a discussion or check if someone has already answered!",
		},
		{
			"target_key":  "topic_filters",
			"title":       "Filter by Topic",
			"instruction": "Narrow down the conversation to Venues, Meetups, or general Discussion using these chips.",
		},
		{
			"target_key":  "new_post_fab",
			"title":       "Share Your Thoughts",
			"instruction": "Start your own conversation. You can post to either of your tribe boards or the general community.",
		},
	}
}