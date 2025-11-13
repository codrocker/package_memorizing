package logic

import (
	"context"

	"package_memorizing/internal/svc"
	"package_memorizing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommenceLogic {
	return &CommenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommenceLogic) Commence(req *types.CommenceRequest) (resp *types.CommenceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
