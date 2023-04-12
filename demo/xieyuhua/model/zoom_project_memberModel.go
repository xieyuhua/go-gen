

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomProjectMember ...
type ZoomProjectMember struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    SpaceID int `gorm:"column:space_id;type:int(10);" json:"space_id"`
    UserID int `gorm:"column:user_id;type:int(10);" json:"user_id"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomProjectMember) TableName() string {
	return "zoom_project_member"
}

 
