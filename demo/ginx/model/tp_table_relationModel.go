

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpTableRelation ...
type TpTableRelation struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    RelationName string `gorm:"column:relation_name;type:varchar(55);" json:"relation_name"`
    LocalTableName string `gorm:"column:local_table_name;type:varchar(55);" json:"local_table_name"`
    LocalKey string `gorm:"column:local_key;type:varchar(55);" json:"local_key"`
    ForeignTableName string `gorm:"column:foreign_table_name;type:varchar(55);" json:"foreign_table_name"`
    RelationType string `gorm:"column:relation_type;type:varchar(20);" json:"relation_type"`
    ForeignKey string `gorm:"column:foreign_key;type:varchar(55);" json:"foreign_key"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpTableRelation) TableName() string {
	return "tp_table_relation"
}

 
