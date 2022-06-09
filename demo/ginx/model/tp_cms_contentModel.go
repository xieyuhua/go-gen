

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpCmsContent ...
type TpCmsContent struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Title string `gorm:"column:title;type:varchar(55);" json:"title"`
    CategoryID int `gorm:"column:category_id;type:int(10);index:category_id;" json:"category_id"`
    Author string `gorm:"column:author;type:varchar(32);" json:"author"`
    Source string `gorm:"column:source;type:varchar(32);" json:"source"`
    IsRecommend int `gorm:"column:is_recommend;type:tinyint(1);" json:"is_recommend"`
    IsHot int `gorm:"column:is_hot;type:tinyint(1);" json:"is_hot"`
    IsTop int `gorm:"column:is_top;type:tinyint(1);" json:"is_top"`
    Tags string `gorm:"column:tags;type:varchar(255);" json:"tags"`
    Keyword string `gorm:"column:keyword;type:varchar(255);" json:"keyword"`
    Logo string `gorm:"column:logo;type:varchar(255);" json:"logo"`
    Attachment string `gorm:"column:attachment;type:varchar(255);" json:"attachment"`
    Description string `gorm:"column:description;type:varchar(255);" json:"description"`
    Content string `gorm:"column:content;type:text;" json:"content"`
    PublishTime time.Time `gorm:"column:publish_time;type:datetime;" json:"publish_time"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    IsShow int `gorm:"column:is_show;type:tinyint(1);" json:"is_show"`
    Click int `gorm:"column:click;type:int(10);" json:"click"`
    Video string `gorm:"column:video;type:varchar(255);" json:"video"`
    CreateUser int `gorm:"column:create_user;type:int(10);" json:"create_user"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    
}

// TableName ...
func ( TpCmsContent) TableName() string {
	return "tp_cms_content"
}

 
