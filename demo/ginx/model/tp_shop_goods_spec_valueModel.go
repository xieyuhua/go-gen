

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopGoodsSpecValue ...
type TpShopGoodsSpecValue struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    SpecID int `gorm:"column:spec_id;type:int(10);index:model_id;" json:"spec_id"`
    GoodsID int `gorm:"column:goods_id;type:int(10);" json:"goods_id"`
    Value string `gorm:"column:value;type:varchar(55);" json:"value"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShopGoodsSpecValue) TableName() string {
	return "tp_shop_goods_spec_value"
}

 
