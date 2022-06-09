

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopOrderGood ...
type TpShopOrderGood struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    OrderID int `gorm:"column:order_id;type:int(10);index:order_id;" json:"order_id"`
    GoodsID int `gorm:"column:goods_id;type:int(10);index:goods_id;" json:"goods_id"`
    GoodsName string `gorm:"column:goods_name;type:varchar(120);" json:"goods_name"`
    GoodsSn string `gorm:"column:goods_sn;type:varchar(60);" json:"goods_sn"`
    GoodsNum int `gorm:"column:goods_num;type:smallint(5);" json:"goods_num"`
    SalePrice float64 `gorm:"column:sale_price;type:decimal(10,2);" json:"sale_price"`
    MemberPrice float64 `gorm:"column:member_price;type:decimal(10,2);" json:"member_price"`
    GiveIntegral float64 `gorm:"column:give_integral;type:decimal(10,2);" json:"give_integral"`
    SpecKey string `gorm:"column:spec_key;type:varchar(128);" json:"spec_key"`
    SpecKeyName string `gorm:"column:spec_key_name;type:varchar(128);" json:"spec_key_name"`
    IsComment int `gorm:"column:is_comment;type:tinyint(1);" json:"is_comment"`
    PromType int `gorm:"column:prom_type;type:tinyint(1);" json:"prom_type"`
    PromID int `gorm:"column:prom_id;type:int(10);" json:"prom_id"`
    IsSend int `gorm:"column:is_send;type:tinyint(1);" json:"is_send"`
    DeliveryID int `gorm:"column:delivery_id;type:int(10);" json:"delivery_id"`
    GoodsSku string `gorm:"column:goods_sku;type:varchar(128);" json:"goods_sku"`
    ShippingPrice float64 `gorm:"column:shipping_price;type:decimal(10,2);" json:"shipping_price"`
    Weight int `gorm:"column:weight;type:int(10);" json:"weight"`
    Logo string `gorm:"column:logo;type:varchar(255);" json:"logo"`
    
}

// TableName ...
func ( TpShopOrderGood) TableName() string {
	return "tp_shop_order_goods"
}

 
