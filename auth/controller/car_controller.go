package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	auth_dto "github.com/haupc/cartransplant/auth/dto"
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
	ListDriverTrip(ctx *gin.Context)
	RegisterTripUser(ctx *gin.Context)
	FindPendingTrip(ctx *gin.Context)
	UserCancelTrip(ctx *gin.Context)
	MarkUserTripDone(ctx *gin.Context)
	DriverTakeTrip(ctx *gin.Context)
	RegisterActiveZone(ctx *gin.Context)
	ListActiveZone(ctx *gin.Context)
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

func (c *carController) RegisterActiveZone(ctx *gin.Context) {
	var activeZones grpcproto.ActiveZone
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Println("RegisterActiveZone - request: ", string(body))
	err := json.Unmarshal(body, &activeZones)
	if err != nil {
		log.Printf("RegisterActiveZone - Error: %v", err)
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	_, err = c.carClient.RegisterActiveZone(middleware.RPCNewContextFromContext(ctx), &activeZones)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

func (c *carController) ListActiveZone(ctx *gin.Context) {
	activeZones, err := c.carClient.ListActiveZone(middleware.RPCNewContextFromContext(ctx), &empty.Empty{})
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", activeZones.Provinces))
}

func (c *carController) DriverTakeTrip(ctx *gin.Context) {
	var takeTripRequest grpcproto.DriverTakeTripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Println("DriverTakeTrip - request: ", string(body))
	err := json.Unmarshal(body, &takeTripRequest)
	if err != nil {
		log.Printf("DriverTakeTrip - Error: %v", err)
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	_, err = c.carClient.DriverTakeTrip(middleware.RPCNewContextFromContext(ctx), &takeTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

func (c *carController) MarkUserTripDone(ctx *gin.Context) {
	idString := ctx.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		respose := utils.BuildErrorResponse("Param id invalid", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	_, err = c.carClient.MarkUserTripDone(middleware.RPCNewContextFromContext(ctx), &grpcproto.Int{Value: int64(id)})
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

func (c *carController) UserCancelTrip(ctx *gin.Context) {
	idString := ctx.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		respose := utils.BuildErrorResponse("Param id invalid", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	_, err = c.carClient.CancelTrip(middleware.RPCNewContextFromContext(ctx), &grpcproto.Int{Value: int64(id)})
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

func (c *carController) FindPendingTrip(ctx *gin.Context) {
	var findPendingTripRequest grpcproto.FindPendingTripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Println("FindPendingTrip - request: ", string(body))
	err := json.Unmarshal(body, &findPendingTripRequest)
	if err != nil {
		log.Printf("FindPendingTrip - Error: %v", err)
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	response, err := c.carClient.FindPendingTrip(context.Background(), &findPendingTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", response))

}

func (c *carController) RegisterTripUser(ctx *gin.Context) {
	var takeTripRequest grpcproto.TakeTripRequest
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &takeTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	// TODO : logic here
	_, err = c.carClient.UserRegisterTrip(middleware.RPCNewContextFromContext(ctx), &takeTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))
}

func (c *carController) ListDriverTrip(ctx *gin.Context) {
	startdateString := ctx.Query("startdate")
	startdate, _ := strconv.Atoi(startdateString)
	enddateString := ctx.Query("enddate")
	enddate, _ := strconv.Atoi(enddateString)
	stateString := ctx.Query("state")
	state, err := strconv.Atoi(stateString)
	if err != nil {
		log.Printf("Parse limit err")
		state = 1
	}
	request := &grpcproto.ListDriverTripRequest{
		State:     int32(state),
		StartDate: int32(startdate),
		EndDate:   int32(enddate),
	}
	respose, err := c.carClient.ListDriverTrip(middleware.RPCNewContextFromContext(ctx), request)
	if err != nil {
		respose := utils.BuildErrorResponse("Something wrong happened", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	var driverTripResponse []auth_dto.DriverTripResponse
	for _, trip := range respose.Trips {
		driverTripResponse = append(driverTripResponse, auth_dto.DriverTripRPCToDriverTripResponse(trip))
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", driverTripResponse))

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
	log.Println(string(body))
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
		Seat:           int32(registerTripRequest.Seat),
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
	findTripRequest := &grpcproto.FindTripRequest{}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, findTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	if findTripRequest.BeginLeaveTime < time.Now().Unix() {
		findTripRequest.BeginLeaveTime = time.Now().Unix()
	}
	if findTripRequest.EndLeaveTime < findTripRequest.BeginLeaveTime {
		respose := utils.BuildErrorResponse("Wrong time to search", "", body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	response, err := c.carClient.FindTrip(middleware.RPCNewContextFromContext(ctx), findTripRequest)
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
	takeTripRequest := &grpcproto.TakeTripRequest{}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Println(string(body))
	err := json.Unmarshal(body, takeTripRequest)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	if takeTripRequest.Type == 1 || takeTripRequest.DriverTripID == 0 {
		_, err = c.carClient.UserRegisterTrip(middleware.RPCNewContextFromContext(ctx), takeTripRequest)
	} else {
		_, err = c.carClient.TakeTrip(middleware.RPCNewContextFromContext(ctx), takeTripRequest)
	}
	if err != nil {
		respose := utils.BuildErrorResponse("Take trip failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", nil))

}
