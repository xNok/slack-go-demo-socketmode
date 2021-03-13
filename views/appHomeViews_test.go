package views

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-test/deep"
	"github.com/slack-go/slack"
)

var default_blocks []slack.Block = []slack.Block{
	&slack.SectionBlock{
		Type: slack.MBTSection,
		Text: &slack.TextBlockObject{
			Type:     "mrkdwn",
			Text:     "*Welcome Back!* \nThis is a home for Stickers app. You can add small notes here!",
			Emoji:    false,
			Verbatim: false,
		},
		Accessory: &slack.Accessory{
			ButtonElement: &slack.ButtonBlockElement{
				Type: slack.METButton,
				Text: &slack.TextBlockObject{
					Type:     "plain_text",
					Text:     "Add a Stickie",
					Emoji:    false,
					Verbatim: false,
				},
				ActionID: AddStockieNoteActionID,
			},
		},
	},
	slack.NewDividerBlock(),
}

func TestAppHomeTabView(t *testing.T) {
	tests := []struct {
		name string
		want slack.HomeTabViewRequest
	}{
		{
			name: "Empty Home Page",
			want: slack.HomeTabViewRequest{
				Type: slack.VTHomeTab,
				Blocks: slack.Blocks{
					BlockSet: default_blocks,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppHomeTabView(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppHomeTabView() = %v, want %v", got, tt.want)
				spew.Dump(got)
			}
		})
	}
}

func TestCreateStickieNoteModal(t *testing.T) {
	tests := []struct {
		name string
		want slack.ModalViewRequest
	}{
		{
			name: "Simple Modal",
			want: slack.ModalViewRequest{
				Type: slack.VTModal,
				Title: &slack.TextBlockObject{
					Type:     "plain_text",
					Text:     "Create a stickie note",
					Emoji:    true,
					Verbatim: false,
				},
				Blocks: slack.Blocks{
					BlockSet: []slack.Block{
						&slack.InputBlock{
							Type:    slack.MBTInput,
							BlockID: ModalDescriptionBlockID,
							Label: &slack.TextBlockObject{
								Type:     "plain_text",
								Text:     "Note",
								Emoji:    false,
								Verbatim: false,
							},
							Element: &slack.PlainTextInputBlockElement{
								Type:     slack.METPlainTextInput,
								ActionID: ModalDescriptionActionID,
								Placeholder: &slack.TextBlockObject{
									Type: "plain_text",
									Text: "Take a note... ",
								},
								Multiline: true,
							},
						},
						&slack.InputBlock{
							Type:    slack.MBTInput,
							BlockID: ModalColorBlockID,
							Label: &slack.TextBlockObject{
								Type: "plain_text",
								Text: "Color",
							},
							Element: &slack.SelectBlockElement{
								Type:     slack.OptTypeStatic,
								ActionID: ModalColorActionID,
								Options: []*slack.OptionBlockObject{
									{
										Text: &slack.TextBlockObject{
											Type: "plain_text",
											Text: "yellow",
										},
										Value: "yellow",
									},
									{
										Text: &slack.TextBlockObject{
											Type: "plain_text",
											Text: "blue",
										},
										Value: "blue",
									},
								},
							},
						},
					},
				},
				Submit: &slack.TextBlockObject{
					Type:  "plain_text",
					Text:  "Create",
					Emoji: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := CreateStickieNoteModal(); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CreateStickieNoteModal() = %v, want %v", got, tt.want)
			// 	spew.Dump(got)
			// }
			if diff := deep.Equal(CreateStickieNoteModal(), tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

var note_blocks []slack.Block = []slack.Block{
	&slack.ContextBlock{
		Type: slack.MBTContext,
		ContextElements: slack.ContextElements{
			Elements: []slack.MixedElement{
				&slack.TextBlockObject{
					Type: "mrkdwn",
					Text: "Today",
				},
			},
		},
	},
	&slack.SectionBlock{
		Type: slack.MBTSection,
		Text: &slack.TextBlockObject{
			Type: "mrkdwn",
			Text: "test",
		},
		Accessory: &slack.Accessory{
			ImageElement: &slack.ImageBlockElement{
				Type:     slack.METImage,
				ImageURL: "https://cdn.glitch.com/0d5619da-dfb3-451b-9255-5560cd0da50b%2Fstickie_blue.png",
				AltText:  "blue stickie note",
			},
		},
	},
	slack.NewDividerBlock(),
}

func TestAppHomeCreateStickieNote(t *testing.T) {
	type args struct {
		note StickieNote
	}
	tests := []struct {
		name string
		note StickieNote
		want slack.HomeTabViewRequest
	}{
		{
			name: "1 Note",
			note: StickieNote{
				Description: "test",
				Color:       "blue",
				Timestamp:   "Today",
			},
			want: slack.HomeTabViewRequest{
				Type: slack.VTHomeTab,
				Blocks: slack.Blocks{
					BlockSet: append(default_blocks, note_blocks...),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := AppHomeCreateStickieNote(tt.note); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("AppHomeCreateStickieNote() = %v, want %v", got, tt.want)
			// }
			if diff := deep.Equal(AppHomeCreateStickieNote(tt.note), tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
