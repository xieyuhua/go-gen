

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminPermissionService 服务
type TpAdminPermissionService struct {
}

// TpAdminPermissionCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminPermissionCreateParam struct {
	model.TpAdminPermission
}

// TpAdminPermissionCreateBack  返回参数
type TpAdminPermissionCreateBack struct {
	model.TpAdminPermission
}

// Create 创建
func ( *TpAdminPermissionService) Create(p *TpAdminPermissionCreateParam) (*TpAdminPermissionCreateBack, error) {
	dao := &dao.TpAdminPermissionDao{}
	 
	data, err := dao.Create(&p.TpAdminPermission)

	if err != nil {
		return nil, err

	}

	var back = TpAdminPermissionCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminPermissionService) Delete(id int) error {
	dao := &dao.TpAdminPermissionDao{}
	return dao.Delete(id)
}

 

// TpAdminPermissionSelectBack  返回参数
type TpAdminPermissionSelectBack struct {
	model.TpAdminPermission
}

// Select ...
func (*TpAdminPermissionService) Select(id int) (*TpAdminPermissionSelectBack, error) {
	dao := &dao.TpAdminPermissionDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminPermissionSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminPermissionUpdateParam   参数 
type TpAdminPermissionUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminPermissionUpdateBack  返回参数
type TpAdminPermissionUpdateBack struct {
	model.TpAdminPermission
}

// Update ...
func (*TpAdminPermissionService) Update(p *TpAdminPermissionUpdateParam) (*TpAdminPermissionUpdateBack, error) {
	dao := &dao.TpAdminPermissionDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminPermissionUpdateBack {
		*data,
	}
	 
	return &back, nil
}

