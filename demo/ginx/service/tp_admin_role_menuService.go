

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminRoleMenuService 服务
type TpAdminRoleMenuService struct {
}

// TpAdminRoleMenuCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminRoleMenuCreateParam struct {
	model.TpAdminRoleMenu
}

// TpAdminRoleMenuCreateBack  返回参数
type TpAdminRoleMenuCreateBack struct {
	model.TpAdminRoleMenu
}

// Create 创建
func ( *TpAdminRoleMenuService) Create(p *TpAdminRoleMenuCreateParam) (*TpAdminRoleMenuCreateBack, error) {
	dao := &dao.TpAdminRoleMenuDao{}
	 
	data, err := dao.Create(&p.TpAdminRoleMenu)

	if err != nil {
		return nil, err

	}

	var back = TpAdminRoleMenuCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminRoleMenuService) Delete(id int) error {
	dao := &dao.TpAdminRoleMenuDao{}
	return dao.Delete(id)
}

 

// TpAdminRoleMenuSelectBack  返回参数
type TpAdminRoleMenuSelectBack struct {
	model.TpAdminRoleMenu
}

// Select ...
func (*TpAdminRoleMenuService) Select(id int) (*TpAdminRoleMenuSelectBack, error) {
	dao := &dao.TpAdminRoleMenuDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminRoleMenuSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminRoleMenuUpdateParam   参数 
type TpAdminRoleMenuUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminRoleMenuUpdateBack  返回参数
type TpAdminRoleMenuUpdateBack struct {
	model.TpAdminRoleMenu
}

// Update ...
func (*TpAdminRoleMenuService) Update(p *TpAdminRoleMenuUpdateParam) (*TpAdminRoleMenuUpdateBack, error) {
	dao := &dao.TpAdminRoleMenuDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminRoleMenuUpdateBack {
		*data,
	}
	 
	return &back, nil
}

