

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopCouponList ...
type TpShopCouponList struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    CouponTypeID int `gorm:"column:coupon_type_id;type:int(10);index:cid;" json:"coupon_type_id"`
    CardType int `gorm:"column:card_type;type:smallint(1);index:ctype;" json:"card_type"`
    Type int `gorm:"column:type;type:tinyint(1);" json:"type"`
    MemberID int `gorm:"column:member_id;type:int(10);" json:"member_id"`
    OrderID int `gorm:"column:order_id;type:int(8);" json:"order_id"`
    UseTime time.Time `gorm:"column:use_time;type:datetime;" json:"use_time"`
    Code string `gorm:"column:code;type:varchar(20);index:sn;" json:"code"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    GetTime time.Time `gorm:"column:get_time;type:date;" json:"get_time"`
    No int `gorm:"column:no_;type:int(10);" json:"no_"`
    Status int `gorm:"column:status;type:smallint(1);index:status;" json:"status"`
    DeleteTime time.Time `gorm:"column:delete_time;type:datetime;" json:"delete_time"`
    
}

// TableName ...
func ( TpShopCouponList) TableName() string {
	return "tp_shop_coupon_list"
}

 
