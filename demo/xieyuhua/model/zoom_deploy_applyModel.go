

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomDeployApply ...
type ZoomDeployApply struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    SpaceID int `gorm:"column:space_id;type:int(10);" json:"space_id"`
    ProjectID int `gorm:"column:project_id;type:int(10);" json:"project_id"`
    Name string `gorm:"column:name;type:varchar(100);" json:"name"`
    Description string `gorm:"column:description;type:varchar(500);" json:"description"`
    BranchName string `gorm:"column:branch_name;type:varchar(100);" json:"branch_name"`
    CommitVersion string `gorm:"column:commit_version;type:varchar(100);" json:"commit_version"`
    AuditStatus int `gorm:"column:audit_status;type:int(10);" json:"audit_status"`
    AuditRefusalReason string `gorm:"column:audit_refusal_reason;type:varchar(500);" json:"audit_refusal_reason"`
    Status int `gorm:"column:status;type:int(10);" json:"status"`
    UserID int `gorm:"column:user_id;type:int(10);" json:"user_id"`
    RollbackID int `gorm:"column:rollback_id;type:int(10);" json:"rollback_id"`
    RollbackApplyID int `gorm:"column:rollback_apply_id;type:int(10);" json:"rollback_apply_id"`
    IsRollbackApply int `gorm:"column:is_rollback_apply;type:int(10);" json:"is_rollback_apply"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomDeployApply) TableName() string {
	return "zoom_deploy_apply"
}

 
