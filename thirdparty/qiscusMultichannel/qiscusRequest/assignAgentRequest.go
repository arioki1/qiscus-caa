package qiscusRequest

type AssignAgentRequest struct {
	AgentID            int    `json:"agent_id"`
	RoomID             string `json:"room_id"`
	ReplaceLatestAgent bool   `json:"replace_latest_agent"`
}
