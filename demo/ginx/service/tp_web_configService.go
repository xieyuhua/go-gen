

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpWebConfigService 服务
type TpWebConfigService struct {
}

// TpWebConfigCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpWebConfigCreateParam struct {
	model.TpWebConfig
}

// TpWebConfigCreateBack  返回参数
type TpWebConfigCreateBack struct {
	model.TpWebConfig
}

// Create 创建
func ( *TpWebConfigService) Create(p *TpWebConfigCreateParam) (*TpWebConfigCreateBack, error) {
	dao := &dao.TpWebConfigDao{}
	 
	data, err := dao.Create(&p.TpWebConfig)

	if err != nil {
		return nil, err

	}

	var back = TpWebConfigCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpWebConfigService) Delete(id int) error {
	dao := &dao.TpWebConfigDao{}
	return dao.Delete(id)
}

 

// TpWebConfigSelectBack  返回参数
type TpWebConfigSelectBack struct {
	model.TpWebConfig
}

// Select ...
func (*TpWebConfigService) Select(id int) (*TpWebConfigSelectBack, error) {
	dao := &dao.TpWebConfigDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpWebConfigSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpWebConfigUpdateParam   参数 
type TpWebConfigUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpWebConfigUpdateBack  返回参数
type TpWebConfigUpdateBack struct {
	model.TpWebConfig
}

// Update ...
func (*TpWebConfigService) Update(p *TpWebConfigUpdateParam) (*TpWebConfigUpdateBack, error) {
	dao := &dao.TpWebConfigDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpWebConfigUpdateBack {
		*data,
	}
	 
	return &back, nil
}

