

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpMemberLog ...
type TpMemberLog struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    MemberID int `gorm:"column:member_id;type:int(10);index:idx_member_id;" json:"member_id"`
    Desc string `gorm:"column:desc;type:varchar(500);" json:"desc"`
    Change string `gorm:"column:change;type:varchar(55);" json:"change"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    
}

// TableName ...
func ( TpMemberLog) TableName() string {
	return "tp_member_log"
}

 
