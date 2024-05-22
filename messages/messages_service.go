package messages

import "go_chat/messages/types"

func SendMessage(input types.SendMessageInput) (MessageModel, error) {
	message := NewMessage()

	return message, nil
}
