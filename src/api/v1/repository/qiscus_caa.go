package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/helpers"
	"github.com/arioki1/qiscus-caa/src/api/v1/model"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel/qiscusRequest"
	"github.com/hibiken/asynq"
	"time"
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
	limit := 50
	reqGetAgent := qiscusRequest.GetAgents{
		RoomId: &roomId,
		Limit:  &limit,
	}
	alreadyAssigned := false
	var errorAssigned error

	for !alreadyAssigned {
		agents, err := qClient.GetAgents(ctx, reqGetAgent)
		if err != nil {
			alreadyAssigned = true
			errorAssigned = err
		}

		if len(agents.Data.Agents) > 0 {
			for _, agent := range agents.Data.Agents {
				if agent.IsAvailable && agent.CurrentCustomerCount < 2 && agent.TypeAsString == "agent" {
					reqAssignAgent := qiscusRequest.AssignAgentRequest{
						AgentID:            agent.Id,
						RoomID:             roomId,
						ReplaceLatestAgent: true,
					}

					if _, err := qClient.AssignAgentToChatRoom(ctx, reqAssignAgent); err != nil {
						alreadyAssigned = true
						errorAssigned = err
					} else {
						alreadyAssigned = true
					}
					break
				}
			}
		} else {
			alreadyAssigned = true
			errorAssigned = fmt.Errorf("no agent available")
		}

		if !alreadyAssigned {
			//sleep 30 seconds
			time.Sleep(30 * time.Second)
		}
	}

	return nil, 0, errorAssigned
}

func NewQiscusCAARepository(cfg config.Config) model.QiscusCAARepository {
	return &qiscusTask{
		cfg: cfg,
	}
}
