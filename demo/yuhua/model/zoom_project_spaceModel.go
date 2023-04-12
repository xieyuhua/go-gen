

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomProjectSpace ...
type ZoomProjectSpace struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(100);" json:"name"`
    Description string `gorm:"column:description;type:varchar(2000);" json:"description"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomProjectSpace) TableName() string {
	return "zoom_project_space"
}

 
