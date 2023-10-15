package alarmPage

import (
	"context"
)

type (
	// DeleteReq ...
	DeleteReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
	}

	// DeleteResp ...
	DeleteResp struct {
		// add response params
	}
)

// Delete ...
func (l *AlarmPage) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	// add your code here
	return &DeleteResp{}, nil
}