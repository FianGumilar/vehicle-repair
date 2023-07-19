package utils

import (
	"context"
	"time"

	"github.com/FianGumilar/vehicle-repair/domain"
)

func ResponseInterceptor(ctx context.Context, resp *domain.ApiResponse) {
	traceIdf := ctx.Value("requestid")
	traceId := ""

	if traceIdf != nil {
		traceId = traceIdf.(string)
	}

	resp.Timestamp = time.Now()
	resp.TraceID = traceId
}
