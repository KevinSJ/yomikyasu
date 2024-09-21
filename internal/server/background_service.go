package server

import (
	"context"
	"yomikyasu/internal/runner"
)

func (s *Server) StartBackgroundRunners(ctx context.Context) {
    podCastRunner := runner.New(ctx, s.db)
    podCastRunner.RunOnce(ctx)
}
