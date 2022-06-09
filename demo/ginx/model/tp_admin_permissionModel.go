

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminPermission ...
type TpAdminPermission struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    ModuleName string `gorm:"column:module_name;type:varchar(55);" json:"module_name"`
    Controller string `gorm:"column:controller;type:varchar(100);index:controller;" json:"controller"`
    Action string `gorm:"column:action;type:varchar(55);" json:"action"`
    URL string `gorm:"column:url;type:varchar(100);" json:"url"`
    ActionName string `gorm:"column:action_name;type:varchar(25);" json:"action_name"`
    ActionType int `gorm:"column:action_type;type:tinyint(1);" json:"action_type"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAdminPermission) TableName() string {
	return "tp_admin_permission"
}

 
