package qiscusRequest

type SendMessageTextRequest struct {
	SenderEmail string `json:"sender_email,omitempty"`
	Message     string `json:"message,omitempty"`
	Type        string `json:"type,omitempty"`
	RoomID      string `json:"room_id,omitempty"`
}
