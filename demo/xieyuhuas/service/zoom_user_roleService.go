

package service
import (
	"xieyuhuas/dao"
	"xieyuhuas/model"
)

// ZoomUserRoleService 服务
type ZoomUserRoleService struct {
}

// ZoomUserRoleCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomUserRoleCreateParam struct {
	model.ZoomUserRole
}

// ZoomUserRoleCreateBack  返回参数
type ZoomUserRoleCreateBack struct {
	model.ZoomUserRole
}

// Create 创建
func ( *ZoomUserRoleService) Create(p *ZoomUserRoleCreateParam) (*ZoomUserRoleCreateBack, error) {
	dao := &dao.ZoomUserRoleDao{}
	 
	data, err := dao.Create(&p.ZoomUserRole)

	if err != nil {
		return nil, err

	}

	var back = ZoomUserRoleCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomUserRoleService) Delete(id int) error {
	dao := &dao.ZoomUserRoleDao{}
	return dao.Delete(id)
}

 

// ZoomUserRoleSelectBack  返回参数
type ZoomUserRoleSelectBack struct {
	model.ZoomUserRole
}

// All ...
func (*ZoomUserRoleService) All() (*[]model.ZoomUserRole, error) {
	dao := &dao.ZoomUserRoleDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomUserRoleService) Select(id int) (*ZoomUserRoleSelectBack, error) {
	dao := &dao.ZoomUserRoleDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomUserRoleSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomUserRoleUpdateParam   参数 
type ZoomUserRoleUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomUserRoleUpdateBack  返回参数
type ZoomUserRoleUpdateBack struct {
	model.ZoomUserRole
}

// Update ...
func (*ZoomUserRoleService) Update(p *ZoomUserRoleUpdateParam) (*ZoomUserRoleUpdateBack, error) {
	dao := &dao.ZoomUserRoleDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomUserRoleUpdateBack {
		*data,
	}
	 
	return &back, nil
}

