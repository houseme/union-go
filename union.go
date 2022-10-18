// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package jd

const (
	version = "0.0.1"
)

// UnionJD is a jd service.
type UnionJD struct {
	// contains filtered or unexported fields
}

// NewUnionJdGo returns a jd service.
func NewUnionJdGo() *UnionJD {
	return &UnionJD{}
}

// Version returns the version of the jd service.
func (u *UnionJD) Version() string {
	return version
}
