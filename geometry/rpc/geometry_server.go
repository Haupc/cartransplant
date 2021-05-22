package geometry

type geometryServer struct{}

const (
	locationHost = "http://localhost"
	routingHost  = "http://localhost:5000"
)

func NewGeometryServer() *geometryServer {
	return &geometryServer{}
}
