

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpWebConfig ...
type TpWebConfig struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Key string `gorm:"column:key;type:varchar(55);unique_index:key;" json:"key"`
    File string `gorm:"column:file;type:varchar(255);" json:"file"`
    Title string `gorm:"column:title;type:varchar(55);" json:"title"`
    Config string `gorm:"column:config;type:text;" json:"config"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpWebConfig) TableName() string {
	return "tp_web_config"
}

 
