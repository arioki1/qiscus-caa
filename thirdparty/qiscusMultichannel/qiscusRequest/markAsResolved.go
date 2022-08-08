package qiscusRequest

type MarkAsResolved struct {
	RoomId        string `json:"room_id"`
	LastCommentId string `json:"last_comment_id"`
	Notes         string `json:"notes,omitempty"`
	IsSendEmail   string `json:"is_send_email,omitempty"`
	Extras        string `json:"extras,omitempty"`
}
