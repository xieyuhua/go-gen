

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopGoodsSpec ...
type TpShopGoodsSpec struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Value string `gorm:"column:value;type:varchar(1000);" json:"value"`
    GoodsID int `gorm:"column:goods_id;type:int(10);index:goods_id;" json:"goods_id"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShopGoodsSpec) TableName() string {
	return "tp_shop_goods_spec"
}

 
