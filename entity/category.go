/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-go.
 */

package entity

// OpenCategoryGoodsGetResponse .
type OpenCategoryGoodsGetResponse struct {
	JDUnionOpenCategoryGoodsGetResponse *JDUnionOpenCategoryGoodsGetResponse `json:"jd_union_open_category_goods_get_response"`
}

// JDUnionOpenCategoryGoodsGetResponse jd union
type JDUnionOpenCategoryGoodsGetResponse struct {
	GetResult *GetResult `json:"getResult"`
}

// GetResult get result
type GetResult struct {
	Code    string         `json:"code"`
	Data    *GetResultData `json:"data"`
	Message string         `json:"message"`
}

// GetResultData get result data
type GetResultData struct {
	CategoryResp []*CategoryResp `json:"categoryResp"`
}

// CategoryResp category item response
type CategoryResp struct {
	Name     string `json:"name"`
	Grade    uint   `json:"grade,string"`
	ID       uint   `json:"id,string"`
	ParentID uint   `json:"parentId,string"`
}

// OpenCategoryGoodsGetRequest .
type OpenCategoryGoodsGetRequest struct {
	CategoryReq *CategoryReq `json:"req"`
}

// CategoryReq .
type CategoryReq struct {
	ParentID uint `json:"parentId,string"`
	Grade    uint `json:"grade,string"`
}
