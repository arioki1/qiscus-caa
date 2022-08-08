package delivery

import (
	"github.com/arioki1/qiscus-caa/helpers"
	"github.com/arioki1/qiscus-caa/src/api/v1/model"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel/qiscusRequest"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"net/http"
)

type qiscusWebhookDelivery struct {
	qiscusWebhookUC model.QiscusWebhookUseCase
	clientAsynq     *asynq.Client
}
type QiscusWebhookDelivery interface {
	Mount(group *gin.RouterGroup)
}

func (qw qiscusWebhookDelivery) Mount(group *gin.RouterGroup) {
	group.POST("/custom-agent-allocation", qw.CustomAgentAllocation)
}

func NewQiscusWebhookDelivery(qiscusWebhookUC model.QiscusWebhookUseCase, asynq *asynq.Client) QiscusWebhookDelivery {
	return &qiscusWebhookDelivery{
		qiscusWebhookUC: qiscusWebhookUC,
		clientAsynq:     asynq,
	}
}

func (qw *qiscusWebhookDelivery) CustomAgentAllocation(c *gin.Context) {
	var requestWebhook qiscusRequest.CustomAgentAllocationRequest
	if err := c.Bind(&requestWebhook); err != nil {
		helpers.PrintErrStringLog(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid JSON body",
		})
		return
	}
	res, _, err := qw.qiscusWebhookUC.CustomAgentAllocation(c, &requestWebhook, qw.clientAsynq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
