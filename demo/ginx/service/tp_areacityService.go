

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAreacityService 服务
type TpAreacityService struct {
}

// TpAreacityCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAreacityCreateParam struct {
	model.TpAreacity
}

// TpAreacityCreateBack  返回参数
type TpAreacityCreateBack struct {
	model.TpAreacity
}

// Create 创建
func ( *TpAreacityService) Create(p *TpAreacityCreateParam) (*TpAreacityCreateBack, error) {
	dao := &dao.TpAreacityDao{}
	 
	data, err := dao.Create(&p.TpAreacity)

	if err != nil {
		return nil, err

	}

	var back = TpAreacityCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAreacityService) Delete(id int) error {
	dao := &dao.TpAreacityDao{}
	return dao.Delete(id)
}

 

// TpAreacitySelectBack  返回参数
type TpAreacitySelectBack struct {
	model.TpAreacity
}

// Select ...
func (*TpAreacityService) Select(id int) (*TpAreacitySelectBack, error) {
	dao := &dao.TpAreacityDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAreacitySelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAreacityUpdateParam   参数 
type TpAreacityUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAreacityUpdateBack  返回参数
type TpAreacityUpdateBack struct {
	model.TpAreacity
}

// Update ...
func (*TpAreacityService) Update(p *TpAreacityUpdateParam) (*TpAreacityUpdateBack, error) {
	dao := &dao.TpAreacityDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAreacityUpdateBack {
		*data,
	}
	 
	return &back, nil
}

