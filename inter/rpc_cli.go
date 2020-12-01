package inter

import (
	"context"

	"google.golang.org/grpc"
)

type RpcInter interface {
	//获取连接
	GetRpc() grpc.ClientConnInterface
	//获取上下文
	GetCtx() context.Context
	//释放资源
	Close()
}

type GetRpcConn interface {
	//获取Rpc连接信息
	Conn(addr string) (RpcInter, error)
}
