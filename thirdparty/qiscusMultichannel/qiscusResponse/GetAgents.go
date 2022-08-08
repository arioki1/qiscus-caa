package qiscusResponse

type GetAgents struct {
	Data struct {
		Agents []struct {
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
		} `json:"agents"`
	} `json:"data"`
	Meta struct {
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Status int `json:"status"`
}
