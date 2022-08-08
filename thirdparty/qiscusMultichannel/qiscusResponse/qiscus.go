package qiscusResponse

type QiscusResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
