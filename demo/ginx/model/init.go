
package model
import (
	"ginx/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// DB 全局数据库连接
var DB *gorm.DB

// LinkDB 初始化连接db
func LinkDB() {

	// 数据连接  刚开机有可能数据未启动，所以循环直到连接成功
	for {
		log.Println("初始化数据库链接", config.GConf.Dbipport)
		if b, db := link(); b == true {
			DB = db
			AutoMigrate()
			return 
		}
		time.Sleep(1 * time.Second)
	}

}
func link() (bool, *gorm.DB) {
	var err error
	dbstr := config.GConf.Dbuser + ":" + config.GConf.Dbpass + "@tcp(" + config.GConf.Dbipport + ")/" + config.GConf.Dbname + "?charset=utf8&parseTime=true"
	tmpDB, err := gorm.Open("mysql", dbstr)
	if err != nil {
		return false, nil
	}
	//空闲
	tmpDB.DB().SetMaxIdleConns(20)
	//打开
	tmpDB.DB().SetMaxOpenConns(500)
	//超时
	tmpDB.DB().SetConnMaxLifetime(time.Second * 60)
	return true, tmpDB
}

// AutoMigrate ...
func AutoMigrate() {
	
	if er := DB.AutoMigrate(&SzyLoginUser{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminGroup{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminMenu{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminOperationLog{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminPermission{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminRole{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminRoleMenu{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminRolePermission{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAdminUser{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAgentLevel{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAreacity{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpAttachment{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpCmsBanner{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpCmsCategory{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpCmsContent{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpCmsPosition{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpCmsTag{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpDeliveryLog{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpExtdemo{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpExtension{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpMember{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpMemberAccount{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpMemberAddres{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpMemberLevel{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpMemberLog{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpMemberRecharge{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShippingArea{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShippingAreaItem{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShippingCom{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopBrand{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopCart{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopCategory{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopCouponList{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopCouponType{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopGood{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopGoodsAttr{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopGoodsSpec{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopGoodsSpecPrice{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopGoodsSpecValue{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopOrder{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopOrderAction{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopOrderGood{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpShopTag{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpSzyLog{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpSzyUser{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpTableRelation{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
	if er := DB.AutoMigrate(&TpWebConfig{}).Error; er != nil {
		log.Println("自动迁移错误:", er)
	}
    
}

