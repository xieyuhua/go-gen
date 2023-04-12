

package service
import (
	"xieyuhuas/dao"
	"xieyuhuas/model"
)

// ZoomDeployApplyService 服务
type ZoomDeployApplyService struct {
}

// ZoomDeployApplyCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomDeployApplyCreateParam struct {
	model.ZoomDeployApply
}

// ZoomDeployApplyCreateBack  返回参数
type ZoomDeployApplyCreateBack struct {
	model.ZoomDeployApply
}

// Create 创建
func ( *ZoomDeployApplyService) Create(p *ZoomDeployApplyCreateParam) (*ZoomDeployApplyCreateBack, error) {
	dao := &dao.ZoomDeployApplyDao{}
	 
	data, err := dao.Create(&p.ZoomDeployApply)

	if err != nil {
		return nil, err

	}

	var back = ZoomDeployApplyCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomDeployApplyService) Delete(id int) error {
	dao := &dao.ZoomDeployApplyDao{}
	return dao.Delete(id)
}

 

// ZoomDeployApplySelectBack  返回参数
type ZoomDeployApplySelectBack struct {
	model.ZoomDeployApply
}

// All ...
func (*ZoomDeployApplyService) All() (*[]model.ZoomDeployApply, error) {
	dao := &dao.ZoomDeployApplyDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomDeployApplyService) Select(id int) (*ZoomDeployApplySelectBack, error) {
	dao := &dao.ZoomDeployApplyDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomDeployApplySelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomDeployApplyUpdateParam   参数 
type ZoomDeployApplyUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomDeployApplyUpdateBack  返回参数
type ZoomDeployApplyUpdateBack struct {
	model.ZoomDeployApply
}

// Update ...
func (*ZoomDeployApplyService) Update(p *ZoomDeployApplyUpdateParam) (*ZoomDeployApplyUpdateBack, error) {
	dao := &dao.ZoomDeployApplyDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomDeployApplyUpdateBack {
		*data,
	}
	 
	return &back, nil
}

