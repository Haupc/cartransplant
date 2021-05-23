package dto

type WaypointDTO struct {
	Distance float64 `json:"distance"`
	Name     string  `json:"name"`
	Location Point   `json:"location"`
}

type StepDTO struct {
	Name     string `json:"name"`
	Location Point  `json:"location"`
}

type RouteDTO struct {
	Steps []StepDTO `json:"steps"`
}
type RoutingDTO struct {
	Waypoints []WaypointDTO `json:"waypoints"`
	Routes    []RouteDTO    `json:"routes"`
}
