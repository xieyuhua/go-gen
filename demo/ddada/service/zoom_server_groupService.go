

package service
import (
	"ddada/dao"
	"ddada/model"
)

// ZoomServerGroupService 服务
type ZoomServerGroupService struct {
}

// ZoomServerGroupCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomServerGroupCreateParam struct {
	model.ZoomServerGroup
}

// ZoomServerGroupCreateBack  返回参数
type ZoomServerGroupCreateBack struct {
	model.ZoomServerGroup
}

// Create 创建
func ( *ZoomServerGroupService) Create(p *ZoomServerGroupCreateParam) (*ZoomServerGroupCreateBack, error) {
	dao := &dao.ZoomServerGroupDao{}
	 
	data, err := dao.Create(&p.ZoomServerGroup)

	if err != nil {
		return nil, err

	}

	var back = ZoomServerGroupCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomServerGroupService) Delete(id int) error {
	dao := &dao.ZoomServerGroupDao{}
	return dao.Delete(id)
}

 

// ZoomServerGroupSelectBack  返回参数
type ZoomServerGroupSelectBack struct {
	model.ZoomServerGroup
}

// All ...
func (*ZoomServerGroupService) All() (*[]model.ZoomServerGroup, error) {
	dao := &dao.ZoomServerGroupDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return &data, nil
}


// Select ...
func (*ZoomServerGroupService) Select(id int) (*ZoomServerGroupSelectBack, error) {
	dao := &dao.ZoomServerGroupDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomServerGroupSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomServerGroupUpdateParam   参数 
type ZoomServerGroupUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomServerGroupUpdateBack  返回参数
type ZoomServerGroupUpdateBack struct {
	model.ZoomServerGroup
}

// Update ...
func (*ZoomServerGroupService) Update(p *ZoomServerGroupUpdateParam) (*ZoomServerGroupUpdateBack, error) {
	dao := &dao.ZoomServerGroupDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomServerGroupUpdateBack {
		*data,
	}
	 
	return &back, nil
}

