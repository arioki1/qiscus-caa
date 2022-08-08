package qiscusResponse

type LoginQiscusResponse struct {
	Data struct {
		User struct {
			Id                  int           `json:"id"`
			Name                string        `json:"name"`
			Email               string        `json:"email"`
			AuthenticationToken string        `json:"authentication_token"`
			CreatedAt           string        `json:"created_at"`
			UpdatedAt           string        `json:"updated_at"`
			SdkEmail            string        `json:"sdk_email"`
			SdkKey              string        `json:"sdk_key"`
			IsAvailable         bool          `json:"is_available"`
			Type                int           `json:"type"`
			AvatarUrl           string        `json:"avatar_url"`
			AppId               int           `json:"app_id"`
			IsVerified          bool          `json:"is_verified"`
			NotificationsRoomId string        `json:"notifications_room_id"`
			BubbleColor         interface{}   `json:"bubble_color"`
			QismoKey            string        `json:"qismo_key"`
			DirectLoginToken    interface{}   `json:"direct_login_token"`
			LastLogin           string        `json:"last_login"`
			ForceOffline        bool          `json:"force_offline"`
			DeletedAt           interface{}   `json:"deleted_at"`
			TypeAsString        string        `json:"type_as_string"`
			AssignedRules       []interface{} `json:"assigned_rules"`
			App                 struct {
				Id                             int         `json:"id"`
				Name                           string      `json:"name"`
				AppCode                        string      `json:"app_code"`
				SecretKey                      string      `json:"secret_key"`
				CreatedAt                      string      `json:"created_at"`
				UpdatedAt                      string      `json:"updated_at"`
				BotWebhookUrl                  string      `json:"bot_webhook_url"`
				IsBotEnabled                   bool        `json:"is_bot_enabled"`
				AllocateAgentWebhookUrl        string      `json:"allocate_agent_webhook_url"`
				IsAllocateAgentWebhookEnabled  bool        `json:"is_allocate_agent_webhook_enabled"`
				MarkAsResolvedWebhookUrl       string      `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool        `json:"is_mark_as_resolved_webhook_enabled"`
				IsMobilePnEnabled              bool        `json:"is_mobile_pn_enabled"`
				IsActive                       bool        `json:"is_active"`
				IsSessional                    bool        `json:"is_sessional"`
				IsAgentAllocationEnabled       bool        `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool        `json:"is_agent_takeover_enabled"`
				IsTokenExpiring                bool        `json:"is_token_expiring"`
				PaidChannelApproved            interface{} `json:"paid_channel_approved"`
				UseLatest                      bool        `json:"use_latest"`
				FreeSessions                   int         `json:"free_sessions"`
				IsForceSendBot                 bool        `json:"is_force_send_bot"`
				AppConfig                      struct {
					Id                     int         `json:"id"`
					AppId                  int         `json:"app_id"`
					Widget                 interface{} `json:"widget"`
					CreatedAt              string      `json:"created_at"`
					UpdatedAt              string      `json:"updated_at"`
					OfflineMessage         interface{} `json:"offline_message"`
					OnlineMessage          interface{} `json:"online_message"`
					Timezone               string      `json:"timezone"`
					EnableBulkAssign       bool        `json:"enable_bulk_assign"`
					SendOnlineIfResolved   bool        `json:"send_online_if_resolved"`
					SendOfflineEachMessage bool        `json:"send_offline_each_message"`
				} `json:"app_config"`
				AgentRoles []struct {
					Id            int    `json:"id"`
					AppId         int    `json:"app_id"`
					Name          string `json:"name"`
					IsDefaultRole bool   `json:"is_default_role"`
					CreatedAt     string `json:"created_at"`
					UpdatedAt     string `json:"updated_at"`
				} `json:"agent_roles"`
			} `json:"app"`
		} `json:"user"`
		Details struct {
			IsIntegrated bool `json:"is_integrated"`
			SdkUser      struct {
				Id          int    `json:"id"`
				Token       string `json:"token"`
				Email       string `json:"email"`
				Password    string `json:"password"`
				DisplayName string `json:"display_name"`
				AvatarUrl   string `json:"avatar_url"`
				Extras      struct {
					Type            string      `json:"type"`
					UserBubbleColor interface{} `json:"user_bubble_color"`
				} `json:"extras"`
			} `json:"sdk_user"`
			App struct {
				AppCode                        string `json:"app_code"`
				SecretKey                      string `json:"secret_key"`
				Name                           string `json:"name"`
				BotWebhookUrl                  string `json:"bot_webhook_url"`
				IsBotEnabled                   bool   `json:"is_bot_enabled"`
				IsAllocateAgentWebhookEnabled  bool   `json:"is_allocate_agent_webhook_enabled"`
				AllocateAgentWebhookUrl        string `json:"allocate_agent_webhook_url"`
				MarkAsResolvedWebhookUrl       string `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool   `json:"is_mark_as_resolved_webhook_enabled"`
				IsActive                       bool   `json:"is_active"`
				IsSessional                    bool   `json:"is_sessional"`
				IsAgentAllocationEnabled       bool   `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool   `json:"is_agent_takeover_enabled"`
				UseLatest                      bool   `json:"use_latest"`
				IsBulkAssignmentEnabled        bool   `json:"is_bulk_assignment_enabled"`
			} `json:"app"`
		} `json:"details"`
		LongLivedToken string `json:"long_lived_token"`
		UserConfigs    struct {
			Notifagentjoining           interface{} `json:"notifagentjoining"`
			IsNotifagentjoiningEnabled  bool        `json:"is_notifagentjoining_enabled"`
			Notifmessagecoming          interface{} `json:"notifmessagecoming"`
			IsNotifmessagecomingEnabled bool        `json:"is_notifmessagecoming_enabled"`
		} `json:"user_configs"`
	} `json:"data"`
}
