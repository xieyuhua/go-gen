

package service
import (
	"xieyuhua/dao"
	"xieyuhua/model"
)

// ZoomServerService 服务
type ZoomServerService struct {
}

// ZoomServerCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomServerCreateParam struct {
	model.ZoomServer
}

// ZoomServerCreateBack  返回参数
type ZoomServerCreateBack struct {
	model.ZoomServer
}

// Create 创建
func ( *ZoomServerService) Create(p *ZoomServerCreateParam) (*ZoomServerCreateBack, error) {
	dao := &dao.ZoomServerDao{}
	 
	data, err := dao.Create(&p.ZoomServer)

	if err != nil {
		return nil, err

	}

	var back = ZoomServerCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomServerService) Delete(id int) error {
	dao := &dao.ZoomServerDao{}
	return dao.Delete(id)
}

 

// ZoomServerSelectBack  返回参数
type ZoomServerSelectBack struct {
	model.ZoomServer
}

// All ...
func (*ZoomServerService) All() (*[]model.ZoomServer, error) {
	dao := &dao.ZoomServerDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomServerService) Select(id int) (*ZoomServerSelectBack, error) {
	dao := &dao.ZoomServerDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomServerSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomServerUpdateParam   参数 
type ZoomServerUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomServerUpdateBack  返回参数
type ZoomServerUpdateBack struct {
	model.ZoomServer
}

// Update ...
func (*ZoomServerService) Update(p *ZoomServerUpdateParam) (*ZoomServerUpdateBack, error) {
	dao := &dao.ZoomServerDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomServerUpdateBack {
		*data,
	}
	 
	return &back, nil
}

