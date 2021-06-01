package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haupc/cartransplant/auth/middleware"
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
	RegisterCar(ctx *gin.Context)
	UpdateCar(ctx *gin.Context)
	DeleteCar(ctx *gin.Context)
	ListMyCar(ctx *gin.Context)
	ListUserTrip(ctx *gin.Context)
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

func (c *carController) ListMyCar(ctx *gin.Context) {
	limitString := ctx.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit <= 0 {
		log.Printf("Parse limit err")
		limit = 10
	}
	respose, _ := c.carClient.ListMyCar(middleware.RPCNewContextFromContext(ctx), &grpcproto.Int{Value: int64(limit)})
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", respose.Cars))
}

func (c *carController) DeleteCar(ctx *gin.Context) {
	var deleteCarIDList []int32
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &deleteCarIDList)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	respose, err := c.carClient.DeleteCar(middleware.RPCNewContextFromContext(ctx), &grpcproto.DeleteCarRequest{Ids: deleteCarIDList})
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(respose.Value, "success", nil))
}

func (c *carController) UpdateCar(ctx *gin.Context) {
	var updateCarRequest grpcproto.CarObject
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &updateCarRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	respose, err := c.carClient.UpdateCar(middleware.RPCNewContextFromContext(ctx), &updateCarRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(respose.Value, "success", nil))
}

func (c *carController) RegisterCar(ctx *gin.Context) {
	var registerCarRequest grpcproto.CarObject
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &registerCarRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	respose, err := c.carClient.RegisterCar(middleware.RPCNewContextFromContext(ctx), &registerCarRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(respose.Value, "success", nil))

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
		BeginLeaveTime: registerTripRequest.BeginLeaveTime,
		EndLeaveTime:   registerTripRequest.EndLeaveTime,
		From:           registerTripRequest.From.ToGrpcPoint(),
		To:             registerTripRequest.To.ToGrpcPoint(),
		MaxDistance:    registerTripRequest.MaxDistance,
		CarID:          registerTripRequest.CarID,
		FeeEachKm:      registerTripRequest.FeeEachKm,
	}
	respose, err := c.carClient.RegisterTrip(middleware.RPCNewContextFromContext(ctx), request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(respose.Value, "success", nil))
}

func (c *carController) FindTrip(ctx *gin.Context) {
	var findTripRequest dto.TripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &findTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	if findTripRequest.BeginLeaveTime < time.Now().Unix() {
		findTripRequest.BeginLeaveTime = time.Now().Unix()
	}
	if findTripRequest.EndLeaveTime < findTripRequest.BeginLeaveTime {
		respose := utils.BuildErrorResponse("Wrong time to search", err.Error(), body)
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
	response, err := c.carClient.FindTrip(middleware.RPCNewContextFromContext(ctx), request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	var respData []dto.FindTripResponse
	json.Unmarshal(response.JsonResponse, &respData)
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", respData))

}

func (c *carController) ListUserTrip(ctx *gin.Context) {
	stateString := ctx.Query("state")
	var state int
	var err error
	if stateString != "" {
		state, err = strconv.Atoi(stateString)
		if err != nil {
			log.Printf("Parse state err")
			state = 0
		}
	}
	response, err := c.carClient.ListUserTrip(middleware.RPCNewContextFromContext(ctx), &grpcproto.Int{Value: int64(state)})
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", response.UserTrip))
}

func (c *carController) TakeTrip(ctx *gin.Context) {
	var findTripRequest dto.TripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &findTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	request := &grpcproto.TakeTripRequest{
		DriverTripID:   int32(findTripRequest.DriverTripID),
		BeginLeaveTime: findTripRequest.BeginLeaveTime,
		EndLeaveTime:   findTripRequest.EndLeaveTime,
		From:           findTripRequest.From.ToGrpcPoint(),
		To:             findTripRequest.From.ToGrpcPoint(),
		Seat:           int32(findTripRequest.Seat),
	}
	_, err = c.carClient.TakeTrip(middleware.RPCNewContextFromContext(ctx), request)
	if err != nil {
		respose := utils.BuildErrorResponse("Take trip failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))

}
