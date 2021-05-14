package views

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/slack-go/slack"
)

func TestLaunchRocketAnnoncement(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   []slack.Block
	}{
		{
			name:   "Count down is correct",
			number: 3,
			want: []slack.Block{
				&slack.SectionBlock{
					Type: slack.MBTSection,
					Fields: []*slack.TextBlockObject{
						{
							Type: "mrkdwn",
							Text: "*Rocket:*\nFalcon 9",
						},
						{
							Type: "mrkdwn",
							Text: "*When:*\n3s count down",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			blocks := LaunchRocketAnnoncement(tt.number)

			if diff := deep.Equal(blocks[1], tt.want[0]); diff != nil {
				t.Error(diff)
			}
		})
	}
}
