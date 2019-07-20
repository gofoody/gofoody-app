package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofoody/gofoody-app/pkg/config"
	"github.com/gofoody/gofoody-app/pkg/ctrl"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()

	initLogger(cfg.GetLogLevel())

	router := mountEndpoints(cfg)
	startService(cfg.GetHttpPort(), router)
}

func initLogger(logLevel string) {
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func mountEndpoints(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	statusCtrl := ctrl.NewStatusCtrl()
	r.HandleFunc("/api/status", statusCtrl.Show)

	loginCtrl := ctrl.NewLoginCtrl(cfg)
	r.HandleFunc("/", loginCtrl.Home)
	r.HandleFunc("/login", loginCtrl.Login).Methods("GET", "POST")
	r.HandleFunc("/signin", loginCtrl.Signin).Methods("POST")
	r.HandleFunc("/signup", loginCtrl.Signup).Methods("POST")
	r.HandleFunc("/register", loginCtrl.Register).Methods("POST")

	return r
}

func startService(port int, router *mux.Router) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Infof("gofoody app running at:%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("failed to start gofoody app, error:%v", err)
	}
}
