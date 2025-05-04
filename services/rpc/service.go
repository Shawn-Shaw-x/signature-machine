package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"signature-machine-eth/common/cliapp"
	"signature-machine-eth/leveldb"
	"signature-machine-eth/protobuf/wallet"
	"strconv"
	"sync/atomic"
)

const MaxReceivedMessageSize = 1024 * 1024 * 30000

type RpcServer struct {
	*RpcServerConfig
	db *leveldb.Keys
	wallet.UnimplementedWalletServiceServer
	stopped atomic.Bool
}

func (r *RpcServer) Start(ctx context.Context) error {
	go func(r *RpcServer) {
		log.Info("rpc server started", "host", r.GrpcHostName, "port", r.GrpcPort)
		listenner, err := net.Listen("tcp", r.GrpcHostName+":"+strconv.Itoa(r.GrpcPort))
		if err != nil {
			log.Error("failed to listen", "err", err)
		}
		opt := grpc.MaxRecvMsgSize(MaxReceivedMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(nil),
		)
		reflection.Register(gs)

		wallet.RegisterWalletServiceServer(gs, r)

		log.Info("rpc server started", "host", r.GrpcHostName, "port", r.GrpcPort)
		if err := gs.Serve(listenner); err != nil {
			log.Error("failed to serve", "err", err)
		}

	}(r)
	return nil
}

func (r RpcServer) Stop(ctx context.Context) error {
	r.stopped.Store(true)
	return nil
}

func (r RpcServer) Stopped() bool {
	return r.stopped.Load()
}

type RpcServerConfig struct {
	GrpcHostName string
	GrpcPort     int
}

func NewRpcServer(db *leveldb.Keys, cfg *RpcServerConfig) (cliapp.Lifecycle, error) {
	return &RpcServer{
		RpcServerConfig: cfg,
		db:              db,
	}, nil

}
