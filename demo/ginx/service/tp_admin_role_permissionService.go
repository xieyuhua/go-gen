

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminRolePermissionService 服务
type TpAdminRolePermissionService struct {
}

// TpAdminRolePermissionCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminRolePermissionCreateParam struct {
	model.TpAdminRolePermission
}

// TpAdminRolePermissionCreateBack  返回参数
type TpAdminRolePermissionCreateBack struct {
	model.TpAdminRolePermission
}

// Create 创建
func ( *TpAdminRolePermissionService) Create(p *TpAdminRolePermissionCreateParam) (*TpAdminRolePermissionCreateBack, error) {
	dao := &dao.TpAdminRolePermissionDao{}
	 
	data, err := dao.Create(&p.TpAdminRolePermission)

	if err != nil {
		return nil, err

	}

	var back = TpAdminRolePermissionCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminRolePermissionService) Delete(id int) error {
	dao := &dao.TpAdminRolePermissionDao{}
	return dao.Delete(id)
}

 

// TpAdminRolePermissionSelectBack  返回参数
type TpAdminRolePermissionSelectBack struct {
	model.TpAdminRolePermission
}

// Select ...
func (*TpAdminRolePermissionService) Select(id int) (*TpAdminRolePermissionSelectBack, error) {
	dao := &dao.TpAdminRolePermissionDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminRolePermissionSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminRolePermissionUpdateParam   参数 
type TpAdminRolePermissionUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminRolePermissionUpdateBack  返回参数
type TpAdminRolePermissionUpdateBack struct {
	model.TpAdminRolePermission
}

// Update ...
func (*TpAdminRolePermissionService) Update(p *TpAdminRolePermissionUpdateParam) (*TpAdminRolePermissionUpdateBack, error) {
	dao := &dao.TpAdminRolePermissionDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminRolePermissionUpdateBack {
		*data,
	}
	 
	return &back, nil
}

