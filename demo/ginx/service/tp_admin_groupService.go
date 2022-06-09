

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminGroupService 服务
type TpAdminGroupService struct {
}

// TpAdminGroupCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminGroupCreateParam struct {
	model.TpAdminGroup
}

// TpAdminGroupCreateBack  返回参数
type TpAdminGroupCreateBack struct {
	model.TpAdminGroup
}

// Create 创建
func ( *TpAdminGroupService) Create(p *TpAdminGroupCreateParam) (*TpAdminGroupCreateBack, error) {
	dao := &dao.TpAdminGroupDao{}
	 
	data, err := dao.Create(&p.TpAdminGroup)

	if err != nil {
		return nil, err

	}

	var back = TpAdminGroupCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminGroupService) Delete(id int) error {
	dao := &dao.TpAdminGroupDao{}
	return dao.Delete(id)
}

 

// TpAdminGroupSelectBack  返回参数
type TpAdminGroupSelectBack struct {
	model.TpAdminGroup
}

// Select ...
func (*TpAdminGroupService) Select(id int) (*TpAdminGroupSelectBack, error) {
	dao := &dao.TpAdminGroupDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminGroupSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminGroupUpdateParam   参数 
type TpAdminGroupUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminGroupUpdateBack  返回参数
type TpAdminGroupUpdateBack struct {
	model.TpAdminGroup
}

// Update ...
func (*TpAdminGroupService) Update(p *TpAdminGroupUpdateParam) (*TpAdminGroupUpdateBack, error) {
	dao := &dao.TpAdminGroupDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminGroupUpdateBack {
		*data,
	}
	 
	return &back, nil
}

