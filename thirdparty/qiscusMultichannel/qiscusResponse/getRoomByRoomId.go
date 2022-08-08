package qiscusResponse

import "time"

type GetRoomByRoomId struct {
	Data struct {
		CustomerRoom struct {
			ChannelId               int         `json:"channel_id"`
			ContactId               int         `json:"contact_id"`
			Id                      interface{} `json:"id"`
			IsHandledByBot          bool        `json:"is_handled_by_bot"`
			IsResolved              bool        `json:"is_resolved"`
			IsWaiting               bool        `json:"is_waiting"`
			LastCommentSender       string      `json:"last_comment_sender"`
			LastCommentSenderType   string      `json:"last_comment_sender_type"`
			LastCommentText         string      `json:"last_comment_text"`
			LastCommentTimestamp    time.Time   `json:"last_comment_timestamp"`
			LastCustomerCommentText interface{} `json:"last_customer_comment_text"`
			LastCustomerTimestamp   time.Time   `json:"last_customer_timestamp"`
			Name                    string      `json:"name"`
			RoomBadge               string      `json:"room_badge"`
			RoomId                  string      `json:"room_id"`
			RoomType                string      `json:"room_type"`
			Source                  string      `json:"source"`
			UserAvatarUrl           string      `json:"user_avatar_url"`
			UserId                  string      `json:"user_id"`
		} `json:"customer_room"`
	} `json:"data"`
	Status int `json:"status"`
}
