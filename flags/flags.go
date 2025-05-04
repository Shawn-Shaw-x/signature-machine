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
	CredentialsFileFlag = &cli.StringFlag{
		Name:    "credentials-file",
		Usage:   "the credentials file of cloud hsm",
		EnvVars: prefixEnvVars("CREDENTIALS_FILE"),
	}
	KeyNameFlag = &cli.StringFlag{
		Name:    "key-name",
		Usage:   "The key name of cloud hsm",
		EnvVars: prefixEnvVars("KEY_NAME"),
	}
	HsmEnable = &cli.BoolFlag{
		Name:    "hsm-enable",
		Usage:   "Hsm enable",
		EnvVars: prefixEnvVars("HSM_ENABLE"),
		Value:   false,
	}
)

var requireFlags = []cli.Flag{
	RpcHostFlag,
	RpcPortFlag,
	LevelDbPathFlag,
}

var optionalFlags = []cli.Flag{
	CredentialsFileFlag,
	KeyNameFlag,
	HsmEnable,
}

var Flags []cli.Flag

func init() {
	Flags = append(requireFlags, optionalFlags...)
}
