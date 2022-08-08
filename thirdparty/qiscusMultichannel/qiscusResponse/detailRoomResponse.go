package qiscusResponse

type DetailRoomResponse struct {
	Data struct {
		CustomerRoom struct {
			ChannelID             int         `json:"channel_id"`
			ID                    int         `json:"id"`
			IsHandledByBot        bool        `json:"is_handled_by_bot"`
			IsResolved            bool        `json:"is_resolved"`
			IsWaiting             bool        `json:"is_waiting"`
			LastCommentText       interface{} `json:"last_comment_text"`
			LastCommentTimestamp  string      `json:"last_comment_timestamp"`
			LastCustomerTimestamp interface{} `json:"last_customer_timestamp"`
			Name                  string      `json:"name"`
			RoomBadge             interface{} `json:"room_badge"`
			RoomID                string      `json:"room_id"`
			Source                string      `json:"source"`
			UserAvatarURL         string      `json:"user_avatar_url"`
			UserID                string      `json:"user_id"`
		} `json:"customer_room"`
	} `json:"data"`
	Status int `json:"status"`
}
