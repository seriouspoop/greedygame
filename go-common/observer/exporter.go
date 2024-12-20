package observer

type ExporterType int

const (
	ConsoleExporter ExporterType = iota
	OTLPHttpExporter
	OTLPGrpcExporter
)

type Exporter struct {
	Type         ExporterType
	HttpEndpoint string
	GrcpEndpoint string
}

func NewDevelopmentExporter() *Exporter {
	return &Exporter{
		Type:         ConsoleExporter,
		HttpEndpoint: "localhost:4318",
		GrcpEndpoint: "localhost:4317",
	}
}

// returns a grpc otlp exporter.
// serviceName is the opentelemetry collector service name.
func NewProductionExporter(serviceName string) *Exporter {
	return &Exporter{
		Type:         OTLPGrpcExporter,
		HttpEndpoint: serviceName + ":4318",
		GrcpEndpoint: serviceName + ":4317",
	}
}
