package qiscusRequest

type CustomAgentAllocationRequest struct {
	AppID          string `json:"app_id"`
	CandidateAgent struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		IsAvailable bool   `json:"is_available"`
	} `json:"candidate_agent"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	RoomId       string `json:"room_id"`
	IsNewSession bool   `json:"is_new_session"`
	IsResolved   bool   `json:"is_resolved"`
}
