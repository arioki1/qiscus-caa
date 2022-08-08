package model

import (
	"context"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel/qiscusRequest"
	"github.com/hibiken/asynq"
)

type (
	QiscusCAATaskPayload struct {
		Payload    interface{}
		RoomId     string
		TypeAction string
	}
	QiscusWebhook        struct{}
	QiscusWebhookUseCase interface {
		CustomAgentAllocation(ctx context.Context, req *qiscusRequest.CustomAgentAllocationRequest, clientAsynq *asynq.Client) (interface{}, int, error)
	}
	QiscusCAARepository interface {
		NewTaskAssignAgent(roomId string) (*asynq.Task, error)
		AssignAgent(ctx context.Context, roomId string) (interface{}, int, error)
	}
)
