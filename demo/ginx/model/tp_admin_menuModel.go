

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminMenu ...
type TpAdminMenu struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    ParentID int `gorm:"column:parent_id;type:int(10);index:parent_id;" json:"parent_id"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    Title string `gorm:"column:title;type:varchar(55);" json:"title"`
    URL string `gorm:"column:url;type:varchar(100);" json:"url"`
    Icon string `gorm:"column:icon;type:varchar(50);" json:"icon"`
    Module string `gorm:"column:module;type:varchar(55);" json:"module"`
    Enable int `gorm:"column:enable;type:tinyint(1);" json:"enable"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAdminMenu) TableName() string {
	return "tp_admin_menu"
}

 
