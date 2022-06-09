

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminOperationLog ...
type TpAdminOperationLog struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    UserID int `gorm:"column:user_id;type:int(10);index:user_id;" json:"user_id"`
    Path string `gorm:"column:path;type:varchar(200);" json:"path"`
    Method string `gorm:"column:method;type:varchar(10);" json:"method"`
    IP string `gorm:"column:ip;type:varchar(120);" json:"ip"`
    Data string `gorm:"column:data;type:text;" json:"data"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    
}

// TableName ...
func ( TpAdminOperationLog) TableName() string {
	return "tp_admin_operation_log"
}

 
