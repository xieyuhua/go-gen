

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpExtension ...
type TpExtension struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Key string `gorm:"column:key;type:varchar(55);unique_index:key;" json:"key"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Version string `gorm:"column:version;type:varchar(55);" json:"version"`
    Title string `gorm:"column:title;type:varchar(55);" json:"title"`
    Description string `gorm:"column:description;type:varchar(50);" json:"description"`
    Tags string `gorm:"column:tags;type:varchar(255);" json:"tags"`
    Install int `gorm:"column:install;type:tinyint(1);" json:"install"`
    Enable int `gorm:"column:enable;type:tinyint(1);" json:"enable"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpExtension) TableName() string {
	return "tp_extension"
}

 
