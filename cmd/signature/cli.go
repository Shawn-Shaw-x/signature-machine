package main

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
	"signature-machine-eth/common/cliapp"
	"signature-machine-eth/config"
	"signature-machine-eth/flags"
	"signature-machine-eth/leveldb"
	"signature-machine-eth/services/rpc"
)

func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("running grpc services...")
	cfg := config.NewConfig(ctx)
	grpcServerCfg := &rpc.RpcServerConfig{
		GrpcHostName: cfg.RpcServer.Host,
		GrpcPort:     cfg.RpcServer.Port,
	}
	db, err := leveldb.NewKeyStore(cfg.LevelDbPath)
	if err != nil {
		log.Error("failed to open keystore", "err", err)
	}
	return rpc.NewRpcServer(db, grpcServerCfg)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData),
		Description:          "Signature machine eth tool",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "version",
				Description: "Show version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run RPC",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
		},
	}
}
