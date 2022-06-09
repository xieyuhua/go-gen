

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpCmsPosition ...
type TpCmsPosition struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Logo string `gorm:"column:logo;type:varchar(255);" json:"logo"`
    Type int `gorm:"column:type;type:tinyint(1);" json:"type"`
    IsShow int `gorm:"column:is_show;type:tinyint(1);" json:"is_show"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    StartTime time.Time `gorm:"column:start_time;type:datetime;" json:"start_time"`
    EndTime time.Time `gorm:"column:end_time;type:datetime;" json:"end_time"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpCmsPosition) TableName() string {
	return "tp_cms_position"
}

 
