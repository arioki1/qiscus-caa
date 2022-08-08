package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/helpers"
	"github.com/arioki1/qiscus-caa/src/api/v1/model"
	"github.com/arioki1/qiscus-caa/src/api/v1/repository"
	"github.com/hibiken/asynq"
)

type workerServices struct {
	conf config.Config
}

func (w workerServices) HandleTaskAssignAgent(c context.Context, t *asynq.Task) error {

	payload := t.Payload()
	data := new(model.QiscusCAATaskPayload)

	if err := json.Unmarshal(payload, data); err != nil {
		helpers.PrintErrStringLog(fmt.Sprintf("error unmarshal roomid: %+v e:%+v", data.RoomId, err))
		return err
	}

	helpers.PrintInfoStringLog(fmt.Sprintf("Handle task %+v", data.TypeAction))

	repo := repository.NewQiscusCAARepository(w.conf)
	ctx := context.Background()

	if data.TypeAction == "assign_agent" {
		if _, _, e := repo.AssignAgent(ctx, data.RoomId); e != nil {
			helpers.PrintErrStringLog(fmt.Sprintf("error assign_agent room id: %+v e:%+v", data.RoomId, e))
			return e
		} else {
			helpers.PrintInfoStringLog("success assign_agent room id: " + data.RoomId)
		}
	}
	return nil
}

type Services interface {
	HandleTaskAssignAgent(c context.Context, t *asynq.Task) error
}

func NewWorkerServices(conf config.Config) Services {
	return &workerServices{
		conf: conf,
	}
}
