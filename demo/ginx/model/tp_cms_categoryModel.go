

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpCmsCategory ...
type TpCmsCategory struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    ParentID int `gorm:"column:parent_id;type:int(10);index:parent_id;" json:"parent_id"`
    Logo string `gorm:"column:logo;type:varchar(255);" json:"logo"`
    Type int `gorm:"column:type;type:tinyint(1);" json:"type"`
    Link string `gorm:"column:link;type:varchar(255);" json:"link"`
    Deep int `gorm:"column:deep;type:tinyint(2);" json:"deep"`
    Path string `gorm:"column:path;type:varchar(55);" json:"path"`
    IsShow int `gorm:"column:is_show;type:tinyint(1);" json:"is_show"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpCmsCategory) TableName() string {
	return "tp_cms_category"
}

 
