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
	Steps    []StepDTO `json:"steps"`
	Duration float64   `json:"duration"`
	Distance float64   `json:"distance"`
}
type RoutingDTO struct {
	Waypoints []WaypointDTO `json:"waypoints"`
	Routes    []RouteDTO    `json:"routes"`
}
