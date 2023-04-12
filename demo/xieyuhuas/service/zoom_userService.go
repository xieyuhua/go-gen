

package service
import (
	"xieyuhuas/dao"
	"xieyuhuas/model"
)

// ZoomUserService 服务
type ZoomUserService struct {
}

// ZoomUserCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomUserCreateParam struct {
	model.ZoomUser
}

// ZoomUserCreateBack  返回参数
type ZoomUserCreateBack struct {
	model.ZoomUser
}

// Create 创建
func ( *ZoomUserService) Create(p *ZoomUserCreateParam) (*ZoomUserCreateBack, error) {
	dao := &dao.ZoomUserDao{}
	 
	data, err := dao.Create(&p.ZoomUser)

	if err != nil {
		return nil, err

	}

	var back = ZoomUserCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomUserService) Delete(id int) error {
	dao := &dao.ZoomUserDao{}
	return dao.Delete(id)
}

 

// ZoomUserSelectBack  返回参数
type ZoomUserSelectBack struct {
	model.ZoomUser
}

// All ...
func (*ZoomUserService) All() (*[]model.ZoomUser, error) {
	dao := &dao.ZoomUserDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomUserService) Select(id int) (*ZoomUserSelectBack, error) {
	dao := &dao.ZoomUserDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomUserSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomUserUpdateParam   参数 
type ZoomUserUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomUserUpdateBack  返回参数
type ZoomUserUpdateBack struct {
	model.ZoomUser
}

// Update ...
func (*ZoomUserService) Update(p *ZoomUserUpdateParam) (*ZoomUserUpdateBack, error) {
	dao := &dao.ZoomUserDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomUserUpdateBack {
		*data,
	}
	 
	return &back, nil
}

