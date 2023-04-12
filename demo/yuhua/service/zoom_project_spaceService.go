

package service
import (
	"yuhua/dao"
	"yuhua/model"
)

// ZoomProjectSpaceService 服务
type ZoomProjectSpaceService struct {
}

// ZoomProjectSpaceCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomProjectSpaceCreateParam struct {
	model.ZoomProjectSpace
}

// ZoomProjectSpaceCreateBack  返回参数
type ZoomProjectSpaceCreateBack struct {
	model.ZoomProjectSpace
}

// Create 创建
func ( *ZoomProjectSpaceService) Create(p *ZoomProjectSpaceCreateParam) (*ZoomProjectSpaceCreateBack, error) {
	dao := &dao.ZoomProjectSpaceDao{}
	 
	data, err := dao.Create(&p.ZoomProjectSpace)

	if err != nil {
		return nil, err

	}

	var back = ZoomProjectSpaceCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomProjectSpaceService) Delete(id int) error {
	dao := &dao.ZoomProjectSpaceDao{}
	return dao.Delete(id)
}

 

// ZoomProjectSpaceSelectBack  返回参数
type ZoomProjectSpaceSelectBack struct {
	model.ZoomProjectSpace
}

// Select ...
func (*ZoomProjectSpaceService) Select(id int) (*ZoomProjectSpaceSelectBack, error) {
	dao := &dao.ZoomProjectSpaceDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomProjectSpaceSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomProjectSpaceUpdateParam   参数 
type ZoomProjectSpaceUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomProjectSpaceUpdateBack  返回参数
type ZoomProjectSpaceUpdateBack struct {
	model.ZoomProjectSpace
}

// Update ...
func (*ZoomProjectSpaceService) Update(p *ZoomProjectSpaceUpdateParam) (*ZoomProjectSpaceUpdateBack, error) {
	dao := &dao.ZoomProjectSpaceDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomProjectSpaceUpdateBack {
		*data,
	}
	 
	return &back, nil
}

