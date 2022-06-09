

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopGood ...
type TpShopGood struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Name string `gorm:"column:name;type:varchar(55);" json:"name"`
    CategoryID int `gorm:"column:category_id;type:int(10);index:category_id;" json:"category_id"`
    BrandID int `gorm:"column:brand_id;type:int(10);index:brand_id;" json:"brand_id"`
    SalePrice float64 `gorm:"column:sale_price;type:decimal(10,2);" json:"sale_price"`
    MarketPrice float64 `gorm:"column:market_price;type:decimal(10,2);" json:"market_price"`
    CostPrice float64 `gorm:"column:cost_price;type:decimal(10,2);" json:"cost_price"`
    Stock int `gorm:"column:stock;type:int(10);" json:"stock"`
    IsRecommend int `gorm:"column:is_recommend;type:tinyint(1);" json:"is_recommend"`
    IsHot int `gorm:"column:is_hot;type:tinyint(1);" json:"is_hot"`
    IsTop int `gorm:"column:is_top;type:tinyint(1);" json:"is_top"`
    Tags string `gorm:"column:tags;type:varchar(255);" json:"tags"`
    Keyword string `gorm:"column:keyword;type:varchar(255);" json:"keyword"`
    Spu string `gorm:"column:spu;type:varchar(100);" json:"spu"`
    Logo string `gorm:"column:logo;type:varchar(255);" json:"logo"`
    Images string `gorm:"column:images;type:varchar(1000);" json:"images"`
    Video string `gorm:"column:video;type:varchar(255);" json:"video"`
    Description string `gorm:"column:description;type:varchar(255);" json:"description"`
    Content string `gorm:"column:content;type:text;" json:"content"`
    PublishTime time.Time `gorm:"column:publish_time;type:datetime;" json:"publish_time"`
    Sort int `gorm:"column:sort;type:mediumint(8);" json:"sort"`
    OnSale int `gorm:"column:on_sale;type:tinyint(1);" json:"on_sale"`
    IsShow int `gorm:"column:is_show;type:tinyint(1);" json:"is_show"`
    Click int `gorm:"column:click;type:int(10);" json:"click"`
    Weight int `gorm:"column:weight;type:int(10);" json:"weight"`
    SalesSum int `gorm:"column:sales_sum;type:int(10);" json:"sales_sum"`
    CreateUser int `gorm:"column:create_user;type:int(10);" json:"create_user"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    DeleteTime time.Time `gorm:"column:delete_time;type:datetime;" json:"delete_time"`
    AdminGroupID int `gorm:"column:admin_group_id;type:int(10);index:idx_admin_group_id;" json:"admin_group_id"`
    ShareCommission float64 `gorm:"column:share_commission;type:decimal(10,2);" json:"share_commission"`
    
}

// TableName ...
func ( TpShopGood) TableName() string {
	return "tp_shop_goods"
}

 
