

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAdminGroup ...
type TpAdminGroup struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    ParentID int `gorm:"column:parent_id;type:int(10);index:parent_id;" json:"parent_id"`
    Description string `gorm:"column:description;type:varchar(100);" json:"description"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Tags string `gorm:"column:tags;type:varchar(500);" json:"tags"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAdminGroup) TableName() string {
	return "tp_admin_group"
}

 
