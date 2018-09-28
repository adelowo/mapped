package main

import (
	"bufio"
	"flag"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/adelowo/gotils/registry"
	"github.com/adelowo/mapped/config"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/pborman/uuid"
)

var (
	BuildDate string

	Version string
)

func main() {

	var (
		httpAddr   = flag.String("http.addr", ":1400", "Port to run HTTP server at")
		consulAddr = flag.String("discovery.consul", "localhost:1500", "Consul discovery address")
	)

	shutDownChan := make(chan os.Signal)
	signal.Notify(shutDownChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var bufferedWriter = bufio.NewWriter(os.Stderr)

	logger := log.NewEntry(&log.Logger{
		Handler: logfmt.New(os.Stderr),
		Level:   log.ErrorLevel,
	})

	host, err := os.Hostname()
	if err != nil {
		logger.WithError(err).Fatal("could not fetch host name")
	}

	logger = logger.WithField("app", "mapped").
		WithField("host", host)

	defaultCfg := consul.DefaultConfig()
	defaultCfg.Address = *consulAddr

	client, err := consul.NewClient(defaultCfg)
	if err != nil {
		logger.WithError(err).Fatal("could not create consul client")
	}

	cfg, err := config.FromConsul(client)
	if err != nil {
		logger.WithError(err).Fatal("error while fetching config from consul")
	}

	cfg.HTTPPort = *httpAddr
	cfg.ServerID = host

	ip, err := ipAddr()
	if err != nil {
		logger.WithError(err).Fatalf("could not determine IP address to register this service with... %v", err)
	}

	registrar := registry.NewWithClient(client)

	pp, err := strconv.Atoi(cfg.HTTPPort[1:])
	if err != nil {
		logger.WithError(err).Fatal("could not convert http addr port to an int")
	}

	svc := &consul.AgentServiceRegistration{
		ID:   uuid.New(),
		Name: "mapped",
		Port: pp,
		Tags: []string{"urlprefix-/mapped"},
		Check: &consul.AgentServiceCheck{
			TLSSkipVerify: true,
			Method:        "GET",
			Timeout:       "20s",
			Interval:      "1m",
			HTTP:          "http://" + ip.String() + cfg.HTTPPort + "/health",
			Name:          "HTTP check for mapped",
		},
	}

	if err := registrar.RegisterService(svc); err != nil {
		logger.WithError(err).Fatal("could not register service in consul")
	}

	logger.Logger.Handler = logfmt.New(bufferedWriter)

	<-shutDownChan
	bufferedWriter.Flush()
}

func ipAddr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return nil, nil

}
