

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShippingCom ...
type TpShippingCom struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(55);unique_index:unq_name;" json:"name"`
    Code string `gorm:"column:code;type:varchar(55);unique_index:unq_code;" json:"code"`
    Enable int `gorm:"column:enable;type:tinyint(1);" json:"enable"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShippingCom) TableName() string {
	return "tp_shipping_com"
}

 
