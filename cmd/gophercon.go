package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ghouscht/k8s-gopherconeu/pkg/router"
	"github.com/ghouscht/k8s-gopherconeu/pkg/server"
	"github.com/ghouscht/k8s-gopherconeu/version"
)

func main() {
	log.Printf(
		"Service is starting, version is %s, commit is %s, build time %s ...",
		version.Release, version.Commit, version.BuildTime,
	)

	sig := make(chan os.Signal, 1)
	shutdown := make(chan error, 2)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	servicePort := os.Getenv("SERVICE_PORT")
	if len(servicePort) == 0 {
		log.Print("use SERVICE_PORT to define a coustom port where the server listens (default: 8080)")
		servicePort = "8080"
	}

	diagPort := os.Getenv("DIAG_PORT")
	if len(diagPort) == 0 {
		log.Print("use DIAG_PORT to define a coustom port where the diag server listens (default: 8081)")
		diagPort = "8081"
	}

	srv := server.New("", servicePort, router.BaseRouter())
	go func() {
		if err := srv.Start(); err != nil {
			shutdown <- err
		}
	}()

	diagSrv := server.New("", diagPort, router.DiagnosticsRouter())
	go func() {
		if err := diagSrv.Start(); err != nil {
			shutdown <- err
		}
	}()

	select {
	case err := <-shutdown:
		log.Fatal(err)
	case s := <-sig:
		log.Printf("Got signal %s. Stopping...", s)
		if err := srv.Stop(); err != nil {
			log.Print(err)
		}
		if err := diagSrv.Stop(); err != nil {
			log.Print(err)
		}
		log.Print("stopped...")
	}
}
