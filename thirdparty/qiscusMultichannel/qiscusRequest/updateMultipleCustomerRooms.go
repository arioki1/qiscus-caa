package qiscusRequest

type UpdateMultipleCustomerRooms []CustomerRoom
type CustomerRoom struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}
