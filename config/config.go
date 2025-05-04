package config

import (
	"github.com/urfave/cli/v2"
	"signature-machine/flags"
)

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	LevelDbPath string
	RpcServer   ServerConfig
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		LevelDbPath: ctx.String(flags.LevelDbPathFlag.Name),
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
