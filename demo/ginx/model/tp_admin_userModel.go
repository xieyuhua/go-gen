

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminUser ...
type TpAdminUser struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    RoleID int `gorm:"column:role_id;type:int(10);index:role_id;" json:"role_id"`
    Username string `gorm:"column:username;type:varchar(30);unique_index:username;" json:"username"`
    Password string `gorm:"column:password;type:varchar(100);" json:"password"`
    Salt string `gorm:"column:salt;type:varchar(10);" json:"salt"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Avatar string `gorm:"column:avatar;type:varchar(120);" json:"avatar"`
    Phone string `gorm:"column:phone;type:varchar(20);" json:"phone"`
    Email string `gorm:"column:email;type:varchar(55);" json:"email"`
    Errors int `gorm:"column:errors;type:int(10);" json:"errors"`
    Enable int `gorm:"column:enable;type:tinyint(1);" json:"enable"`
    Tags string `gorm:"column:tags;type:varchar(500);" json:"tags"`
    GroupID int `gorm:"column:group_id;type:int(10);index:group_id;" json:"group_id"`
    LoginTime time.Time `gorm:"column:login_time;type:datetime;" json:"login_time"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAdminUser) TableName() string {
	return "tp_admin_user"
}

 
