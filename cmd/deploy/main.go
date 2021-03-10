package main

import (
	"context"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/navikt/deployment/pkg/deployer"
	"github.com/navikt/deployment/pkg/pb"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func main() {
	err := run()
	if err == nil {
		return
	}
	code := deployer.ErrorExitCode(err)
	if code == deployer.ExitInvocationFailure {
		flag.Usage()
	}
	log.Errorf("fatal: %s", err)
	os.Exit(int(code))
}

func run() error {
	// Configuration and context
	cfg := deployer.NewConfig()
	deployer.InitConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	// Logging
	deployer.SetupLogging(*cfg)

	// Welcome
	log.Infof("Starting NAIS deploy")

	// Prepare request
	request, err := deployer.Prepare(ctx, cfg)
	if err != nil {
		return err
	}

	// Set up asynchronous gRPC connection
	grpcConnection, err := deployer.NewGrpcConnection(*cfg)
	if err != nil {
		return err
	}
	defer func() {
		err := grpcConnection.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	d := deployer.Deployer{
		Client: pb.NewDeployClient(grpcConnection),
	}

	if cfg.PrintPayload {
		marsh := jsonpb.Marshaler{Indent: "  "}
		err = marsh.Marshal(os.Stdout, request)
		if err != nil {
			log.Errorf("print payload: %s", err)
		}
	}

	if cfg.DryRun {
		return nil
	}

	return d.Deploy(ctx, cfg, request)
}
