package qiscusRequest

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

type CustomChannelWebhookRequest struct {
	AppCode string `json:"app_code"`
	Payload struct {
		From struct {
			AvatarUrl string `json:"avatar_url"`
			Email     string `json:"email"`
			Id        int    `json:"id"`
			IdStr     string `json:"id_str"`
			Name      string `json:"name"`
		} `json:"from"`
		Message struct {
			CommentBeforeId    int         `json:"comment_before_id"`
			CommentBeforeIdStr string      `json:"comment_before_id_str"`
			CreatedAt          time.Time   `json:"created_at"`
			DisableLinkPreview bool        `json:"disable_link_preview"`
			Extras             interface{} `json:"extras"`
			Id                 int         `json:"id"`
			IdStr              string      `json:"id_str"`
			Payload            interface{} `json:"payload"`
			Text               string      `json:"text"`
			Timestamp          time.Time   `json:"timestamp"`
			Type               string      `json:"type"`
			UniqueTempId       string      `json:"unique_temp_id"`
			UnixNanoTimestamp  int64       `json:"unix_nano_timestamp"`
			UnixTimestamp      int         `json:"unix_timestamp"`
		} `json:"message"`
		Room struct {
			Id              int    `json:"id"`
			IdStr           string `json:"id_str"`
			IsPublicChannel bool   `json:"is_public_channel"`
			Name            string `json:"name"`
			Options         string `json:"options"`
			Participants    []struct {
				Active                   bool        `json:"active"`
				AvatarUrl                string      `json:"avatar_url"`
				Email                    string      `json:"email"`
				Extras                   interface{} `json:"extras"`
				Id                       int         `json:"id"`
				IdStr                    string      `json:"id_str"`
				LastCommentReadId        int         `json:"last_comment_read_id"`
				LastCommentReadIdStr     string      `json:"last_comment_read_id_str"`
				LastCommentReceivedId    int         `json:"last_comment_received_id"`
				LastCommentReceivedIdStr string      `json:"last_comment_received_id_str"`
				Username                 string      `json:"username"`
			} `json:"participants"`
			RoomAvatar string `json:"room_avatar"`
			TopicId    int    `json:"topic_id"`
			TopicIdStr string `json:"topic_id_str"`
			Type       string `json:"type"`
		} `json:"room"`
	} `json:"payload"`
	Type string `json:"type"`
}

type googlePlayReviewID struct {
	GooglePlayReviewID *string
	MultichannelEmail  *string
}

func (r *CustomChannelWebhookRequest) GetGooglePlayReviewID() (*googlePlayReviewID, error) {
	var errResult error
	var result *googlePlayReviewID

	isGooglePlay := false
	payload := r.Payload
	fromEmail := payload.From.Email
	listParticipant := payload.Room.Participants

	if len(listParticipant) == 0 {
		return nil, errors.New("participant is empty")
	}

	for _, participant := range listParticipant {
		participantEmail := participant.Email
		if participantEmail == fromEmail {
			continue
		}

		if len(participantEmail) > 3 {
			firstCharacter := participantEmail[0:3]
			character := "gp_"
			if firstCharacter == character {
				isGooglePlay = true
				encodeId := fmt.Sprintf(participantEmail[3:])
				reviewId, err := base64.RawURLEncoding.DecodeString(encodeId)
				if err != nil {
					errResult = err
					break
				}
				gReviewId := string(reviewId)
				result = &googlePlayReviewID{
					GooglePlayReviewID: &gReviewId,
					MultichannelEmail:  &participantEmail,
				}
			}
		}
	}
	if !isGooglePlay {
		errResult = errors.New("is not google play custom channel")
	}

	return result, errResult
}
