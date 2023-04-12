

package service
import (
	"xieyuhuas/dao"
	"xieyuhuas/model"
)

// ZoomDeployTaskService 服务
type ZoomDeployTaskService struct {
}

// ZoomDeployTaskCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomDeployTaskCreateParam struct {
	model.ZoomDeployTask
}

// ZoomDeployTaskCreateBack  返回参数
type ZoomDeployTaskCreateBack struct {
	model.ZoomDeployTask
}

// Create 创建
func ( *ZoomDeployTaskService) Create(p *ZoomDeployTaskCreateParam) (*ZoomDeployTaskCreateBack, error) {
	dao := &dao.ZoomDeployTaskDao{}
	 
	data, err := dao.Create(&p.ZoomDeployTask)

	if err != nil {
		return nil, err

	}

	var back = ZoomDeployTaskCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomDeployTaskService) Delete(id int) error {
	dao := &dao.ZoomDeployTaskDao{}
	return dao.Delete(id)
}

 

// ZoomDeployTaskSelectBack  返回参数
type ZoomDeployTaskSelectBack struct {
	model.ZoomDeployTask
}

// All ...
func (*ZoomDeployTaskService) All() (*[]model.ZoomDeployTask, error) {
	dao := &dao.ZoomDeployTaskDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomDeployTaskService) Select(id int) (*ZoomDeployTaskSelectBack, error) {
	dao := &dao.ZoomDeployTaskDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomDeployTaskSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomDeployTaskUpdateParam   参数 
type ZoomDeployTaskUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomDeployTaskUpdateBack  返回参数
type ZoomDeployTaskUpdateBack struct {
	model.ZoomDeployTask
}

// Update ...
func (*ZoomDeployTaskService) Update(p *ZoomDeployTaskUpdateParam) (*ZoomDeployTaskUpdateBack, error) {
	dao := &dao.ZoomDeployTaskDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomDeployTaskUpdateBack {
		*data,
	}
	 
	return &back, nil
}

