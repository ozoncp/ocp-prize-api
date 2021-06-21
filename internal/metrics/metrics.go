package metrics

import "github.com/prometheus/client_golang/prometheus"

var crudProcessesByHandler *prometheus.CounterVec = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "crud_processes_by_handler", // metric name
		Help: "Number of successful calls of crud functions by handler",
	},
	[]string{"handler"}, // labels
)

// RegisterMetrics by default crud handler
func RegisterMetrics() {
	prometheus.MustRegister(crudProcessesByHandler)
}

func incSuccessfulRequestByHandler(handler string, counts int) {
	crudProcessesByHandler.With(prometheus.Labels{"handler": handler}).Add(float64(counts))
}

// IncrementSuccessfulCreate with label create
func IncrementSuccessfulCreate(counts int) {
	incSuccessfulRequestByHandler("create", counts)
}

// IncrementSuccessfulMultiCreate with label multi_create
func IncrementSuccessfulMultiCreate(counts int) {
	incSuccessfulRequestByHandler("multi_create", counts)
}

// IncrementSuccessfulUpdate with label update
func IncrementSuccessfulUpdate(counts int) {
	incSuccessfulRequestByHandler("update", counts)
}

// IncrementSuccessfulDescribe with label describe
func IncrementSuccessfulDescribe(counts int) {
	incSuccessfulRequestByHandler("describe", counts)
}

// IncrementSuccessfulList with label list
func IncrementSuccessfulList(counts int) {
	incSuccessfulRequestByHandler("list", counts)
}

// IncrementSuccessfulRemove with label remove
func IncrementSuccessfulRemove(counts int) {
	incSuccessfulRequestByHandler("remove", counts)
}
