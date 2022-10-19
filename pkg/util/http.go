// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
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

// Sign md5 sign
func Sign(params map[string]string, appSecret string) string {
	// params key list
	keyList := make([]string, 0, len(params))
	for k := range params {
		keyList = append(keyList, k)
	}
	// sort
	sort.Strings(keyList)
	var b strings.Builder
	for _, key := range keyList {
		if val, ok := params[key]; ok {
			if _, err := b.WriteString(fmt.Sprintf("%s%s", key, val)); err != nil {
				return ""
			}
		}
	}

	hash := md5.New()
	if _, err := hash.Write([]byte(appSecret + b.String() + appSecret)); err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}
