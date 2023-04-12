

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomServer ...
type ZoomServer struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    GroupID int `gorm:"column:group_id;type:int(10);" json:"group_id"`
    Name string `gorm:"column:name;type:varchar(100);" json:"name"`
    IP string `gorm:"column:ip;type:varchar(100);" json:"ip"`
    SSHPort int `gorm:"column:ssh_port;type:int(10);" json:"ssh_port"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomServer) TableName() string {
	return "zoom_server"
}

 
