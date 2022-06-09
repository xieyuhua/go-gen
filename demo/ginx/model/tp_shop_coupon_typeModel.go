

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopCouponType ...
type TpShopCouponType struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(50);" json:"name"`
    CardType int `gorm:"column:card_type;type:smallint(1);index:idx_card_type;" json:"card_type"`
    Type int `gorm:"column:type;type:tinyint(1);index:idx_type;" json:"type"`
    Money float64 `gorm:"column:money;type:decimal(10,2);" json:"money"`
    Condition float64 `gorm:"column:condition;type:decimal(10,2);" json:"condition"`
    CreateNum int `gorm:"column:create_num;type:int(11);" json:"create_num"`
    SendNum int `gorm:"column:send_num;type:int(11);" json:"send_num"`
    UseNum int `gorm:"column:use_num;type:int(11);" json:"use_num"`
    GetNum int `gorm:"column:get_num;type:int(11);" json:"get_num"`
    DelNum int `gorm:"column:del_num;type:int(11);" json:"del_num"`
    SendStartTime time.Time `gorm:"column:send_start_time;type:date;" json:"send_start_time"`
    SendEndTime time.Time `gorm:"column:send_end_time;type:date;" json:"send_end_time"`
    UseStartTime time.Time `gorm:"column:use_start_time;type:date;" json:"use_start_time"`
    UseEndTime time.Time `gorm:"column:use_end_time;type:date;" json:"use_end_time"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    Status int `gorm:"column:status;type:smallint(1);index:idx_status;" json:"status"`
    AllyID int `gorm:"column:ally_id;type:int(11);" json:"ally_id"`
    ForGoods string `gorm:"column:for_goods;type:varchar(555);" json:"for_goods"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpShopCouponType) TableName() string {
	return "tp_shop_coupon_type"
}

 
