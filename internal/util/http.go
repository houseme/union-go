// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package util

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytedance/sonic/decoder"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HTTPGet get http request
func HTTPGet(url string) (string, error) {
	c, err := client.NewClient()
	if err != nil {
		return "", err
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("https://www.jd.com")
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v", string(res.Body()))

	var o = map[string]interface{}{}
	var r = strings.NewReader(`{"a":"b"}{"1":"2"}`)
	var dec = decoder.NewStreamDecoder(r)
	if err = dec.Decode(&o); err != nil {
		return "", err
	}
	fmt.Printf("%+v", o) // map[1:2 a:b]

	return string(res.Body()), nil
}
