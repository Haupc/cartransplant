package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haupc/cartransplant/auth/utils"
	"github.com/haupc/cartransplant/car/client"
	"github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

var _carController *carController

type CarController interface {
	TakeTrip(ctx *gin.Context)
	RegisterTrip(ctx *gin.Context)
	FindTrip(ctx *gin.Context)
}

type carController struct {
	carClient grpcproto.CarClient
}

func GetCarController() CarController {
	if _carController == nil {
		_carController = &carController{
			carClient: client.GetCarClient(),
		}
	}
	return _carController
}

func (c *carController) TakeTrip(ctx *gin.Context) {

}

func (c *carController) RegisterTrip(ctx *gin.Context) {
	var registerTripRequest dto.RegisterTripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &registerTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	request := &grpcproto.RegisterTripRequest{
		StartTime: registerTripRequest.StartTime,
		From:      registerTripRequest.From.ToGrpcPoint(),
		To:        registerTripRequest.To.ToGrpcPoint(),
	}
	respose, err := c.carClient.RegisterTrip(ctx, request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(respose.Value, "success", nil))
}

func (c *carController) FindTrip(ctx *gin.Context) {
	var findTripRequest dto.FindTripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &findTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	request := &grpcproto.FindTripRequest{
		BeginLeaveTime: findTripRequest.BeginLeaveTime,
		EndLeaveTime:   findTripRequest.EndLeaveTime,
		From:           findTripRequest.From.ToGrpcPoint(),
		To:             findTripRequest.To.ToGrpcPoint(),
		Option:         findTripRequest.Opt,
	}
	response, err := c.carClient.FindTrip(ctx, request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	var respData []dto.FindTripResponse
	json.Unmarshal(response.JsonResponse, &respData)
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", respData))

}
