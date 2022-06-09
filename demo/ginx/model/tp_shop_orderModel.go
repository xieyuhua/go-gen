

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpShopOrder ...
type TpShopOrder struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    OrderSn string `gorm:"column:order_sn;type:varchar(30);unique_index:order_sn;" json:"order_sn"`
    MemberID int `gorm:"column:member_id;type:mediumint(8);index:member_id;" json:"member_id"`
    Consignee string `gorm:"column:consignee;type:varchar(20);" json:"consignee"`
    Province int64 `gorm:"column:province;type:bigint(20);index:province;" json:"province"`
    City int64 `gorm:"column:city;type:bigint(20);index:city;" json:"city"`
    Area int64 `gorm:"column:area;type:bigint(20);index:area;" json:"area"`
    Town int64 `gorm:"column:town;type:bigint(20);" json:"town"`
    Address string `gorm:"column:address;type:varchar(255);" json:"address"`
    Mobile string `gorm:"column:mobile;type:varchar(60);index:mobile;" json:"mobile"`
    ShippingCode string `gorm:"column:shipping_code;type:varchar(32);index:shipping_code;" json:"shipping_code"`
    ShippingName string `gorm:"column:shipping_name;type:varchar(120);" json:"shipping_name"`
    InvoiceTitle string `gorm:"column:invoice_title;type:varchar(256);" json:"invoice_title"`
    GoodsPrice float64 `gorm:"column:goods_price;type:decimal(10,2);" json:"goods_price"`
    ShippingPrice float64 `gorm:"column:shipping_price;type:decimal(10,2);" json:"shipping_price"`
    UseMoney float64 `gorm:"column:use_money;type:decimal(10,2);" json:"use_money"`
    UsePoints float64 `gorm:"column:use_points;type:decimal(10,2);" json:"use_points"`
    PointsMoney float64 `gorm:"column:points_money;type:decimal(10,2);" json:"points_money"`
    CouponCardType int `gorm:"column:coupon_card_type;type:tinyint(1);" json:"coupon_card_type"`
    CouponPrice float64 `gorm:"column:coupon_price;type:decimal(10,2);" json:"coupon_price"`
    OrderAmount float64 `gorm:"column:order_amount;type:decimal(10,2);" json:"order_amount"`
    TotalAmount float64 `gorm:"column:total_amount;type:decimal(10,2);" json:"total_amount"`
    OrderPromID int `gorm:"column:order_prom_id;type:smallint(6);" json:"order_prom_id"`
    OrderPromAmount float64 `gorm:"column:order_prom_amount;type:decimal(10,2);" json:"order_prom_amount"`
    Discount float64 `gorm:"column:discount;type:decimal(10,2);" json:"discount"`
    UserNote string `gorm:"column:user_note;type:varchar(255);" json:"user_note"`
    AdminNote string `gorm:"column:admin_note;type:varchar(255);" json:"admin_note"`
    HasDistribut int `gorm:"column:has_distribut;type:tinyint(1);" json:"has_distribut"`
    OrderStatus int `gorm:"column:order_status;type:tinyint(1);index:order_status;" json:"order_status"`
    PayCode string `gorm:"column:pay_code;type:varchar(32);index:pay_code;" json:"pay_code"`
    PayStatus int `gorm:"column:pay_status;type:tinyint(1);index:pay_status;" json:"pay_status"`
    PayTime time.Time `gorm:"column:pay_time;type:datetime;" json:"pay_time"`
    ShippingStatus int `gorm:"column:shipping_status;type:tinyint(1);index:shipping_status;" json:"shipping_status"`
    TransactionID string `gorm:"column:transaction_id;type:varchar(80);" json:"transaction_id"`
    ShippingTime time.Time `gorm:"column:shipping_time;type:datetime;" json:"shipping_time"`
    ConfirmTime time.Time `gorm:"column:confirm_time;type:datetime;" json:"confirm_time"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;type:datetime;" json:"update_time"`
    DeleteTime time.Time `gorm:"column:delete_time;type:datetime;" json:"delete_time"`
    AddressID int `gorm:"column:address_id;type:int(10);" json:"address_id"`
    UseCommission float64 `gorm:"column:use_commission;type:decimal(10,2);" json:"use_commission"`
    
}

// TableName ...
func ( TpShopOrder) TableName() string {
	return "tp_shop_order"
}

 
