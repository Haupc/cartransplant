package helper

import (
	"fmt"

	"github.com/haupc/cartransplant/geometry/dto"
)

func RouteResponseToDTO(response *dto.RouteResponse) dto.RoutingDTO {
	waypoints := []dto.WaypointDTO{}
	for _, waypoint := range response.Waypoints {
		waypoints = append(waypoints, WaypointToWaypointDTO(waypoint))
	}
	result := dto.RoutingDTO{
		Waypoints: waypoints,
		Routes:    []dto.RouteDTO{},
	}

	for _, route := range response.Routes {
		steps := []dto.StepDTO{}
		for _, step := range route.Legs[0].Steps {
			steps = append(steps, StepToStepDTO(step))
		}
		result.Routes = append(result.Routes, dto.RouteDTO{
			Steps: steps,
		})
	}
	return result

}

func LocationToPointDTO(location []float64) dto.Point {
	return dto.Point{
		Latitue:   fmt.Sprintf("%v", location[0]),
		Longitude: fmt.Sprintf("%v", location[1]),
	}
}

func WaypointToWaypointDTO(waypoint dto.Waypoint) dto.WaypointDTO {
	return dto.WaypointDTO{
		Distance: waypoint.Distance,
		Name:     waypoint.Name,
		Location: LocationToPointDTO(waypoint.Location),
	}
}

func StepToStepDTO(step dto.Step) dto.StepDTO {
	return dto.StepDTO{
		Name:     step.Name,
		Location: LocationToPointDTO(step.Location),
	}
}
