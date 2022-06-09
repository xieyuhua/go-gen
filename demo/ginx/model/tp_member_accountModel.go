

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpMemberAccount ...
type TpMemberAccount struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    MemberID int `gorm:"column:member_id;type:int(10);index:idx_member_id;" json:"member_id"`
    AdminID int `gorm:"column:admin_id;type:int(10);" json:"admin_id"`
    Points float64 `gorm:"column:points;type:decimal(10,2);" json:"points"`
    Money float64 `gorm:"column:money;type:decimal(10,2);" json:"money"`
    Commission float64 `gorm:"column:commission;type:decimal(10,2);" json:"commission"`
    Remark string `gorm:"column:remark;type:varchar(255);" json:"remark"`
    OrderID int `gorm:"column:order_id;type:int(10);" json:"order_id"`
    Tag string `gorm:"column:tag;type:varchar(55);" json:"tag"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    
}

// TableName ...
func ( TpMemberAccount) TableName() string {
	return "tp_member_account"
}

 
