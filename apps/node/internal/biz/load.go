package biz

import (
	"context"
	"golang.org/x/sync/errgroup"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"

	"prometheus-manager/api"
	pb "prometheus-manager/api/strategy/v1/load"

	"prometheus-manager/pkg/curl"
	"prometheus-manager/pkg/util/dir"

	"prometheus-manager/apps/node/internal/conf"
	"prometheus-manager/apps/node/internal/service"
)

type (
	ILoadRepo interface {
		V1Repo
		LoadStrategy(ctx context.Context, path []string) error
	}

	LoadLogic struct {
		logger *log.Helper
		repo   ILoadRepo
	}
)

var _ service.ILoadLogic = (*LoadLogic)(nil)

func NewLoadLogic(repo ILoadRepo, logger log.Logger) *LoadLogic {
	return &LoadLogic{repo: repo, logger: log.NewHelper(log.With(logger, "module", loadModuleName))}
}

func (l *LoadLogic) Reload(ctx context.Context, _ *pb.ReloadRequest) (*pb.ReloadReply, error) {
	ctx, span := otel.Tracer(loadModuleName).Start(ctx, "LoadLogic.Reload")
	defer span.End()

	var eg errgroup.Group

	datasource := conf.Get().GetStrategy().GetPromDatasources()
	for _, promDatasource := range datasource {
		eg.Go(func() error {
			dirList := promDatasource.GetPath()
			configPath := conf.GetConfigPath()
			err := l.repo.LoadStrategy(ctx, dir.MakeDirs(configPath, dirList...))
			if err != nil {
				l.logger.Errorf("LoadStrategy err: %v", err)
				return err
			}

			out, err := curl.Curl(ctx, promDatasource.GetReloadPath())
			if err != nil {
				l.logger.Errorf("Curl err: %v", err)
				return err
			}
			l.logger.Infof("Curl out: %v", out)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		l.logger.Errorf("Wait err: %v", err)
		return nil, err
	}

	return &pb.ReloadReply{
		Response:  &api.Response{Message: l.repo.V1(ctx)},
		Timestamp: time.Now().Unix(),
	}, nil
}
