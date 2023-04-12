

package service
import (
	"ddada/dao"
	"ddada/model"
)

// ZoomDeployBuildService 服务
type ZoomDeployBuildService struct {
}

// ZoomDeployBuildCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomDeployBuildCreateParam struct {
	model.ZoomDeployBuild
}

// ZoomDeployBuildCreateBack  返回参数
type ZoomDeployBuildCreateBack struct {
	model.ZoomDeployBuild
}

// Create 创建
func ( *ZoomDeployBuildService) Create(p *ZoomDeployBuildCreateParam) (*ZoomDeployBuildCreateBack, error) {
	dao := &dao.ZoomDeployBuildDao{}
	 
	data, err := dao.Create(&p.ZoomDeployBuild)

	if err != nil {
		return nil, err

	}

	var back = ZoomDeployBuildCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomDeployBuildService) Delete(id int) error {
	dao := &dao.ZoomDeployBuildDao{}
	return dao.Delete(id)
}

 

// ZoomDeployBuildSelectBack  返回参数
type ZoomDeployBuildSelectBack struct {
	model.ZoomDeployBuild
}

// All ...
func (*ZoomDeployBuildService) All() (*[]model.ZoomDeployBuild, error) {
	dao := &dao.ZoomDeployBuildDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return &data, nil
}


// Select ...
func (*ZoomDeployBuildService) Select(id int) (*ZoomDeployBuildSelectBack, error) {
	dao := &dao.ZoomDeployBuildDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomDeployBuildSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomDeployBuildUpdateParam   参数 
type ZoomDeployBuildUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomDeployBuildUpdateBack  返回参数
type ZoomDeployBuildUpdateBack struct {
	model.ZoomDeployBuild
}

// Update ...
func (*ZoomDeployBuildService) Update(p *ZoomDeployBuildUpdateParam) (*ZoomDeployBuildUpdateBack, error) {
	dao := &dao.ZoomDeployBuildDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomDeployBuildUpdateBack {
		*data,
	}
	 
	return &back, nil
}

