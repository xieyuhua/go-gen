

package model
import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// ZoomProject ...
type ZoomProject struct {
    ID int `gorm:"column:id;type:int(10) AUTO_INCREMENT;primary_key;" json:"id"`
    SpaceID int `gorm:"column:space_id;type:int(10);" json:"space_id"`
    Name string `gorm:"column:name;type:varchar(100);" json:"name"`
    Description string `gorm:"column:description;type:varchar(500);" json:"description"`
    ProjectType int `gorm:"column:project_type;type:tinyint(4);" json:"project_type"`
    NeedAudit int `gorm:"column:need_audit;type:int(10);" json:"need_audit"`
    Status int `gorm:"column:status;type:int(10);" json:"status"`
    RepoURL string `gorm:"column:repo_url;type:varchar(500);" json:"repo_url"`
    DeployMode int `gorm:"column:deploy_mode;type:int(10);" json:"deploy_mode"`
    RepoBranch string `gorm:"column:repo_branch;type:varchar(100);" json:"repo_branch"`
    PreReleaseCluster int `gorm:"column:pre_release_cluster;type:int(10);" json:"pre_release_cluster"`
    OnlineCluster string `gorm:"column:online_cluster;type:varchar(2000);" json:"online_cluster"`
    DeployUser string `gorm:"column:deploy_user;type:varchar(50);" json:"deploy_user"`
    DeployPath string `gorm:"column:deploy_path;type:varchar(500);" json:"deploy_path"`
    BuildScript string `gorm:"column:build_script;type:text;" json:"build_script"`
    BuildHookScript string `gorm:"column:build_hook_script;type:text;" json:"build_hook_script"`
    DeployHookScript string `gorm:"column:deploy_hook_script;type:text;" json:"deploy_hook_script"`
    PreDeployCmd string `gorm:"column:pre_deploy_cmd;type:text;" json:"pre_deploy_cmd"`
    AfterDeployCmd string `gorm:"column:after_deploy_cmd;type:text;" json:"after_deploy_cmd"`
    AuditNotice string `gorm:"column:audit_notice;type:varchar(2000);" json:"audit_notice"`
    DeployNotice string `gorm:"column:deploy_notice;type:varchar(2000);" json:"deploy_notice"`
    Ctime int `gorm:"column:ctime;type:int(10);" json:"ctime"`
    
}

// TableName ...
func ( ZoomProject) TableName() string {
	return "zoom_project"
}

 
