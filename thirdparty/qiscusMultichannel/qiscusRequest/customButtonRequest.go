package qiscusRequest

type CustomButtonRequest struct {
	AdditionalInfo []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"additional_info"`
	Agent struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"agent"`
	ChannelId   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelType string `json:"channel_type"`
	Customer    struct {
		Avatar string `json:"avatar"`
		Name   string `json:"name"`
		UserId string `json:"user_id"`
	} `json:"customer"`
	CustomerProperties []struct {
		Id    int    `json:"id"`
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"customer_properties"`
	Notes  string `json:"notes"`
	RoomId int    `json:"room_id"`
	Tag    []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tag"`
}
