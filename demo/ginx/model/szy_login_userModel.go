

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// SzyLoginUser ...
type SzyLoginUser struct {
    UserID int `gorm:"column:user_id;type:int(11);" json:"user_id"`
    UserName string `gorm:"column:user_name;type:varchar(60);" json:"user_name"`
    Mobile string `gorm:"column:mobile;type:varchar(20);" json:"mobile"`
    RegTime int `gorm:"column:reg_time;type:int(11);" json:"reg_time"`
    RegIP string `gorm:"column:reg_ip;type:varchar(40);" json:"reg_ip"`
    LastTime int `gorm:"column:last_time;type:int(11);" json:"last_time"`
    LastIP string `gorm:"column:last_ip;type:varchar(40);" json:"last_ip"`
    
}

// TableName ...
func ( SzyLoginUser) TableName() string {
	return "szy_login_user"
}

 
