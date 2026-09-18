package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var httpRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições recebidas pelo endpoint /projeto-korp",
	},
	[]string{"method", "status"},
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

	resposta := Resposta{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	rec.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rec).Encode(resposta)

	httpRequestsTotal.WithLabelValues(r.Method, http.StatusText(rec.status)).Inc()
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
