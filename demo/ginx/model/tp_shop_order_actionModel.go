

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopOrderAction ...
type TpShopOrderAction struct {
    ID int `gorm:"column:id;type:mediumint(8);primary_key;" json:"id"`
    OrderID int `gorm:"column:order_id;type:mediumint(8);index:order_id;" json:"order_id"`
    OrderSn string `gorm:"column:order_sn;type:varchar(30);" json:"order_sn"`
    AdminID int `gorm:"column:admin_id;type:int(10);" json:"admin_id"`
    OrderStatus int `gorm:"column:order_status;type:tinyint(1);" json:"order_status"`
    ShippingStatus int `gorm:"column:shipping_status;type:tinyint(1);" json:"shipping_status"`
    PayStatus int `gorm:"column:pay_status;type:tinyint(1);" json:"pay_status"`
    ActionNote string `gorm:"column:action_note;type:varchar(255);" json:"action_note"`
    StatusDesc string `gorm:"column:status_desc;type:varchar(255);" json:"status_desc"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    
}

// TableName ...
func ( TpShopOrderAction) TableName() string {
	return "tp_shop_order_action"
}

 
