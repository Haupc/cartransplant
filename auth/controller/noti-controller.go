package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/middleware"
	"github.com/haupc/cartransplant/auth/utils"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/notify/client"
)

var _notifyController *notifyController

type NotifyController interface {
	PushNotify(ctx *gin.Context)
	GetNotify(ctx *gin.Context)
	RegisterToken(ctx *gin.Context)
}

type notifyController struct {
	fcmClient    *messaging.Client
	notifyClient grpcproto.NotifyClient
}

func GetNotifyController() NotifyController {
	if _notifyController == nil {
		_notifyController = &notifyController{
			fcmClient:    config.GetFcmClient(),
			notifyClient: client.GetNotifyClient(),
		}
	}
	return _notifyController
}

func (n *notifyController) RegisterToken(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		respose := utils.BuildErrorResponse("Token empty", "", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	_, err := n.notifyClient.AddUserToken(middleware.RPCNewContextFromContext(ctx), &grpcproto.AddUserTokenReq{
		Token: token,
	})
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

func (n *notifyController) GetNotify(ctx *gin.Context) {
	request := &grpcproto.GetNotifyRequest{
		Limit:  1000,
		Offset: 0,
	}
	response, err := n.notifyClient.GetNotify(middleware.RPCNewContextFromContext(ctx), request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", response.Notifications))
}

func (n *notifyController) PushNotify(ctx *gin.Context) {
	var pushNotiRQ PushNotifyRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &pushNotiRQ)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	request := &grpcproto.PushNotifyReq{
		Notification: &grpcproto.NotifyMessage{
			UserID:      pushNotiRQ.UserID,
			CreatedTime: time.Now().Unix(),
			Title:       pushNotiRQ.Title,
			Message:     pushNotiRQ.Message,
			Image:       pushNotiRQ.Image,
		},
	}
	_, err = n.notifyClient.PushNotify(middleware.RPCNewContextFromContext(ctx), request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

type PushNotifyRequest struct {
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Message string `json:"message"`
}
