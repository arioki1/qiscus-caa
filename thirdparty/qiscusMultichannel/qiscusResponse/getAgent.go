package qiscusResponse

type GetAgent struct {
	Data struct {
		Agent struct {
			AvatarUrl            string      `json:"avatar_url"`
			CreatedAt            string      `json:"created_at"`
			CurrentCustomerCount int         `json:"current_customer_count"`
			Email                string      `json:"email"`
			ForceOffline         bool        `json:"force_offline"`
			Id                   int         `json:"id"`
			IsAvailable          bool        `json:"is_available"`
			LastLogin            interface{} `json:"last_login"`
			Name                 string      `json:"name"`
			SdkEmail             string      `json:"sdk_email"`
			SdkKey               string      `json:"sdk_key"`
			Type                 int         `json:"type"`
			TypeAsString         string      `json:"type_as_string"`
			UserChannels         []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"user_channels"`
			UserRoles []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"user_roles"`
		} `json:"agent"`
	} `json:"data"`
	Status int `json:"status"`
}

func (agent *GetAgent) IsHandleCorrectChannel(channelId int) bool {
	result := false
	listChannel := agent.Data.Agent.UserChannels
	for i := 0; i < len(listChannel); i++ {
		if listChannel[i].Id == channelId {
			result = true
			break
		}
	}
	return result
}
