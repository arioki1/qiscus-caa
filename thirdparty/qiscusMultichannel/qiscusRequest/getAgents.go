package qiscusRequest

type GetAgents struct {
	RoomId            *string `url:"room_id,omitempty"`
	Limit             *int    `url:"limit,omitempty"`
	CursorAfter       *string `url:"cursor_after,omitempty"`
	CursorBefore      *string `url:"cursor_before,omitempty"`
	IsAvailableInRoom *bool   `url:"is_available_in_room,omitempty"`
}
