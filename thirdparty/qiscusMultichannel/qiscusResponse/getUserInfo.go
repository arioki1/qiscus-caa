package qiscusResponse

import (
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel/qiscusRequest"
	"time"
)

type GetUserInfo struct {
	Data struct {
		Extras struct {
			AdditionalExtras struct {
				TimezoneOffset int `json:"timezone_offset"`
			} `json:"additional_extras"`
			Notes          interface{}                    `json:"notes"`
			TimezoneOffset interface{}                    `json:"timezone_offset"`
			UserProperties []qiscusRequest.UserProperties `json:"user_properties"`
		} `json:"extras"`
		FirstInitiated         time.Time   `json:"first_initiated"`
		FirstAgentResponseTime interface{} `json:"first_agent_response_time"`
		UserId                 string      `json:"user_id"`
		ChannelId              int         `json:"channel_id"`
		IsBlocked              bool        `json:"is_blocked"`
		ChannelName            string      `json:"channel_name"`
		Channel                struct {
			Id                         int         `json:"id"`
			Base64UserCredential       string      `json:"base64_user_credential"`
			BaseUrl                    string      `json:"base_url"`
			Token                      string      `json:"token"`
			IsActive                   bool        `json:"is_active"`
			CreatedAt                  string      `json:"created_at"`
			UpdatedAt                  string      `json:"updated_at"`
			IsSslEnabled               bool        `json:"is_ssl_enabled"`
			Platform                   string      `json:"platform"`
			Extras                     interface{} `json:"extras"`
			AppId                      int         `json:"app_id"`
			Name                       string      `json:"name"`
			BadgeUrl                   string      `json:"badge_url"`
			Hsm24Enabled               bool        `json:"hsm_24_enabled"`
			PhoneNumber                string      `json:"phone_number"`
			ForwardUrl                 interface{} `json:"forward_url"`
			ForwardEnabled             bool        `json:"forward_enabled"`
			ReadEnabled                bool        `json:"read_enabled"`
			BusinessId                 string      `json:"business_id"`
			AllowIntlHsm               bool        `json:"allow_intl_hsm"`
			OnSync                     bool        `json:"on_sync"`
			UseChannelResponder        bool        `json:"use_channel_responder"`
			PhoneNumberStatus          interface{} `json:"phone_number_status"`
			BusinessVerificationStatus interface{} `json:"business_verification_status"`
			AlterAppId                 int         `json:"alter_app_id"`
			IsAutoResponderEnabled     bool        `json:"is_auto_responder_enabled"`
			IsBotEnabled               bool        `json:"is_bot_enabled"`
			IsOnCloud                  bool        `json:"is_on_cloud"`
			PhoneNumberId              string      `json:"phone_number_id"`
			Pin                        string      `json:"pin"`
		} `json:"channel"`
		Hsm24Enabled bool `json:"hsm_24_enabled"`
	} `json:"data"`
}
