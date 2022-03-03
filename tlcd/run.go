package tlcd

import (
	"context"

	"github.com/lwch/tlc/proto"
)

// Run handle run command
func (sv *Service) Run(ctx context.Context, cfg *proto.RunConfig) (*proto.RunResponse, error) {
	sv.createContainer(cfg)
	return &proto.RunResponse{}, nil
}
