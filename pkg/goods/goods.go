// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package goods

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/houseme/union-jd-go/pkg/entity"
)

type iGoods struct {
}

// Goods .
func Goods() *iGoods {
	return &iGoods{}
}

// QueryCate query category
func (i *iGoods) QueryCate(ctx context.Context, req *entity.OpenCategoryGoodsGetRequest, logger glog.ILogger) (res *entity.OpenCategoryGoodsGetResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-goods-QueryCate")
	defer span.End()

	logger.Info(ctx, "query category start params:", req)
	client := g.Client().Timeout(5*time.Second).SetHeader("", "")
	client.Get(ctx, "")
	return nil, nil
}
