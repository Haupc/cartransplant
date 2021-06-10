package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/haupc/cartransplant/auth/utils"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/grpcproto"
)

var _geometryController *geometryController

type GeometryController interface {
	GetCurrentAddress(ctx *gin.Context)
	SearchAddress(ctx *gin.Context)
	GetRouting(ctx *gin.Context)
}

type geometryController struct {
	geomClient grpcproto.GeometryClient
}

func GetGeometryController() GeometryController {
	if _geometryController == nil {
		_geometryController = &geometryController{
			geomClient: client.GetGeomClient(),
		}
	}
	return _geometryController
}

func (c *geometryController) GetCurrentAddress(ctx *gin.Context) {
	latitude := ctx.Query("lat")
	longitude := ctx.Query("long")
	point := &grpcproto.Point{
		Latitude:  latitude,
		Longitude: longitude,
	}
	response, err := c.geomClient.GetCurrentAddress(ctx, point)
	if err != nil {
		glog.V(3).Infof("GetCurrentAddress - Error: %v", err)
		errResp := utils.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
	}
	var resp map[string]interface{}
	err = json.Unmarshal(response.JsonResponse, &resp)
	if err != nil {
		glog.V(3).Infof("GetCurrentAddress - Error: %v", err)
		errResp := utils.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", resp))
}

func (c *geometryController) SearchAddress(ctx *gin.Context) {
	query := ctx.Query("query")
	// query, _ = base.Normalize(query)
	glog.V(4).Infof("SearchAddress - query: %s", query)
	if query == "" {
		errResp := utils.BuildErrorResponse("Invalid Request", "Query Empty", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	q := &grpcproto.SearchRequest{
		Query: query,
	}
	response, err := c.geomClient.SearchAddress(ctx, q)
	if err != nil {
		glog.V(3).Infof("SearchAddress - Error: %v", err)
		errResp := utils.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}
	var resp []map[string]interface{}
	err = json.Unmarshal(response.JsonResponse, &resp)
	if err != nil {
		glog.V(3).Infof("SearchAddress - Error: %v", err)
		errResp := utils.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", resp))
}

func (c *geometryController) GetRouting(ctx *gin.Context) {
	from := &grpcproto.Point{
		Latitude:  ctx.Query("fromLat"),
		Longitude: ctx.Query("fromLong"),
	}
	to := &grpcproto.Point{
		Latitude:  ctx.Query("toLat"),
		Longitude: ctx.Query("toLong"),
	}
	routingRequest := &grpcproto.RouteRequest{
		From: from,
		To:   to,
	}
	response, err := c.geomClient.GetRouting(ctx, routingRequest)

	if err != nil {
		errResp := utils.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}
	var resp map[string]interface{}
	err = json.Unmarshal(response.JsonResponse, &resp)
	if err != nil {
		errResp := utils.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(true, "success", resp))

}
