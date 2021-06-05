package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/utils"
)

var _notifyController *notifyController

type NotifyController interface {
	PushNotify(ctx *gin.Context)
}

type notifyController struct {
	fcmClient *messaging.Client
}

func GetNotifyController() NotifyController {
	if _notifyController == nil {
		_notifyController = &notifyController{
			fcmClient: config.GetFcmClient(),
		}
	}
	return _notifyController
}

func (n *notifyController) PushNotify(ctx *gin.Context) {
	var notifyRequest messaging.Message
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &notifyRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	response, err := n.fcmClient.Send(ctx, &notifyRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Push noti failed", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", response))
}
