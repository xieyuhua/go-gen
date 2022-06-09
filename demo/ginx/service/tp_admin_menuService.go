

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminMenuService 服务
type TpAdminMenuService struct {
}

// TpAdminMenuCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminMenuCreateParam struct {
	model.TpAdminMenu
}

// TpAdminMenuCreateBack  返回参数
type TpAdminMenuCreateBack struct {
	model.TpAdminMenu
}

// Create 创建
func ( *TpAdminMenuService) Create(p *TpAdminMenuCreateParam) (*TpAdminMenuCreateBack, error) {
	dao := &dao.TpAdminMenuDao{}
	 
	data, err := dao.Create(&p.TpAdminMenu)

	if err != nil {
		return nil, err

	}

	var back = TpAdminMenuCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminMenuService) Delete(id int) error {
	dao := &dao.TpAdminMenuDao{}
	return dao.Delete(id)
}

 

// TpAdminMenuSelectBack  返回参数
type TpAdminMenuSelectBack struct {
	model.TpAdminMenu
}

// Select ...
func (*TpAdminMenuService) Select(id int) (*TpAdminMenuSelectBack, error) {
	dao := &dao.TpAdminMenuDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminMenuSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminMenuUpdateParam   参数 
type TpAdminMenuUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminMenuUpdateBack  返回参数
type TpAdminMenuUpdateBack struct {
	model.TpAdminMenu
}

// Update ...
func (*TpAdminMenuService) Update(p *TpAdminMenuUpdateParam) (*TpAdminMenuUpdateBack, error) {
	dao := &dao.TpAdminMenuDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminMenuUpdateBack {
		*data,
	}
	 
	return &back, nil
}

