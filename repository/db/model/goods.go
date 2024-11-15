package model

import "time"

type Goods struct {
	GoodsID     int       `gorm:"primaryKey;autoIncrement;column:goodsID"`
	GoodsName   string    `gorm:"type:varchar(30);not null;column:goodsName"`
	UserID      int       `gorm:"not null;column:userID"`
	Price       float64   `gorm:"type:decimal(10,2);not null;check:price >= 0;column:price"`
	CategoryID  int       `gorm:"not null;column:categoryID"`
	Details     string    `gorm:"type:text;column:details"`
	IsSold      bool      `gorm:"type:tinyint;not null;default:0;column:isSold"`
	GoodsImages string    `gorm:"type:varchar(256);column:goodsImages"`
	CreatedTime time.Time `gorm:"type:datetime;not null;column:createdTime"`
}

func (Goods) TableName() string {
	return "goods"
}
