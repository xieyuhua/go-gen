

package service
import (
	"ddada/dao"
	"ddada/model"
)

// ZoomProjectService 服务
type ZoomProjectService struct {
}

// ZoomProjectCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomProjectCreateParam struct {
	model.ZoomProject
}

// ZoomProjectCreateBack  返回参数
type ZoomProjectCreateBack struct {
	model.ZoomProject
}

// Create 创建
func ( *ZoomProjectService) Create(p *ZoomProjectCreateParam) (*ZoomProjectCreateBack, error) {
	dao := &dao.ZoomProjectDao{}
	 
	data, err := dao.Create(&p.ZoomProject)

	if err != nil {
		return nil, err

	}

	var back = ZoomProjectCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomProjectService) Delete(id int) error {
	dao := &dao.ZoomProjectDao{}
	return dao.Delete(id)
}

 

// ZoomProjectSelectBack  返回参数
type ZoomProjectSelectBack struct {
	model.ZoomProject
}

// All ...
func (*ZoomProjectService) All() (*[]model.ZoomProject, error) {
	dao := &dao.ZoomProjectDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return &data, nil
}


// Select ...
func (*ZoomProjectService) Select(id int) (*ZoomProjectSelectBack, error) {
	dao := &dao.ZoomProjectDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomProjectSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomProjectUpdateParam   参数 
type ZoomProjectUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomProjectUpdateBack  返回参数
type ZoomProjectUpdateBack struct {
	model.ZoomProject
}

// Update ...
func (*ZoomProjectService) Update(p *ZoomProjectUpdateParam) (*ZoomProjectUpdateBack, error) {
	dao := &dao.ZoomProjectDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomProjectUpdateBack {
		*data,
	}
	 
	return &back, nil
}

