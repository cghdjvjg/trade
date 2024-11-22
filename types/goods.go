package types

import "time"

type Goods struct {
	GoodsID     int       `json:"goodsID"`     // 商品ID
	GoodsName   string    `json:"goodsName"`   // 商品名称
	UserID      int       `json:"userID"`      // 用户ID
	Price       float64   `json:"price"`       // 价格
	CategoryID  int       `json:"categoryID"`  // 分类ID
	Details     string    `json:"details"`     // 商品详情
	IsSold      int       `json:"isSold"`      // 是否已售：0 未售，1 已售
	GoodsImages string    `json:"goodsImages"` // 商品图片
	CreatedTime time.Time `json:"createdTime"` // 创建时间
}

type GoodsListResp struct {
	GoodsList []Goods `json:"goodsList"` // 商品列表
	Total     int     `json:"total"`     // 总记录数
	PageNum   int     `json:"pageNum"`   // 当前页码
}

type ShowGoodsReq struct {
	SearchQuery string  `form:"searchQuery" json:"searchQuery"` // 商品名称模糊查询
	CategoryID  int     `form:"categoryID" json:"categoryID"`   // 商品分类
	PriceMin    float64 `form:"priceMin" json:"priceMin"`       // 最低价格
	PriceMax    float64 `form:"priceMax" json:"priceMax"`       // 最高价格
	IsSold      int     `form:"isSold" json:"isSold"`           // 商品是否售出（0:未售，1:已售）
	PageNum     int     `form:"pageNum" json:"pageNum"`         // 当前页码
	PageSize    int     `form:"pageSize" json:"pageSize"`       // 每页记录数
}

type UpdateGoodsReq struct {
	GoodsID     int     `json:"goodsID"`     // 商品ID
	GoodsName   string  `json:"goodsName"`   // 商品名称
	Price       float64 `json:"price"`       // 价格
	CategoryID  int     `json:"categoryID"`  // 分类ID
	Details     string  `json:"details"`     // 商品详情
	IsSold      int     `json:"isSold"`      // 是否售出
	GoodsImages string  `json:"goodsImages"` // 商品图片
}

type CreateGoodsReq struct {
	UserID      int     `json:"userID"`      // 用户ID
	GoodsName   string  `json:"goodsName"`   // 商品名称
	Price       float64 `json:"price"`       // 价格
	CategoryID  int     `json:"categoryID"`  // 商品分类
	Details     string  `json:"details"`     // 商品详情
	GoodsImages string  `json:"goodsImages"` // 商品图片
}