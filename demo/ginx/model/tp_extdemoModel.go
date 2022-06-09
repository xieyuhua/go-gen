

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpExtdemo ...
type TpExtdemo struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpExtdemo) TableName() string {
	return "tp_extdemo"
}

 
