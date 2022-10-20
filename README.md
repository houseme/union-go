# UNION-GO

[![Go Doc](https://godoc.org/github.com/houseme/union-go?status.svg)](https://godoc.org/github.com/houseme/union-go)
[![GoFrame CI](https://github.com/houseme/union-go/actions/workflows/go.yml/badge.svg)](https://github.com/houseme/union-go/actions/workflows/go.yml)
[![Go Report](https://goreportcard.com/badge/github.com/houseme/union-go?v=1)](https://goreportcard.com/report/github.com/houseme/union-go)
[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/houseme/union-go)
[![License](https://img.shields.io/github/license/houseme/union-go.svg?style=flat)](https://github.com/houseme/union-go)

JD.com Alliance Open platform API Interface Golang SDK (Development Kit)

JD.com Alliance optimizes the core competence of the alliance, such as commodity pool, chain switching ability, effect query, marketing tools, etc. API.

# 京东联盟开放平台API（golang）

> https://union.jd.com/openplatform/api

## APIs

| 功能 | API列表 |
| :--- | :--- |
| 京粉精选商品查询接口 | jd.union.open.goods.jingfen.query |
| 订单查询接口 | jd.union.open.order.query |
| 关键词商品查询接口【申请】| jd.union.open.goods.query |
| 网站/APP获取推广链接接口 | jd.union.open.promotion.common.get |
| 根据skuid查询商品信息接口 | jd.union.open.goods.promotiongoodsinfo.query |
| 优惠券领取情况查询接口【申请】 | jd.union.open.coupon.query |
| 社交媒体获取推广链接接口【申请】 | jd.union.open.promotion.bysubunionid.get |
| 工具商获取推广链接接口【申请】 | jd.union.open.promotion.byunionid.get |
| 查询推广位【申请】 | jd.union.open.position.query |
| 创建推广位【申请】 | jd.union.open.position.create |
| 获取PID【申请】 | jd.union.open.user.pid.get |
| 秒杀商品查询接口【申请】 | jd.union.open.goods.seckill.query |
| 学生价商品查询接口【申请】 | jd.union.open.goods.stuprice.query |
| 商品类目查询接口 | jd.union.open.category.goods.get |
| 商品详情查询接口【申请】 | jd.union.open.goods.bigfield.query |
| 奖励订单查询接口【申请】 | jd.union.open.order.bonus.query |
| 礼金停止【申请】 | jd.union.open.coupon.gift.stop |
| 礼金创建【申请】 | jd.union.open.coupon.gift.get |
| 礼金效果数据 | jd.union.open.statistics.giftcoupon.query |
| 活动查询接口 | jd.union.open.activity.query |
| 订单行查询接口 | jd.union.open.order.row.query |

# Installation
Enter your repo. directory and execute following command:

```bash
go get -u -v github.com/houseme/union-go
```

# Limitation
```
golang version >= 1.18
```

# License

`UNION-JD-GO` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.