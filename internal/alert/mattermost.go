package alert

import "github.com/slack-go/slack"

func toMattermost(URL string, color string, log map[string]string) {
	msg := generateAttachments(color, log)
	slack.PostWebhook(URL, &slack.WebhookMessage{Attachments: msg}) // nolint:errcheck
}
