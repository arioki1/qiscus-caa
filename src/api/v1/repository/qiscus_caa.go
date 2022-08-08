package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/helpers"
	"github.com/arioki1/qiscus-caa/src/api/v1/model"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel"
	"github.com/hibiken/asynq"
)

type qiscusTask struct {
	cfg config.Config
}

func (q *qiscusTask) NewTaskAssignAgent(roomId string) (*asynq.Task, error) {
	payload := model.QiscusCAATaskPayload{
		Payload:    nil,
		RoomId:     roomId,
		TypeAction: "assign_agent",
	}
	marshal, err := json.Marshal(payload)
	if err != nil {
		helpers.PrintErrStringLog(fmt.Sprintf("error marshal room id: %+v e:%+v", roomId, err))
		return nil, err
	}
	return asynq.NewTask(q.cfg.GetAppName()+":assign_agent", marshal), nil
}

func (q *qiscusTask) AssignAgent(ctx context.Context, roomId string) (interface{}, int, error) {
	qConfig := qiscusMultichannel.MultichannelConfig{
		AppID:     q.cfg.GetQiscusAppId(),
		SecretKey: q.cfg.GetQiscusSecretKey(),
		BaseUrl:   q.cfg.GetQismoBaseURL(),
	}
	qClient := qiscusMultichannel.NewQiscusMultichannelClient(qConfig)

	//TODO ASSIGN AGENT
	//delay 1 minute
	fmt.Println(qClient)
	return nil, 0, nil
}

func NewQiscusCAARepository(cfg config.Config) model.QiscusCAARepository {
	return &qiscusTask{
		cfg: cfg,
	}
}
