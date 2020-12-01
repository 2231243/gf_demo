package packed

import (
	"context"
	"log"
	"time"

	"gf-app/inter"

	"github.com/gogf/gf/errors/gerror"
	"github.com/shimingyah/pool"
	"google.golang.org/grpc"
)

type GetRpcConn interface {
	//获取Rpc连接信息
	Conn(addr string) (inter.RpcInter, error)
}

type RpcPool struct {
	pool map[string]pool.Pool
}

type RpcClient struct {
	conn pool.Conn
	ctx  context.Context
}

func (rcc RpcClient) GetRpc() grpc.ClientConnInterface {
	return rcc.conn.Value()
}

func (rcc RpcClient) GetCtx() context.Context {
	return rcc.ctx
}

func (rcc RpcClient) Close() {
	if rcc.conn != nil {
		_ = rcc.conn.Close()
	}
	rcc.conn = nil
	rcc.ctx = nil
}

func (rc RpcPool) Conn(addr string) (inter.RpcInter, error) {
	if p, ok := rc.pool[addr]; ok {
		conn, err := p.Get()
		if err != nil {
			return nil, err
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		rcc := RpcClient{
			conn: conn,
			ctx:  ctx,
		}
		return rcc, nil
	}
	return nil, gerror.New("not found rpc " + addr)
}

func NewRpcPool(addrs map[string]pool.Options) GetRpcConn {
	r := RpcPool{pool: make(map[string]pool.Pool)}
	for k, v := range addrs {
		p, err := pool.New(k, v)
		if err != nil {
			log.Fatal(k + "is not connect")
		}
		r.pool[k] = p
	}
	return r
}
