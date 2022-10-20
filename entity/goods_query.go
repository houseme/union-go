/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package entity

// JdUnionOpenGoodsPromotionGoodsInfoQueryResponseTop is a top struct.
type JdUnionOpenGoodsPromotionGoodsInfoQueryResponseTop struct {
	JdUnionOpenGoodsPromotionGoodsInfoQueryResponse *JdUnionOpenGoodsPromotionGoodsInfoQueryResponse `json:"jd_union_open_goods_promotiongoodsinfo_query_response"`
}

// JdUnionOpenGoodsPromotionGoodsInfoQueryResponse is a response struct.
type JdUnionOpenGoodsPromotionGoodsInfoQueryResponse struct {
	QueryResult *JdUnionOpenGoodsPromotionGoodsInfoQueryQueryResult `json:"queryResult"`
}

// JdUnionOpenGoodsPromotionGoodsInfoQueryQueryResult is a queryResult struct.
type JdUnionOpenGoodsPromotionGoodsInfoQueryQueryResult struct {
	Code    string                                                  `json:"code"`
	Data    *JdUnionOpenGoodsPromotionGoodsInfoQueryQueryResultData `json:"data"`
	Message string                                                  `json:"message"`
}

// JdUnionOpenGoodsPromotionGoodsInfoQueryQueryResultData is a data struct.
type JdUnionOpenGoodsPromotionGoodsInfoQueryQueryResultData struct {
	PromotionGoodsResp *PromotionGoodsResp `json:"promotionGoodsResp"`
}

// PromotionGoodsResp promotion goods resp.
type PromotionGoodsResp struct {
	UnitPrice         string `json:"unitPrice"`
	MaterialURL       string `json:"materialUrl"`
	EndDate           string `json:"endDate"`
	IsFreeFreightRisk string `json:"isFreeFreightRisk"`
	IsFreeShipping    string `json:"isFreeShipping"`
	CommisionRatioWl  string `json:"commisionRatioWl"`
	CommisionRatioPc  string `json:"commisionRatioPc"`
	ImgURL            string `json:"imgUrl"`
	Vid               string `json:"vid"`
	CidName           string `json:"cidName"`
	WlUnitPrice       string `json:"wlUnitPrice"`
	Cid2Name          string `json:"cid2Name"`
	IsSeckill         string `json:"isSeckill"`
	Cid2              string `json:"cid2"`
	Cid3Name          string `json:"cid3Name"`
	InOrderCount      string `json:"inOrderCount"`
	Cid3              string `json:"cid3"`
	ShopID            string `json:"shopId"`
	IsJdSale          string `json:"isJdSale"`
	GoodsName         string `json:"goodsName"`
	SkuID             string `json:"skuId"`
	StartDate         string `json:"startDate"`
	Cid               string `json:"cid"`
}

// JdUnionOpenGoodsPromotionGoodsInfoQueryRequest is a request struct.
type JdUnionOpenGoodsPromotionGoodsInfoQueryRequest struct {
	// skuIds	必填	商品skuId，支持批量，最多不超过100个，英文逗号分隔
	SkuIds string `json:"skuIds" dc:"开发示例如param_json={'skuIds':'5225346,7275691'}"`
}
