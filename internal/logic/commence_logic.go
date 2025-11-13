package logic

import (
	"context"
	"fmt"

	"package_memorizing/internal/svc"
	"package_memorizing/internal/types"
	"package_memorizing/model"

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

// Commence demonstrates how to use the package_memorizing table via the generated model
// It performs a lookup by (uid, package_id). If not found, it inserts a new record.
// If found, it updates the record (example: toggles the PackageType) and shows how to delete.
func (l *CommenceLogic) Commence(req *types.CommenceRequest) (resp *types.CommenceResponse, err error) {
	// contract:
	// - input: req.Atom.UID (int64) and req.PackageID (int64)
	// - output: CommenceResponse (empty for this demo) ; error on DB failure

	uid := uint64(req.UID)
	packageID := uint64(req.PackageID)

	// get model bound to the correct shard/table for this package id
	m := l.svcCtx.GetPackageModelById(int64(packageID))

	// 1) Try to find existing record by (uid, package_id)
	pm, err := m.FindOneByUidPackageId(l.ctx, uid, packageID)
	if err != nil {
		if err == model.ErrNotFound {
			// not found -> create one
			newRec := &model.PackageMemorizing{
				Uid:         uid,
				PackageType: 1, // example default type
				PackageId:   packageID,
			}
			if _, err = m.Insert(l.ctx, newRec); err != nil {
				l.Logger.Error(err)
				return nil, fmt.Errorf("failed to insert package_memorizing: %w", err)
			}
			l.Logger.Infof("inserted package_memorizing for uid=%d package_id=%d", uid, packageID)
			return &types.CommenceResponse{}, nil
		}
		// other DB error
		l.Logger.Error(err)
		return nil, err
	}

	// 2) If found, update example: change PackageType (this is just illustrative)
	pm.PackageType = pm.PackageType + 1
	if err = m.Update(l.ctx, pm); err != nil {
		l.Logger.Error(err)
		return nil, fmt.Errorf("failed to update package_memorizing: %w", err)
	}
	l.Logger.Infof("updated package_memorizing id=%d uid=%d package_id=%d", pm.Id, pm.Uid, pm.PackageId)

	// 3) Example delete logic: if PackageType reaches a certain value, delete it (demo)
	if pm.PackageType > 10 {
		if err = m.Delete(l.ctx, pm.Id); err != nil {
			l.Logger.Error(err)
			return nil, fmt.Errorf("failed to delete package_memorizing: %w", err)
		}
		l.Logger.Infof("deleted package_memorizing id=%d", pm.Id)
	}

	return &types.CommenceResponse{}, nil
}
