

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpAgentLevel ...
type TpAgentLevel struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    Level int `gorm:"column:level;type:tinyint(3);" json:"level"`
    Description string `gorm:"column:description;type:varchar(255);" json:"description"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpAgentLevel) TableName() string {
	return "tp_agent_level"
}

 
