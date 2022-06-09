

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminRoleService 服务
type TpAdminRoleService struct {
}

// TpAdminRoleCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminRoleCreateParam struct {
	model.TpAdminRole
}

// TpAdminRoleCreateBack  返回参数
type TpAdminRoleCreateBack struct {
	model.TpAdminRole
}

// Create 创建
func ( *TpAdminRoleService) Create(p *TpAdminRoleCreateParam) (*TpAdminRoleCreateBack, error) {
	dao := &dao.TpAdminRoleDao{}
	 
	data, err := dao.Create(&p.TpAdminRole)

	if err != nil {
		return nil, err

	}

	var back = TpAdminRoleCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminRoleService) Delete(id int) error {
	dao := &dao.TpAdminRoleDao{}
	return dao.Delete(id)
}

 

// TpAdminRoleSelectBack  返回参数
type TpAdminRoleSelectBack struct {
	model.TpAdminRole
}

// Select ...
func (*TpAdminRoleService) Select(id int) (*TpAdminRoleSelectBack, error) {
	dao := &dao.TpAdminRoleDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminRoleSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminRoleUpdateParam   参数 
type TpAdminRoleUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminRoleUpdateBack  返回参数
type TpAdminRoleUpdateBack struct {
	model.TpAdminRole
}

// Update ...
func (*TpAdminRoleService) Update(p *TpAdminRoleUpdateParam) (*TpAdminRoleUpdateBack, error) {
	dao := &dao.TpAdminRoleDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminRoleUpdateBack {
		*data,
	}
	 
	return &back, nil
}

