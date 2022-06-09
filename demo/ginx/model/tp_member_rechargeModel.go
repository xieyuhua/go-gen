

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpMemberRecharge ...
type TpMemberRecharge struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    MemberID int `gorm:"column:member_id;type:int(10);" json:"member_id"`
    OrderSn string `gorm:"column:order_sn;type:varchar(30);" json:"order_sn"`
    Account float64 `gorm:"column:account;type:decimal(10,2);" json:"account"`
    PayTime time.Time `gorm:"column:pay_time;type:datetime;" json:"pay_time"`
    PayCode string `gorm:"column:pay_code;type:varchar(20);" json:"pay_code"`
    PayStatus int `gorm:"column:pay_status;type:tinyint(1);" json:"pay_status"`
    PkgID int `gorm:"column:pkg_id;type:int(10);" json:"pkg_id"`
    Remark string `gorm:"column:remark;type:varchar(255);" json:"remark"`
    TransactionID string `gorm:"column:transaction_id;type:varchar(80);" json:"transaction_id"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UseChildAccount int `gorm:"column:use_child_account;type:tinyint(1);" json:"use_child_account"`
    GoodsMethod int `gorm:"column:goods_method;type:tinyint(1);" json:"goods_method"`
    BuyNum int `gorm:"column:buy_num;type:tinyint(1);" json:"buy_num"`
    FirstLeader int `gorm:"column:first_leader;type:int(10);" json:"first_leader"`
    
}

// TableName ...
func ( TpMemberRecharge) TableName() string {
	return "tp_member_recharge"
}

 
