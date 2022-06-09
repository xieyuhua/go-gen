

package service
import (
	"ginx/dao"
	"ginx/model"
)

// SzyLoginUserService 服务
type SzyLoginUserService struct {
}

// SzyLoginUserCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type SzyLoginUserCreateParam struct {
	model.SzyLoginUser
}

// SzyLoginUserCreateBack  返回参数
type SzyLoginUserCreateBack struct {
	model.SzyLoginUser
}

// Create 创建
func ( *SzyLoginUserService) Create(p *SzyLoginUserCreateParam) (*SzyLoginUserCreateBack, error) {
	dao := &dao.SzyLoginUserDao{}
	 
	data, err := dao.Create(&p.SzyLoginUser)

	if err != nil {
		return nil, err

	}

	var back = SzyLoginUserCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *SzyLoginUserService) Delete(id int) error {
	dao := &dao.SzyLoginUserDao{}
	return dao.Delete(id)
}

 

// SzyLoginUserSelectBack  返回参数
type SzyLoginUserSelectBack struct {
	model.SzyLoginUser
}

// Select ...
func (*SzyLoginUserService) Select(id int) (*SzyLoginUserSelectBack, error) {
	dao := &dao.SzyLoginUserDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = SzyLoginUserSelectBack{
		*data,
	}
	  
	return &back, nil
}

// SzyLoginUserUpdateParam   参数 
type SzyLoginUserUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// SzyLoginUserUpdateBack  返回参数
type SzyLoginUserUpdateBack struct {
	model.SzyLoginUser
}

// Update ...
func (*SzyLoginUserService) Update(p *SzyLoginUserUpdateParam) (*SzyLoginUserUpdateBack, error) {
	dao := &dao.SzyLoginUserDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = SzyLoginUserUpdateBack {
		*data,
	}
	 
	return &back, nil
}

