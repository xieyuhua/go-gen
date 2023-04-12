

package service
import (
	"xieyuhuas/dao"
	"xieyuhuas/model"
)

// ZoomUserTokenService 服务
type ZoomUserTokenService struct {
}

// ZoomUserTokenCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomUserTokenCreateParam struct {
	model.ZoomUserToken
}

// ZoomUserTokenCreateBack  返回参数
type ZoomUserTokenCreateBack struct {
	model.ZoomUserToken
}

// Create 创建
func ( *ZoomUserTokenService) Create(p *ZoomUserTokenCreateParam) (*ZoomUserTokenCreateBack, error) {
	dao := &dao.ZoomUserTokenDao{}
	 
	data, err := dao.Create(&p.ZoomUserToken)

	if err != nil {
		return nil, err

	}

	var back = ZoomUserTokenCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomUserTokenService) Delete(id int) error {
	dao := &dao.ZoomUserTokenDao{}
	return dao.Delete(id)
}

 

// ZoomUserTokenSelectBack  返回参数
type ZoomUserTokenSelectBack struct {
	model.ZoomUserToken
}

// All ...
func (*ZoomUserTokenService) All() (*[]model.ZoomUserToken, error) {
	dao := &dao.ZoomUserTokenDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomUserTokenService) Select(id int) (*ZoomUserTokenSelectBack, error) {
	dao := &dao.ZoomUserTokenDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomUserTokenSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomUserTokenUpdateParam   参数 
type ZoomUserTokenUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomUserTokenUpdateBack  返回参数
type ZoomUserTokenUpdateBack struct {
	model.ZoomUserToken
}

// Update ...
func (*ZoomUserTokenService) Update(p *ZoomUserTokenUpdateParam) (*ZoomUserTokenUpdateBack, error) {
	dao := &dao.ZoomUserTokenDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomUserTokenUpdateBack {
		*data,
	}
	 
	return &back, nil
}

