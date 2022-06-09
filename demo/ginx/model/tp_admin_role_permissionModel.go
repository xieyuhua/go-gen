

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminRolePermission ...
type TpAdminRolePermission struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    RoleID int `gorm:"column:role_id;type:int(10);index:role_id;" json:"role_id"`
    ControllerID int `gorm:"column:controller_id;type:int(10);" json:"controller_id"`
    PermissionID int `gorm:"column:permission_id;type:int(10);" json:"permission_id"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAdminRolePermission) TableName() string {
	return "tp_admin_role_permission"
}

 
