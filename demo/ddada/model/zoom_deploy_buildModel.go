

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomDeployBuild ...
type ZoomDeployBuild struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    ApplyID int `gorm:"column:apply_id;type:int(10);unique_index:idx_apply_id;" json:"apply_id"`
    StartTime int `gorm:"column:start_time;type:int(10);" json:"start_time"`
    FinishTime int `gorm:"column:finish_time;type:int(10);" json:"finish_time"`
    Status int `gorm:"column:status;type:int(10);" json:"status"`
    Tar string `gorm:"column:tar;type:varchar(2000);" json:"tar"`
    Output string `gorm:"column:output;type:mediumtext;" json:"output"`
    Errmsg string `gorm:"column:errmsg;type:text;" json:"errmsg"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomDeployBuild) TableName() string {
	return "zoom_deploy_build"
}

 
