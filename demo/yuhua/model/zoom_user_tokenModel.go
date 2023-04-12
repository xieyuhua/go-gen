

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomUserToken ...
type ZoomUserToken struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    UserID int `gorm:"column:user_id;type:int(10);unique_index:idx_user_id;" json:"user_id"`
    Token string `gorm:"column:token;type:varchar(100);" json:"token"`
    Expire int `gorm:"column:expire;type:int(10);" json:"expire"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomUserToken) TableName() string {
	return "zoom_user_token"
}

 
