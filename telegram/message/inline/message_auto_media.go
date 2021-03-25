package inline

import (
	"github.com/gotd/td/telegram/message/entity"
	"github.com/gotd/td/telegram/message/markup"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"
)

// MessageMediaAutoBuilder is a builder of inline result text message.
type MessageMediaAutoBuilder struct {
	message *tg.InputBotInlineMessageMediaAuto
	option  styling.StyledTextOption
	options []styling.StyledTextOption
}

func (b *MessageMediaAutoBuilder) apply() (tg.InputBotInlineMessageClass, error) {
	tb := entity.Builder{}
	if err := styling.Perform(&tb, b.option, b.options...); err != nil {
		return nil, err
	}
	msg, entities := tb.Complete()
	r := *b.message

	r.Message = msg
	r.Entities = entities
	return &r, nil
}

// MediaAuto creates new message text option builder.
func MediaAuto(msg string) *MessageMediaAutoBuilder {
	return MediaAutoStyled(styling.Plain(msg))
}

// MediaAutoStyled creates new message text option builder.
func MediaAutoStyled(text styling.StyledTextOption, texts ...styling.StyledTextOption) *MessageMediaAutoBuilder {
	return &MessageMediaAutoBuilder{
		message: &tg.InputBotInlineMessageMediaAuto{},
		option:  text,
		options: texts,
	}
}

// Markup sets reply markup for sending bot buttons.
// NB: markup will not be used, if you send multiple media attachments.
func (b *MessageMediaAutoBuilder) Markup(m tg.ReplyMarkupClass) *MessageMediaAutoBuilder {
	b.message.ReplyMarkup = m
	return b
}

// Row sets single row keyboard markup  for sending bot buttons.
// NB: markup will not be used, if you send multiple media attachments.
func (b *MessageMediaAutoBuilder) Row(
	button tg.KeyboardButtonClass,
	buttons ...tg.KeyboardButtonClass,
) *MessageMediaAutoBuilder {
	return b.Markup(markup.InlineRow(button, buttons...))
}
