

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpSzyUserService 服务
type TpSzyUserService struct {
}

// TpSzyUserCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpSzyUserCreateParam struct {
	model.TpSzyUser
}

// TpSzyUserCreateBack  返回参数
type TpSzyUserCreateBack struct {
	model.TpSzyUser
}

// Create 创建
func ( *TpSzyUserService) Create(p *TpSzyUserCreateParam) (*TpSzyUserCreateBack, error) {
	dao := &dao.TpSzyUserDao{}
	 
	data, err := dao.Create(&p.TpSzyUser)

	if err != nil {
		return nil, err

	}

	var back = TpSzyUserCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpSzyUserService) Delete(id int) error {
	dao := &dao.TpSzyUserDao{}
	return dao.Delete(id)
}

 

// TpSzyUserSelectBack  返回参数
type TpSzyUserSelectBack struct {
	model.TpSzyUser
}

// Select ...
func (*TpSzyUserService) Select(id int) (*TpSzyUserSelectBack, error) {
	dao := &dao.TpSzyUserDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpSzyUserSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpSzyUserUpdateParam   参数 
type TpSzyUserUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpSzyUserUpdateBack  返回参数
type TpSzyUserUpdateBack struct {
	model.TpSzyUser
}

// Update ...
func (*TpSzyUserService) Update(p *TpSzyUserUpdateParam) (*TpSzyUserUpdateBack, error) {
	dao := &dao.TpSzyUserDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpSzyUserUpdateBack {
		*data,
	}
	 
	return &back, nil
}

