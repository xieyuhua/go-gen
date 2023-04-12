

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomServerGroup ...
type ZoomServerGroup struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(100);" json:"name"`
    Ctime int `gorm:"column:ctime;type:int(11);" json:"ctime"`
    
}

// TableName ...
func ( ZoomServerGroup) TableName() string {
	return "zoom_server_group"
}

 
