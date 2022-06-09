

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopGoodsSpecPrice ...
type TpShopGoodsSpecPrice struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    GoodsID int `gorm:"column:goods_id;type:int(10);index:goods_id;" json:"goods_id"`
    SpecIds string `gorm:"column:spec_ids;type:varchar(55);" json:"spec_ids"`
    SpecKey string `gorm:"column:spec_key;type:varchar(128);unique_index:spec_key;" json:"spec_key"`
    SpecKeyName string `gorm:"column:spec_key_name;type:varchar(128);" json:"spec_key_name"`
    Data string `gorm:"column:data;type:varchar(255);" json:"data"`
    Sku string `gorm:"column:sku;type:varchar(55);" json:"sku"`
    SalePrice float64 `gorm:"column:sale_price;type:decimal(10,2);" json:"sale_price"`
    Stock int `gorm:"column:stock;type:int(10);" json:"stock"`
    Img string `gorm:"column:img;type:varchar(255);" json:"img"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShopGoodsSpecPrice) TableName() string {
	return "tp_shop_goods_spec_price"
}

 
