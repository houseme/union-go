// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package jd

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

const (
	version   = "0.0.1"
	serverURL = "https://router.jd.com/api"
)

// UnionJD is a jd service.
type UnionJD struct {
	accessToken string
	appKey      string
	appSecret   string
	logger      glog.ILogger
}

// NewUnionJD returns a jd service.
func NewUnionJD(ctx context.Context, appKey, appSecret, accessToken string) (*UnionJD, error) {
	return &UnionJD{
		appKey:      appKey,
		appSecret:   appSecret,
		accessToken: accessToken,
	}, nil
}

// Version returns the version of the jd service.
func (u *UnionJD) Version() string {
	return version
}

// QueryCate query category
func (u *UnionJD) QueryCate(ctx context.Context) {

}
