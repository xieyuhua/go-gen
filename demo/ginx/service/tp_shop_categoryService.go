

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopCategoryService 服务
type TpShopCategoryService struct {
}

// TpShopCategoryCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopCategoryCreateParam struct {
	model.TpShopCategory
}

// TpShopCategoryCreateBack  返回参数
type TpShopCategoryCreateBack struct {
	model.TpShopCategory
}

// Create 创建
func ( *TpShopCategoryService) Create(p *TpShopCategoryCreateParam) (*TpShopCategoryCreateBack, error) {
	dao := &dao.TpShopCategoryDao{}
	 
	data, err := dao.Create(&p.TpShopCategory)

	if err != nil {
		return nil, err

	}

	var back = TpShopCategoryCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopCategoryService) Delete(id int) error {
	dao := &dao.TpShopCategoryDao{}
	return dao.Delete(id)
}

 

// TpShopCategorySelectBack  返回参数
type TpShopCategorySelectBack struct {
	model.TpShopCategory
}

// Select ...
func (*TpShopCategoryService) Select(id int) (*TpShopCategorySelectBack, error) {
	dao := &dao.TpShopCategoryDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopCategorySelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopCategoryUpdateParam   参数 
type TpShopCategoryUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopCategoryUpdateBack  返回参数
type TpShopCategoryUpdateBack struct {
	model.TpShopCategory
}

// Update ...
func (*TpShopCategoryService) Update(p *TpShopCategoryUpdateParam) (*TpShopCategoryUpdateBack, error) {
	dao := &dao.TpShopCategoryDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopCategoryUpdateBack {
		*data,
	}
	 
	return &back, nil
}

