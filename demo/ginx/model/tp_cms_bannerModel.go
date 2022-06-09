

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpCmsBanner ...
type TpCmsBanner struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Title string `gorm:"column:title;type:varchar(55);" json:"title"`
    PositionID int `gorm:"column:position_id;type:mediumint(8);index:position_id;" json:"position_id"`
    Description string `gorm:"column:description;type:varchar(255);" json:"description"`
    Image string `gorm:"column:image;type:varchar(255);" json:"image"`
    Link string `gorm:"column:link;type:varchar(255);" json:"link"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    IsShow int `gorm:"column:is_show;type:tinyint(1);" json:"is_show"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpCmsBanner) TableName() string {
	return "tp_cms_banner"
}

 
