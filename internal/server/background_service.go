package server

import (
	"context"
	"yomikyasu/internal/runner"
)

func (s *Server) StartBackgroundRunners(ctx context.Context) {
	podCastRunner := runner.New(ctx, s.db)
	podCastRunner.RunOnce(ctx)
}

// SQLite has disable foreign key support by default, so need to enable it
// everytime.
func (s *Server) ConfigDb() {
	s.db.Config()
}
