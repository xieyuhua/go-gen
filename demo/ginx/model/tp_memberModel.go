

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// TpMember ...
type TpMember struct {
    ID int `gorm:"column:id;type:int(10);primary_key;" json:"id"`
    Username string `gorm:"column:username;type:varchar(20);unique_index:unq_username;" json:"username"`
    Nickname string `gorm:"column:nickname;type:varchar(20);" json:"nickname"`
    Mobile string `gorm:"column:mobile;type:varchar(20);unique_index:unq_mobile;" json:"mobile"`
    Email string `gorm:"column:email;type:varchar(60);" json:"email"`
    Password string `gorm:"column:password;type:varchar(55);" json:"password"`
    Avatar string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
    Age int `gorm:"column:age;type:tinyint(3);" json:"age"`
    Gender string `gorm:"column:gender;type:varchar(5);" json:"gender"`
    Province int64 `gorm:"column:province;type:bigint(20);" json:"province"`
    City int64 `gorm:"column:city;type:bigint(20);" json:"city"`
    Area int64 `gorm:"column:area;type:bigint(20);" json:"area"`
    Status int `gorm:"column:status;type:tinyint(3);" json:"status"`
    Token string `gorm:"column:token;type:varchar(55);" json:"token"`
    LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime;" json:"last_login_time"`
    LastLoginIP string `gorm:"column:last_login_ip;type:varchar(55);" json:"last_login_ip"`
    Level int `gorm:"column:level;type:tinyint(3);index:idx_level;" json:"level"`
    AgentLevel int `gorm:"column:agent_level;type:tinyint(3);index:idx_agent_level;" json:"agent_level"`
    Points float64 `gorm:"column:points;type:decimal(10,2);" json:"points"`
    Money float64 `gorm:"column:money;type:decimal(10,2);" json:"money"`
    Commission float64 `gorm:"column:commission;type:decimal(10,2);" json:"commission"`
    SpendMoney float64 `gorm:"column:spend_money;type:decimal(10,2);" json:"spend_money"`
    Openid string `gorm:"column:openid;type:varchar(75);" json:"openid"`
    Remark string `gorm:"column:remark;type:varchar(255);" json:"remark"`
    DeleteTime time.Time `gorm:"column:delete_time;type:datetime;" json:"delete_time"`
    FirstLeader int `gorm:"column:first_leader;type:int(10);index:idx_first_leader;" json:"first_leader"`
    SecondLeader int `gorm:"column:second_leader;type:int(10);index:idx_second_leader;" json:"second_leader"`
    ThirdLeader int `gorm:"column:third_leader;type:int(10);index:idx_third_leader;" json:"third_leader"`
    Relation string `gorm:"column:relation;type:varchar(1000);" json:"relation"`
    CreateTime time.Time `gorm:"column:create_time;type:datetime;" json:"create_time"`
    UpdateTime string `gorm:"column:update_time;type:varchar(55);" json:"update_time"`
    AgentImg string `gorm:"column:agent_img;type:varchar(255);" json:"agent_img"`
    Salt string `gorm:"column:salt;type:varchar(10);" json:"salt"`
    Errors int `gorm:"column:errors;type:tinyint(3);" json:"errors"`
    PayPassword string `gorm:"column:pay_password;type:varchar(60);" json:"pay_password"`
    
}

// TableName ...
func ( TpMember) TableName() string {
	return "tp_member"
}

 
