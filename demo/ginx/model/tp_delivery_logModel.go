

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpDeliveryLog ...
type TpDeliveryLog struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    OrderID int `gorm:"column:order_id;type:int(10);index:order_id;" json:"order_id"`
    OrderSn string `gorm:"column:order_sn;type:varchar(64);" json:"order_sn"`
    MemberID int `gorm:"column:member_id;type:int(10);index:member_id;" json:"member_id"`
    AdminID int `gorm:"column:admin_id;type:int(10);" json:"admin_id"`
    Consignee string `gorm:"column:consignee;type:varchar(64);" json:"consignee"`
    Mobile string `gorm:"column:mobile;type:varchar(20);" json:"mobile"`
    Province int64 `gorm:"column:province;type:bigint(20);index:idx_province;" json:"province"`
    City int64 `gorm:"column:city;type:bigint(20);index:idx_city;" json:"city"`
    Area int64 `gorm:"column:area;type:bigint(20);index:idx_area;" json:"area"`
    Town int64 `gorm:"column:town;type:bigint(20);" json:"town"`
    Address string `gorm:"column:address;type:varchar(255);" json:"address"`
    ShippingCode string `gorm:"column:shipping_code;type:varchar(32);" json:"shipping_code"`
    ShippingName string `gorm:"column:shipping_name;type:varchar(64);" json:"shipping_name"`
    ShippingPrice float64 `gorm:"column:shipping_price;type:decimal(10,2);" json:"shipping_price"`
    InvoiceNo string `gorm:"column:invoice_no;type:varchar(255);" json:"invoice_no"`
    Note string `gorm:"column:note;type:varchar(500);" json:"note"`
    BestTime time.Time `gorm:"column:best_time;type:datetime;" json:"best_time"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    DeleteTime time.Time `gorm:"column:delete_time;type:datetime;" json:"delete_time"`
    AdminGroupID int `gorm:"column:admin_group_id;type:int(10);index:idx_admin_group_id;" json:"admin_group_id"`
    LastCheckTime time.Time `gorm:"column:last_check_time;type:datetime;" json:"last_check_time"`
    ShippingResult string `gorm:"column:shipping_result;type:text;" json:"shipping_result"`
    
}

// TableName ...
func ( TpDeliveryLog) TableName() string {
	return "tp_delivery_log"
}

 
