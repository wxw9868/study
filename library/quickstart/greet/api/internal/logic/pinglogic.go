package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/wxw9868/study/library/quickstart/greet/api/internal/svc"
	"github.com/wxw9868/study/library/quickstart/greet/api/internal/types"
	"github.com/wxw9868/study/library/quickstart/greet/rpc/greet"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Resp, err error) {
	if _, err = l.svcCtx.GreetRpc.Ping(l.ctx, new(greet.Placeholder)); err != nil {
		return
	}

	resp = new(types.Resp)
	resp.Msg = "pong"

	return
}
