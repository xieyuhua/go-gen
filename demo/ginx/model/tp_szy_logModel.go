

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpSzyLog ...
type TpSzyLog struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Level int `gorm:"column:level;type:int(11);" json:"level"`
    Category string `gorm:"column:category;type:varchar(100);index:log_time;" json:"category"`
    AddTime int `gorm:"column:add_time;type:int(10);" json:"add_time"`
    UserID int `gorm:"column:user_id;type:int(11);" json:"user_id"`
    ShopID int `gorm:"column:shop_id;type:int(11);" json:"shop_id"`
    StoreID int `gorm:"column:store_id;type:int(11);" json:"store_id"`
    AppID string `gorm:"column:app_id;type:varchar(100);" json:"app_id"`
    Content string `gorm:"column:content;type:text;" json:"content"`
    From int `gorm:"column:from;type:tinyint(1);" json:"from"`
    IP string `gorm:"column:ip;type:varchar(20);" json:"ip"`
    
}

// TableName ...
func ( TpSzyLog) TableName() string {
	return "tp_szy_log"
}

 
