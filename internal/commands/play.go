package commands

import "regexp"

var urlPattern = regexp.MustCompile(
	"^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?",
)
