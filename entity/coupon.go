/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package entity

// JDUnionOpenCouponGiftGetTopLevel 优惠券领取情况查询接口
type JDUnionOpenCouponGiftGetTopLevel struct {
	JDUnionOpenCouponGiftGetResponse JDUnionOpenCouponGiftGetResponse `json:"jd_union_open_coupon_gift_get_responce"`
}

// JDUnionOpenCouponGiftGetResponse 优惠券领取情况查询接口
type JDUnionOpenCouponGiftGetResponse struct {
	Result string `json:"getResult"`
	Code   string `json:"code"`
}

// JDUnionOpenCouponGiftGetResult 优惠券领取情况查询接口
type JDUnionOpenCouponGiftGetResult struct {
	Code      int64       `json:"code"`
	Data      *CouponGift `json:"data"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestId"`
}

// CouponGift 优惠券领取情况查询接口
type CouponGift struct {
	GiftCouponKey string `json:"giftCouponKey"`
}

// JDUnionOpenCouponGiftStopTopLevel 优惠券领取情况查询接口
type JDUnionOpenCouponGiftStopTopLevel struct {
	JDUnionOpenCouponGiftStopResponse JDUnionOpenCouponGiftStopResponse `json:"jd_union_open_coupon_gift_stop_responce"`
}

// JDUnionOpenCouponGiftStopResponse 优惠券领取情况查询接口
type JDUnionOpenCouponGiftStopResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

// JDUnionOpenCouponGiftStopResult 优惠券领取情况查询接口
type JDUnionOpenCouponGiftStopResult struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}
