

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpExtensionService 服务
type TpExtensionService struct {
}

// TpExtensionCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpExtensionCreateParam struct {
	model.TpExtension
}

// TpExtensionCreateBack  返回参数
type TpExtensionCreateBack struct {
	model.TpExtension
}

// Create 创建
func ( *TpExtensionService) Create(p *TpExtensionCreateParam) (*TpExtensionCreateBack, error) {
	dao := &dao.TpExtensionDao{}
	 
	data, err := dao.Create(&p.TpExtension)

	if err != nil {
		return nil, err

	}

	var back = TpExtensionCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpExtensionService) Delete(id int) error {
	dao := &dao.TpExtensionDao{}
	return dao.Delete(id)
}

 

// TpExtensionSelectBack  返回参数
type TpExtensionSelectBack struct {
	model.TpExtension
}

// Select ...
func (*TpExtensionService) Select(id int) (*TpExtensionSelectBack, error) {
	dao := &dao.TpExtensionDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpExtensionSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpExtensionUpdateParam   参数 
type TpExtensionUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpExtensionUpdateBack  返回参数
type TpExtensionUpdateBack struct {
	model.TpExtension
}

// Update ...
func (*TpExtensionService) Update(p *TpExtensionUpdateParam) (*TpExtensionUpdateBack, error) {
	dao := &dao.TpExtensionDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpExtensionUpdateBack {
		*data,
	}
	 
	return &back, nil
}

