

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminUserService 服务
type TpAdminUserService struct {
}

// TpAdminUserCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminUserCreateParam struct {
	model.TpAdminUser
}

// TpAdminUserCreateBack  返回参数
type TpAdminUserCreateBack struct {
	model.TpAdminUser
}

// Create 创建
func ( *TpAdminUserService) Create(p *TpAdminUserCreateParam) (*TpAdminUserCreateBack, error) {
	dao := &dao.TpAdminUserDao{}
	 
	data, err := dao.Create(&p.TpAdminUser)

	if err != nil {
		return nil, err

	}

	var back = TpAdminUserCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminUserService) Delete(id int) error {
	dao := &dao.TpAdminUserDao{}
	return dao.Delete(id)
}

 

// TpAdminUserSelectBack  返回参数
type TpAdminUserSelectBack struct {
	model.TpAdminUser
}

// Select ...
func (*TpAdminUserService) Select(id int) (*TpAdminUserSelectBack, error) {
	dao := &dao.TpAdminUserDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminUserSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminUserUpdateParam   参数 
type TpAdminUserUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminUserUpdateBack  返回参数
type TpAdminUserUpdateBack struct {
	model.TpAdminUser
}

// Update ...
func (*TpAdminUserService) Update(p *TpAdminUserUpdateParam) (*TpAdminUserUpdateBack, error) {
	dao := &dao.TpAdminUserDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminUserUpdateBack {
		*data,
	}
	 
	return &back, nil
}

