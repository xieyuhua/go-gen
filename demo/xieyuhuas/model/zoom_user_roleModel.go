

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomUserRole ...
type ZoomUserRole struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(100);" json:"name"`
    Privilege string `gorm:"column:privilege;type:varchar(2000);" json:"privilege"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomUserRole) TableName() string {
	return "zoom_user_role"
}

 
