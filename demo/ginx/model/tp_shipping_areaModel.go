

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShippingArea ...
type TpShippingArea struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    FirstWeight int `gorm:"column:first_weight;type:mediumint(8);" json:"first_weight"`
    Money float64 `gorm:"column:money;type:decimal(6,2);" json:"money"`
    SecondWeight int `gorm:"column:second_weight;type:mediumint(8);" json:"second_weight"`
    AddMoney float64 `gorm:"column:add_money;type:decimal(6,2);" json:"add_money"`
    ComCodes string `gorm:"column:com_codes;type:varchar(255);" json:"com_codes"`
    Enable int `gorm:"column:enable;type:tinyint(1);" json:"enable"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShippingArea) TableName() string {
	return "tp_shipping_area"
}

 
