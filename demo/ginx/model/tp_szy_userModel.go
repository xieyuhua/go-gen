

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpSzyUser ...
type TpSzyUser struct {
    ID int `gorm:"column:id;type:int(11);primary_key;" json:"id"`
    UserName string `gorm:"column:user_name;type:varchar(60);unique_index:user_name;" json:"user_name"`
    Mobile string `gorm:"column:mobile;type:varchar(20);" json:"mobile"`
    Email string `gorm:"column:email;type:varchar(60);" json:"email"`
    IsSeller int `gorm:"column:is_seller;type:tinyint(1);index:is_seller;" json:"is_seller"`
    ShopID int `gorm:"column:shop_id;type:int(11);index:shop_id;" json:"shop_id"`
    StoreID int `gorm:"column:store_id;type:int(11);index:store_id;" json:"store_id"`
    MultiStoreID int `gorm:"column:multi_store_id;type:int(11);" json:"multi_store_id"`
    InitPassword string `gorm:"column:init_password;type:varchar(100);" json:"init_password"`
    ComstoreID int `gorm:"column:comstore_id;type:int(11);index:comstore_id;" json:"comstore_id"`
    Password string `gorm:"column:password;type:varchar(255);" json:"password"`
    PasswordResetToken string `gorm:"column:password_reset_token;type:varchar(255);" json:"password_reset_token"`
    Salt string `gorm:"column:salt;type:char(10);" json:"salt"`
    Nickname string `gorm:"column:nickname;type:varchar(200);" json:"nickname"`
    Sex int `gorm:"column:sex;type:tinyint(1);" json:"sex"`
    Birthday int `gorm:"column:birthday;type:int(11);" json:"birthday"`
    AddressNow string `gorm:"column:address_now;type:varchar(60);" json:"address_now"`
    DetailAddress string `gorm:"column:detail_address;type:varchar(255);" json:"detail_address"`
    Headimg string `gorm:"column:headimg;type:varchar(255);" json:"headimg"`
    Faceimg1 string `gorm:"column:faceimg1;type:text;" json:"faceimg1"`
    Faceimg2 string `gorm:"column:faceimg2;type:text;" json:"faceimg2"`
    UserMoney float64 `gorm:"column:user_money;type:decimal(10,2);" json:"user_money"`
    UserMoneyLimit float64 `gorm:"column:user_money_limit;type:decimal(10,2);" json:"user_money_limit"`
    FrozenMoney float64 `gorm:"column:frozen_money;type:decimal(10,2);" json:"frozen_money"`
    PayPoint string `gorm:"column:pay_point;type:varchar(255);" json:"pay_point"`
    FrozenPoint string `gorm:"column:frozen_point;type:varchar(255);" json:"frozen_point"`
    RankPoint int `gorm:"column:rank_point;type:int(11);" json:"rank_point"`
    AddressID int `gorm:"column:address_id;type:int(11);" json:"address_id"`
    RankID int `gorm:"column:rank_id;type:int(11);" json:"rank_id"`
    RankStartTime int `gorm:"column:rank_start_time;type:int(11);" json:"rank_start_time"`
    RankEndTime int `gorm:"column:rank_end_time;type:int(11);" json:"rank_end_time"`
    MobileValidated int `gorm:"column:mobile_validated;type:tinyint(4);" json:"mobile_validated"`
    EmailValidated int `gorm:"column:email_validated;type:tinyint(4);" json:"email_validated"`
    RegTime int `gorm:"column:reg_time;type:int(11);" json:"reg_time"`
    RegIP string `gorm:"column:reg_ip;type:varchar(40);" json:"reg_ip"`
    LastTime int `gorm:"column:last_time;type:int(11);" json:"last_time"`
    LastIP string `gorm:"column:last_ip;type:varchar(40);" json:"last_ip"`
    VisitCount int `gorm:"column:visit_count;type:int(11);" json:"visit_count"`
    MobileSupplier string `gorm:"column:mobile_supplier;type:varchar(255);" json:"mobile_supplier"`
    MobileProvince string `gorm:"column:mobile_province;type:varchar(255);" json:"mobile_province"`
    MobileCity string `gorm:"column:mobile_city;type:varchar(255);" json:"mobile_city"`
    RegFrom string `gorm:"column:reg_from;type:varchar(10);" json:"reg_from"`
    SurplusPassword string `gorm:"column:surplus_password;type:varchar(255);" json:"surplus_password"`
    Status int `gorm:"column:status;type:int(11);index:status;" json:"status"`
    AuthKey string `gorm:"column:auth_key;type:varchar(255);" json:"auth_key"`
    Type int `gorm:"column:type;type:tinyint(1);" json:"type"`
    IsReal int `gorm:"column:is_real;type:tinyint(1);" json:"is_real"`
    ShoppingStatus int `gorm:"column:shopping_status;type:tinyint(1);" json:"shopping_status"`
    CommentStatus int `gorm:"column:comment_status;type:tinyint(1);" json:"comment_status"`
    RoleID int `gorm:"column:role_id;type:int(11);" json:"role_id"`
    AuthCodes string `gorm:"column:auth_codes;type:text;" json:"auth_codes"`
    CompanyName string `gorm:"column:company_name;type:varchar(60);" json:"company_name"`
    CompanyRegionCode string `gorm:"column:company_region_code;type:varchar(60);" json:"company_region_code"`
    CompanyAddress string `gorm:"column:company_address;type:varchar(255);" json:"company_address"`
    PurposeType string `gorm:"column:purpose_type;type:varchar(255);" json:"purpose_type"`
    ReferralMobile string `gorm:"column:referral_mobile;type:varchar(20);" json:"referral_mobile"`
    Employees int `gorm:"column:employees;type:tinyint(1);" json:"employees"`
    Industry int `gorm:"column:industry;type:tinyint(2);" json:"industry"`
    Nature int `gorm:"column:nature;type:tinyint(1);" json:"nature"`
    ContactName string `gorm:"column:contact_name;type:varchar(60);" json:"contact_name"`
    Department int `gorm:"column:department;type:tinyint(1);" json:"department"`
    CompanyTel string `gorm:"column:company_tel;type:varchar(20);" json:"company_tel"`
    QqKey string `gorm:"column:qq_key;type:varchar(255);" json:"qq_key"`
    WeiboKey string `gorm:"column:weibo_key;type:varchar(255);" json:"weibo_key"`
    WeixinKey string `gorm:"column:weixin_key;type:varchar(255);index:weixin_key;" json:"weixin_key"`
    UserRemark string `gorm:"column:user_remark;type:text;" json:"user_remark"`
    InviteCode string `gorm:"column:invite_code;type:varchar(10);index:invite_code;" json:"invite_code"`
    ParentID int `gorm:"column:parent_id;type:int(11);index:parent_id;" json:"parent_id"`
    IsRecommend int `gorm:"column:is_recommend;type:tinyint(1);" json:"is_recommend"`
    CustomsMoney string `gorm:"column:customs_money;type:text;" json:"customs_money"`
    AuditStatus int `gorm:"column:audit_status;type:int(11);" json:"audit_status"`
    EnterpriseTypeID int `gorm:"column:enterprise_type_id;type:int(11);" json:"enterprise_type_id"`
    Contacts string `gorm:"column:contacts;type:varchar(60);" json:"contacts"`
    CorporationName string `gorm:"column:corporation_name;type:varchar(255);" json:"corporation_name"`
    BusinessAddress string `gorm:"column:business_address;type:varchar(60);" json:"business_address"`
    OfflineID string `gorm:"column:offline_id;type:varchar(60);" json:"offline_id"`
    SubmitTime int `gorm:"column:submit_time;type:int(11);" json:"submit_time"`
    ErrorReason string `gorm:"column:error_reason;type:varchar(255);" json:"error_reason"`
    IsFirstMarket int `gorm:"column:is_first_market;type:int(11);" json:"is_first_market"`
    HdRegionCodes string `gorm:"column:hd_region_codes;type:text;" json:"hd_region_codes"`
    HdGroupID string `gorm:"column:hd_group_id;type:varchar(60);" json:"hd_group_id"`
    BusinessLicenseID string `gorm:"column:business_license_id;type:varchar(60);" json:"business_license_id"`
    BusinessLicenseName string `gorm:"column:business_license_name;type:varchar(255);" json:"business_license_name"`
    BusinessLicenseValidity int `gorm:"column:business_license_validity;type:int(11);" json:"business_license_validity"`
    BusinessQualification string `gorm:"column:business_qualification;type:text;" json:"business_qualification"`
    IsSync int `gorm:"column:is_sync;type:int(11);" json:"is_sync"`
    QualificationImages string `gorm:"column:qualification_images;type:text;" json:"qualification_images"`
    
}

// TableName ...
func ( TpSzyUser) TableName() string {
	return "tp_szy_user"
}

 
