

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopBrandService 服务
type TpShopBrandService struct {
}

// TpShopBrandCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopBrandCreateParam struct {
	model.TpShopBrand
}

// TpShopBrandCreateBack  返回参数
type TpShopBrandCreateBack struct {
	model.TpShopBrand
}

// Create 创建
func ( *TpShopBrandService) Create(p *TpShopBrandCreateParam) (*TpShopBrandCreateBack, error) {
	dao := &dao.TpShopBrandDao{}
	 
	data, err := dao.Create(&p.TpShopBrand)

	if err != nil {
		return nil, err

	}

	var back = TpShopBrandCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopBrandService) Delete(id int) error {
	dao := &dao.TpShopBrandDao{}
	return dao.Delete(id)
}

 

// TpShopBrandSelectBack  返回参数
type TpShopBrandSelectBack struct {
	model.TpShopBrand
}

// Select ...
func (*TpShopBrandService) Select(id int) (*TpShopBrandSelectBack, error) {
	dao := &dao.TpShopBrandDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopBrandSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopBrandUpdateParam   参数 
type TpShopBrandUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopBrandUpdateBack  返回参数
type TpShopBrandUpdateBack struct {
	model.TpShopBrand
}

// Update ...
func (*TpShopBrandService) Update(p *TpShopBrandUpdateParam) (*TpShopBrandUpdateBack, error) {
	dao := &dao.TpShopBrandDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopBrandUpdateBack {
		*data,
	}
	 
	return &back, nil
}

