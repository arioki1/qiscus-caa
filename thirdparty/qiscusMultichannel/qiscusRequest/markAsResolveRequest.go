package qiscusRequest

type MarkAsResolveRequest struct {
	Customer struct {
		AdditionalInfo []interface{} `json:"additional_info"`
		Avatar         string        `json:"avatar"`
		Name           string        `json:"name"`
		UserID         string        `json:"user_id"`
	} `json:"customer"`
	ResolvedBy struct {
		Email       string `json:"email"`
		ID          int    `json:"id"`
		IsAvailable bool   `json:"is_available"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"resolved_by"`
	Service struct {
		FirstCommentID string `json:"first_comment_id"`
		ID             int    `json:"id"`
		IsResolved     bool   `json:"is_resolved"`
		//LastCommentID  int    `json:"last_comment_id"`
		Notes  string `json:"notes"`
		RoomID string `json:"room_id"`
		Source string `json:"source"`
	} `json:"service"`
}
