

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpMemberAddres ...
type TpMemberAddres struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    MemberID int `gorm:"column:member_id;type:mediumint(8);index:idx_member_id;" json:"member_id"`
    Consignee string `gorm:"column:consignee;type:varchar(20);" json:"consignee"`
    Province int64 `gorm:"column:province;type:bigint(20);" json:"province"`
    City int64 `gorm:"column:city;type:bigint(20);" json:"city"`
    Area int64 `gorm:"column:area;type:bigint(20);" json:"area"`
    Town int64 `gorm:"column:town;type:bigint(20);" json:"town"`
    Address string `gorm:"column:address;type:varchar(255);" json:"address"`
    Mobile string `gorm:"column:mobile;type:varchar(60);" json:"mobile"`
    IsDefault int `gorm:"column:is_default;type:tinyint(1);" json:"is_default"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime string `gorm:"column:update_time;type:varchar(55);" json:"update_time"`
    
}

// TableName ...
func ( TpMemberAddres) TableName() string {
	return "tp_member_address"
}

 
