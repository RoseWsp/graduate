package service

import (
	"context"
	v1 "graduate/api/album/job/v1"
	"graduate/app/album/job/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type JobService struct {
	v1.UnimplementedJobServer

	ac  *biz.JobUseCase
	log *log.Helper
}

func NewJobService(ac *biz.JobUseCase, logger log.Logger) *JobService {
	return &JobService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/job")),
	}

}

func (s *JobService) IntegratingJob(ctx context.Context, req *v1.IntegratingJobReq) (*v1.IntegratingJobReply, error) {
	return &v1.IntegratingJobReply{}, s.ac.IntegratingJob(ctx)
}
