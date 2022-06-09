

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminRoleMenu ...
type TpAdminRoleMenu struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    RoleID int `gorm:"column:role_id;type:int(10);index:role_id;" json:"role_id"`
    MenuID int `gorm:"column:menu_id;type:int(10);index:menu_id;" json:"menu_id"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAdminRoleMenu) TableName() string {
	return "tp_admin_role_menu"
}

 
