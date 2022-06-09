

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAttachment ...
type TpAttachment struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    AdminID int `gorm:"column:admin_id;type:int(10);index:admin_id;" json:"admin_id"`
    UserID int `gorm:"column:user_id;type:int(10);index:user_id;" json:"user_id"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Mime string `gorm:"column:mime;type:varchar(125);" json:"mime"`
    Suffix string `gorm:"column:suffix;type:varchar(10);" json:"suffix"`
    Size float64 `gorm:"column:size;type:double(10,3);" json:"size"`
    Sha1 string `gorm:"column:sha1;type:varchar(40);" json:"sha1"`
    Storage string `gorm:"column:storage;type:varchar(40);" json:"storage"`
    URL string `gorm:"column:url;type:varchar(255);" json:"url"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAttachment) TableName() string {
	return "tp_attachment"
}

 
