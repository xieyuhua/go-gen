

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopCart ...
type TpShopCart struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    MemberID int `gorm:"column:member_id;type:int(10);" json:"member_id"`
    SessionID string `gorm:"column:session_id;type:varchar(128);index:session_id;" json:"session_id"`
    GoodsID int `gorm:"column:goods_id;type:int(10);" json:"goods_id"`
    GoodsSpu string `gorm:"column:goods_spu;type:varchar(60);" json:"goods_spu"`
    GoodsName string `gorm:"column:goods_name;type:varchar(120);" json:"goods_name"`
    MarketPrice float64 `gorm:"column:market_price;type:decimal(10,2);" json:"market_price"`
    SalePrice float64 `gorm:"column:sale_price;type:decimal(10,2);" json:"sale_price"`
    MemberPrice float64 `gorm:"column:member_price;type:decimal(10,2);" json:"member_price"`
    GoodsNum int `gorm:"column:goods_num;type:smallint(5);" json:"goods_num"`
    SpecKey string `gorm:"column:spec_key;type:varchar(64);" json:"spec_key"`
    SpecKeyName string `gorm:"column:spec_key_name;type:varchar(64);" json:"spec_key_name"`
    PromType int `gorm:"column:prom_type;type:tinyint(1);" json:"prom_type"`
    PromID int `gorm:"column:prom_id;type:int(11);" json:"prom_id"`
    GoodsSku string `gorm:"column:goods_sku;type:varchar(128);" json:"goods_sku"`
    Weight int `gorm:"column:weight;type:int(11);" json:"weight"`
    Tag string `gorm:"column:tag;type:varchar(20);" json:"tag"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    
}

// TableName ...
func ( TpShopCart) TableName() string {
	return "tp_shop_cart"
}

 
