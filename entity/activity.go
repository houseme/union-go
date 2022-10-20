/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-go.
 */

package entity

// JDUnionOpenActivityQueryTopLevel is a struct.
type JDUnionOpenActivityQueryTopLevel struct {
	JdUnionOpenActivityQueryResponse JDUnionOpenActivityQueryResponse `json:"jd_union_open_activity_query_responce"`
}

// JDUnionOpenActivityQueryResponse is a struct.
type JDUnionOpenActivityQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

// JDUnionOpenActivityQueryResult is a struct.
type JDUnionOpenActivityQueryResult struct {
	Code       int64                           `json:"code"`
	Data       []*JDUnionOpenActivityQueryData `json:"data"`
	Message    string                          `json:"message"`
	TotalCount int                             `json:"totalCount"`
	RequestID  string                          `json:"requestId"`
}

// JDUnionOpenActivityQueryData is a struct.
type JDUnionOpenActivityQueryData struct {
	ActStatus    int                              `json:"actStatus"`
	Advantage    string                           `json:"advantage"`
	Content      string                           `json:"content"`
	DownloadCode string                           `json:"downloadCode"`
	DownloadURL  string                           `json:"downloadUrl"`
	EndTime      int                              `json:"endTime"`
	ID           int                              `json:"id"`
	PlatformType int                              `json:"platformType"`
	StartTime    int                              `json:"startTime"`
	Tag          string                           `json:"tag"`
	Title        string                           `json:"title"`
	UpdateTime   int                              `json:"updateTime"`
	URLM         string                           `json:"urlM"`
	URLPC        string                           `json:"urlPC"`
	CategoryList []*JDUnionOpenActivityQueryCate  `json:"categoryList"`
	ImgList      []*JDUnionOpenActivityQueryImage `json:"imgList"`
}

// JDUnionOpenActivityQueryCate is a struct.
type JDUnionOpenActivityQueryCate struct {
	CategoryID int `json:"categoryId"`
	Type       int `json:"type"`
}

// JDUnionOpenActivityQueryImage is a struct.
type JDUnionOpenActivityQueryImage struct {
	ImgName     string `json:"imgName"`
	ImgURL      string `json:"imgUrl"`
	WidthHeight string `json:"widthHeight"`
}
