package main

import(
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"go_api/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		fmt.Println(err)
	}
}