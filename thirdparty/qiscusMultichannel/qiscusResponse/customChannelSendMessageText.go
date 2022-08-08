package qiscusResponse

type CustomChannelSendMessageText struct {
	Data struct {
		Agent        interface{} `json:"agent"`
		AgentService interface{} `json:"agent_service"`
		RoomLog      struct {
			ChannelId             int         `json:"channel_id"`
			Extras                interface{} `json:"extras"`
			HasNoMessage          bool        `json:"has_no_message"`
			IsWaiting             bool        `json:"is_waiting"`
			Name                  string      `json:"name"`
			Resolved              bool        `json:"resolved"`
			ResolvedTs            interface{} `json:"resolved_ts"`
			RoomBadge             string      `json:"room_badge"`
			RoomId                string      `json:"room_id"`
			Source                string      `json:"source"`
			StartServiceCommentId interface{} `json:"start_service_comment_id"`
			UserAvatarUrl         string      `json:"user_avatar_url"`
			UserId                string      `json:"user_id"`
		} `json:"room_log"`
	} `json:"data"`
	Status int `json:"status"`
}
