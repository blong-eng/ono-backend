package tutorials

// GetWalletSteps returns the instructional sequence for the Ticket Wallet.
func GetWalletSteps() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"target_key":  "wallet_tabs",
			"title":       "Your Ticket Wallet",
			"instruction": "Track your upcoming plans here. Once a night out has ended, it will automatically move to your 'Past History' tab.",
		},
		{
			"target_key":  "rate_night_button",
			"title":       "Archetype Evolution",
			"instruction": "This is the most important step! Rate your nights out so our AI can learn which venues matched your energy and refine your social archetype.",
		},
	}
}