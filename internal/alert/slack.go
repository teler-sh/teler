package alert

import "github.com/slack-go/slack"

func toSlack(token string, channel string, color string, log map[string]string, webhooked bool) {
	msg := generateAttachments(color, log)

	if webhooked {
		slack.PostWebhook(token, &slack.WebhookMessage{Attachments: msg}) // nolint:errcheck
	} else {
		// TODO: Displays an error if it does not exceed the rate-limit
		api := slack.New(token)
		api.PostMessage( // nolint:errcheck
			channel,
			slack.MsgOptionAttachments(msg...),
			slack.MsgOptionAsUser(true),
		)
	}
}
