

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomUser ...
type ZoomUser struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    RoleID int `gorm:"column:role_id;type:int(10);" json:"role_id"`
    Username string `gorm:"column:username;type:varchar(20);index:idx_username;" json:"username"`
    Password string `gorm:"column:password;type:char(32);" json:"password"`
    Salt string `gorm:"column:salt;type:char(10);" json:"salt"`
    Truename string `gorm:"column:truename;type:varchar(10);" json:"truename"`
    Mobile string `gorm:"column:mobile;type:varchar(20);" json:"mobile"`
    Email string `gorm:"column:email;type:varchar(500);" json:"email"`
    Status int `gorm:"column:status;type:int(10);" json:"status"`
    LastLoginTime int `gorm:"column:last_login_time;type:int(10);" json:"last_login_time"`
    LastLoginIP string `gorm:"column:last_login_ip;type:varchar(50);" json:"last_login_ip"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomUser) TableName() string {
	return "zoom_user"
}

 
