

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShippingAreaItem ...
type TpShippingAreaItem struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    ComCodes string `gorm:"column:com_codes;type:varchar(255);" json:"com_codes"`
    AreaID int `gorm:"column:area_id;type:int(10);" json:"area_id"`
    Province int64 `gorm:"column:province;type:bigint(20);" json:"province"`
    City int64 `gorm:"column:city;type:bigint(20);" json:"city"`
    Area int64 `gorm:"column:area;type:bigint(20);" json:"area"`
    Town int64 `gorm:"column:town;type:bigint(20);" json:"town"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShippingAreaItem) TableName() string {
	return "tp_shipping_area_item"
}

 
