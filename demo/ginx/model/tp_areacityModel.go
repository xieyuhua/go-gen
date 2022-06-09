

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAreacity ...
type TpAreacity struct {
    ID int64 `gorm:"column:id;type:bigint(15);primary_key;" json:"id"`
    Pid int64 `gorm:"column:pid;type:bigint(15);index:pid;" json:"pid"`
    Deep int `gorm:"column:deep;type:tinyint(1);" json:"deep"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    PinyinPrefix string `gorm:"column:pinyin_prefix;type:varchar(55);index:pinyin_prefix;" json:"pinyin_prefix"`
    Pinyin string `gorm:"column:pinyin;type:varchar(50);" json:"pinyin"`
    ExtID int64 `gorm:"column:ext_id;type:bigint(15);" json:"ext_id"`
    ExtName string `gorm:"column:ext_name;type:varchar(55);" json:"ext_name"`
    
}

// TableName ...
func ( TpAreacity) TableName() string {
	return "tp_areacity"
}

 
