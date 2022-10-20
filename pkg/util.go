/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package pkg

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
)

type util struct {
	logger glog.ILogger
}

// NewUtil helper util
func NewUtil(log glog.ILogger) *util {
	return &util{
		logger: log,
	}
}

// NewRequest new request
func (l *util) NewRequest(ctx context.Context, c *config.Config, method, params string) (req *entity.Request, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-logic-NewRequest")
	defer span.End()

	l.logger.Info(ctx, "new request start params:", params)
	req = &entity.Request{
		Method:       method,
		Version:      config.RequestVersion,
		Timestamp:    gtime.Now().Format(config.RequestTimestamp),
		Format:       config.RequestFormat,
		SignMethod:   config.RequestSignMethod,
		AppKey:       c.AppKey(),
		BuyParamJSON: params,
	}
	if req.Sign, err = l.MakeSign(ctx, c.AppSecret(), *req); err != nil {
		return nil, gerror.Wrap(err, "make sign failed")
	}
	l.logger.Info(ctx, "new request end result:", req)
	return req, nil
}

// MakeSign make sign
func (l *util) MakeSign(ctx context.Context, appSecret string, req entity.Request) (string, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-logic-MakeSign")
	defer span.End()

	hash := md5.New()
	if _, err := hash.Write([]byte(appSecret + l.ConcatenateSignSource(ctx, req) + appSecret)); err != nil {
		err = gerror.Wrap(err, "write string failed")
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil))), nil
}

// ConcatenateSignSource get sign url 排序并拼接签名的内容信息
func (l *util) ConcatenateSignSource(ctx context.Context, data interface{}) string {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-ConcatenateSignSource")
	defer span.End()

	var (
		tt     = reflect.TypeOf(data)
		v      = reflect.ValueOf(data)
		count  = v.NumField()
		keys   = make([]string, count)
		params = make(map[string]string)
	)

	l.logger.Info(ctx, "helper ConcatenateSignSource tt", tt, " v", v)

	for i := 0; i < count; i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			l.logger.Printf(ctx, "%s %s = %v -tag:%s", tt.Field(i).Name, tt.Field(i).Type, v.Field(i).Interface(), tt.Field(i).Tag)
			tagName := tt.Field(i).Tag.Get("json")
			keys[i] = tagName
			params[tagName] = gconv.String(v.Field(i).Interface())
		}
	}
	// sort params
	sort.Strings(keys)
	var buf bytes.Buffer
	for i := range keys {
		if params[keys[i]] == "" || keys[i] == "sign" {
			continue
		}
		if _, err := buf.WriteString(keys[i] + "=" + params[keys[i]]); err != nil {
			l.logger.Error(ctx, "helper ConcatenateSignSource buf.WriteString", err)
			return ""
		}
	}
	l.logger.Info(ctx, "helper ConcatenateSignSource string start:", buf.String())
	return buf.String()
}
