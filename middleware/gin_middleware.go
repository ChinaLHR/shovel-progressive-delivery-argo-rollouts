package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	gin_http_responses_cnt = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "gin_metric_handler_http_responses_status_total",
		Help: "Total number of scrapes by Gin HTTP status code.",
	}, []string{"code"})

	gin_http_metric_ignore_path = []string{"/ping", "/metrics"}
)

func GinHttpMetric() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		//get reponse http status
		status := ctx.Writer.Status()

		path := ctx.Request.URL.Path
		ignorePath := false
		for _, v := range gin_http_metric_ignore_path {
			if v == path {
				ignorePath = true
				continue
			}
		}
		if !ignorePath {
			gin_http_responses_cnt.With(prometheus.Labels{"code": strconv.Itoa(status)}).Inc()
		}
	}
}
