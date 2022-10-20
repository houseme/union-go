// Copyright union-jd-go Author(https://houseme.github.io/union-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-go.

package jd

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/union-go/config"
	"github.com/houseme/union-go/goods"
	"github.com/houseme/union-go/orders"
)

// UnionJD is a jd service.
type UnionJD struct {
}

// NewUnionJD returns a jd service.
func NewUnionJD() *UnionJD {
	return &UnionJD{}
}

// Version returns the version of the jd service.
func (u *UnionJD) Version() string {
	return config.Version
}

// GetGoods to get goods
func (u *UnionJD) GetGoods(ctx context.Context, c *config.Config) (*goods.Goods, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-GetGoods")
	defer span.End()

	return goods.NewGoods(ctx, c)
}

// GetOrders to get orders
func (u *UnionJD) GetOrders(ctx context.Context, c *config.Config) (*orders.Orders, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-GetOrders")
	defer span.End()

	return orders.NewOrders(ctx, c)
}
