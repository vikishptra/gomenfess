package runregisteruser

import (
	"context"
)

type runregisteruserInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runregisteruserInteractor{
		outport: outputPort,
	}
}

func (r *runregisteruserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	return res, nil
}
