package dto

type Maneuver struct {
	Exit          int       `json:"exit"`
	BearingAfter  int       `json:"bearing_after"`
	BearingBefore int       `json:"bearing_before"`
	Location      []float64 `json:"location"`
	Modifier      string    `json:"modifier"`
	Type          string    `json:"type"`
}

type Step struct {
	Maneuver    `json:"maneuver,omitempty"`
	Mode        string  `json:"mode"`
	DrivingSide string  `json:"driving_side"`
	Name        string  `json:"name"`
	Weight      float64 `json:"weight"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
	Ref         string  `json:"ref,omitempty"`
}

type Leg struct {
	Steps    []Step  `json:"steps"`
	Summary  string  `json:"summary"`
	Weight   float64 `json:"weight"`
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`
}

type Route struct {
	Legs       []Leg   `json:"legs"`
	WeightName string  `json:"weight_name"`
	Weight     float64 `json:"weight"`
	Duration   float64 `json:"duration"`
	Distance   float64 `json:"distance"`
}

type Waypoint struct {
	Hint     string    `json:"hint"`
	Distance float64   `json:"distance"`
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
}

type RouteResponse struct {
	Code      string     `json:"code"`
	Routes    []Route    `json:"routes"`
	Waypoints []Waypoint `json:"waypoints"`
}
