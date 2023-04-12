

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomDeployTask ...
type ZoomDeployTask struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    ApplyID int `gorm:"column:apply_id;type:int(10);index:idx_apply_id;" json:"apply_id"`
    GroupID int `gorm:"column:group_id;type:int(10);" json:"group_id"`
    Status int `gorm:"column:status;type:int(10);" json:"status"`
    Content string `gorm:"column:content;type:text;" json:"content"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomDeployTask) TableName() string {
	return "zoom_deploy_task"
}

 
