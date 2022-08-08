package qiscusRequest

type CustomChannelSendMessageText struct {
	IdentifierKey string `json:"identifier_key,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Message       string `json:"message,omitempty"`
}
