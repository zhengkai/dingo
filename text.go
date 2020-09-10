package dingo

type text struct {
	Msgtype string  `json:"msgtype"`
	Text    textSub `json:"text"`
}

type textSub struct {
	Content string `json:"content"`
}

// Text ...
func (b *Bot) Text(msg string) (err error) {

	d := &text{
		Msgtype: `text`,
		Text: textSub{
			Content: msg,
		},
	}
	return b.Send(d)
}
