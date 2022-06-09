

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpCmsCategoryService 服务
type TpCmsCategoryService struct {
}

// TpCmsCategoryCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpCmsCategoryCreateParam struct {
	model.TpCmsCategory
}

// TpCmsCategoryCreateBack  返回参数
type TpCmsCategoryCreateBack struct {
	model.TpCmsCategory
}

// Create 创建
func ( *TpCmsCategoryService) Create(p *TpCmsCategoryCreateParam) (*TpCmsCategoryCreateBack, error) {
	dao := &dao.TpCmsCategoryDao{}
	 
	data, err := dao.Create(&p.TpCmsCategory)

	if err != nil {
		return nil, err

	}

	var back = TpCmsCategoryCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpCmsCategoryService) Delete(id int) error {
	dao := &dao.TpCmsCategoryDao{}
	return dao.Delete(id)
}

 

// TpCmsCategorySelectBack  返回参数
type TpCmsCategorySelectBack struct {
	model.TpCmsCategory
}

// Select ...
func (*TpCmsCategoryService) Select(id int) (*TpCmsCategorySelectBack, error) {
	dao := &dao.TpCmsCategoryDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpCmsCategorySelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpCmsCategoryUpdateParam   参数 
type TpCmsCategoryUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpCmsCategoryUpdateBack  返回参数
type TpCmsCategoryUpdateBack struct {
	model.TpCmsCategory
}

// Update ...
func (*TpCmsCategoryService) Update(p *TpCmsCategoryUpdateParam) (*TpCmsCategoryUpdateBack, error) {
	dao := &dao.TpCmsCategoryDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpCmsCategoryUpdateBack {
		*data,
	}
	 
	return &back, nil
}

