package qiscusRequest

type SetUserInfo struct {
	UserProperties []UserProperties `json:"user_properties,omitempty"`
}

type UserProperties struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
