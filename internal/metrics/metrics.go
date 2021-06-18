package metrics

import "github.com/prometheus/client_golang/prometheus"

var crudProcessesByHandler *prometheus.CounterVec = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "crud_processes_by_handler", // metric name
		Help: "Number of successful calls of crud functions by handler",
	},
	[]string{"handler"}, // labels
)

func RegisterMetrics() {
	prometheus.MustRegister(crudProcessesByHandler)
}

func incSuccessfulRequestByHandler(handler string, counts int) {
	crudProcessesByHandler.With(prometheus.Labels{"handler": handler}).Add(float64(counts))
}

func IncrementSuccessfulCreate(counts int) {
	incSuccessfulRequestByHandler("create", counts)
}

func IncrementSuccessfulMultiCreate(counts int) {
	incSuccessfulRequestByHandler("multi_create", counts)
}

func IncrementSuccessfulUpdate(counts int) {
	incSuccessfulRequestByHandler("update", counts)
}

func IncrementSuccessfulDescribe(counts int) {
	incSuccessfulRequestByHandler("describe", counts)
}

func IncrementSuccessfulList(counts int) {
	incSuccessfulRequestByHandler("list", counts)
}

func IncrementSuccessfulRemove(counts int) {
	incSuccessfulRequestByHandler("remove", counts)
}
