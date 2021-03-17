package views

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/slack-go/slack"
)

func TestGreetingMessage(t *testing.T) {
	tests := []struct {
		name string
		user string
		want []slack.Block
	}{
		{
			name: "User name is added",
			user: "David",
			want: []slack.Block{
				&slack.SectionBlock{
					Type: slack.MBTSection,
					Text: &slack.TextBlockObject{
						Type: "mrkdwn",
						Text: "Hi David :wave:",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Block section 0
			if diff := deep.Equal(GreetingMessage(tt.user)[0], tt.want[0]); diff != nil {
				t.Error(diff)
			}
		})
	}
}
