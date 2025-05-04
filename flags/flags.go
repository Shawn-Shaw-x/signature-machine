package flags

import "github.com/urfave/cli/v2"

const envVarPrefix = "SIGNATURE"

func prefixEnvVars(name string) []string {
	return []string{envVarPrefix + "_" + name}
}

var (
	// RpcHostFlag RPC Service
	RpcHostFlag = &cli.StringFlag{
		Name:     "rpc-host",
		Usage:    "The host of the rpc",
		EnvVars:  prefixEnvVars("RPC_HOST"),
		Required: true,
	}
	RpcPortFlag = &cli.IntFlag{
		Name:     "rpc-port",
		Usage:    "The port of the rpc",
		EnvVars:  prefixEnvVars("RPC_PORT"),
		Value:    8983,
		Required: true,
	}
	// LevelDbPathFlag Database
	LevelDbPathFlag = &cli.StringFlag{
		Name:    "master-db-host",
		Usage:   "The path of the leveldb",
		EnvVars: prefixEnvVars("LEVEL_DB_PATH"),
		Value:   "./",
	}
)

var requireFlags = []cli.Flag{
	RpcHostFlag,
	RpcPortFlag,
	LevelDbPathFlag,
}

var Flags []cli.Flag

func init() {
	Flags = append(requireFlags)
}
