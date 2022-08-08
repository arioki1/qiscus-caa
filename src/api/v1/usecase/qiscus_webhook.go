package usecase

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

type qiscusWebhook struct {
	cfg         config.Config
	qRepository model.QiscusCAARepository
}

func (q *qiscusWebhook) CustomAgentAllocation(ctx context.Context, req *qiscusRequest.CustomAgentAllocationRequest, clientAsynq *asynq.Client) (interface{}, int, error) {
	task, err := q.qRepository.NewTaskAssignAgent(req.RoomId)
	if err != nil {
		helpers.PrintErrStringLog(fmt.Sprintf("failed to create task set CustomAgentAllocation error: %v", err.Error()))
	}

	isNewTask, err := q.createNewTask(req.RoomId)
	if err != nil {
		return nil, 0, err
	}

	if isNewTask {
		deadline := time.Now().Add(time.Hour * 24)
		if _, err := clientAsynq.Enqueue(task, asynq.MaxRetry(2), asynq.Deadline(deadline)); err != nil {
			if err.Error() != "task ID conflicts with another task" {
				helpers.PrintErrStringLog(fmt.Sprintf("failed to create task set CustomAgentAllocation error: %v", err.Error()))
			}
		} else {
			helpers.PrintInfoStringLog(fmt.Sprintf("add task assign agent in room id %v", req.RoomId))
		}

		qConfig := qiscusMultichannel.MultichannelConfig{
			AppID:     q.cfg.GetQiscusAppId(),
			SecretKey: q.cfg.GetQiscusSecretKey(),
			BaseUrl:   q.cfg.GetQismoBaseURL(),
		}
		qClient := qiscusMultichannel.NewQiscusMultichannelClient(qConfig)
		msgRequest := qiscusRequest.SendMessageTextRequest{
			SenderEmail: q.cfg.GetQiscusAdminEmail(),
			Message:     "Silahkan tunggu agent akan datang menjawab pertanyaan anda",
			Type:        "text",
			RoomID:      req.RoomId,
		}

		if err := qClient.SendMessageText(ctx, msgRequest); err != nil {
			helpers.PrintErrStringLog(fmt.Sprintf("failed send message: %v", err.Error()))
		}
	}

	return struct {
		Message string `json:"message"`
	}{
		Message: "success create task",
	}, 0, nil
}

func (q *qiscusWebhook) createNewTask(roomId string) (bool, error) {
	newTask := true
	redisConnection, err := asynq.ParseRedisURI(q.cfg.GetRedisURL())
	if err != nil {
		helpers.PrintFatalStringLog(fmt.Sprintf("ParseRedisURI Error: %+v", err))
		return false, err
	}
	inspector := asynq.NewInspector(redisConnection)

	//Get Active task
	activeTasks, err := inspector.ListActiveTasks("default")
	if err == nil {
		for _, t := range activeTasks {
			payload := t.Payload
			data := new(model.QiscusCAATaskPayload)

			if err := json.Unmarshal(payload, data); err != nil {
				helpers.PrintErrStringLog(fmt.Sprintf("error unmarshal roomid: %+v e:%+v", data.RoomId, err))
			} else {
				if data.RoomId == roomId {
					newTask = false
				}
			}
		}
	} else {
		helpers.PrintFatalStringLog(fmt.Sprintf("ListActiveTasks Error: %+v", err))
	}

	//Get Pending task
	pendingTasks, err := inspector.ListPendingTasks("default")
	if err == nil {
		for _, t := range pendingTasks {
			payload := t.Payload
			data := new(model.QiscusCAATaskPayload)

			if err := json.Unmarshal(payload, data); err != nil {
				helpers.PrintErrStringLog(fmt.Sprintf("error unmarshal roomid: %+v e:%+v", data.RoomId, err))
			} else {
				if data.RoomId == roomId {
					newTask = false
				}
			}
		}
	} else {
		helpers.PrintFatalStringLog(fmt.Sprintf("ListPendingTasks Error: %+v", err))
	}

	//Get Retry task
	retryTasks, err := inspector.ListRetryTasks("default")
	if err == nil {
		for _, t := range retryTasks {
			payload := t.Payload
			data := new(model.QiscusCAATaskPayload)

			if err := json.Unmarshal(payload, data); err != nil {
				helpers.PrintErrStringLog(fmt.Sprintf("error unmarshal roomid: %+v e:%+v", data.RoomId, err))
			} else {
				if data.RoomId == roomId {
					newTask = false
				}
			}
		}
	} else {
		helpers.PrintFatalStringLog(fmt.Sprintf("ListRetryTasks Error: %+v", err))
	}

	return newTask, nil
}

func NewQiscusUseCase(c config.Config, r model.QiscusCAARepository) model.QiscusWebhookUseCase {
	return &qiscusWebhook{
		cfg:         c,
		qRepository: r,
	}
}
