

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpCmsPositionService 服务
type TpCmsPositionService struct {
}

// TpCmsPositionCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpCmsPositionCreateParam struct {
	model.TpCmsPosition
}

// TpCmsPositionCreateBack  返回参数
type TpCmsPositionCreateBack struct {
	model.TpCmsPosition
}

// Create 创建
func ( *TpCmsPositionService) Create(p *TpCmsPositionCreateParam) (*TpCmsPositionCreateBack, error) {
	dao := &dao.TpCmsPositionDao{}
	 
	data, err := dao.Create(&p.TpCmsPosition)

	if err != nil {
		return nil, err

	}

	var back = TpCmsPositionCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpCmsPositionService) Delete(id int) error {
	dao := &dao.TpCmsPositionDao{}
	return dao.Delete(id)
}

 

// TpCmsPositionSelectBack  返回参数
type TpCmsPositionSelectBack struct {
	model.TpCmsPosition
}

// Select ...
func (*TpCmsPositionService) Select(id int) (*TpCmsPositionSelectBack, error) {
	dao := &dao.TpCmsPositionDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpCmsPositionSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpCmsPositionUpdateParam   参数 
type TpCmsPositionUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpCmsPositionUpdateBack  返回参数
type TpCmsPositionUpdateBack struct {
	model.TpCmsPosition
}

// Update ...
func (*TpCmsPositionService) Update(p *TpCmsPositionUpdateParam) (*TpCmsPositionUpdateBack, error) {
	dao := &dao.TpCmsPositionDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpCmsPositionUpdateBack {
		*data,
	}
	 
	return &back, nil
}

