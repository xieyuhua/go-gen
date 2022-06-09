

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpCmsBannerService 服务
type TpCmsBannerService struct {
}

// TpCmsBannerCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpCmsBannerCreateParam struct {
	model.TpCmsBanner
}

// TpCmsBannerCreateBack  返回参数
type TpCmsBannerCreateBack struct {
	model.TpCmsBanner
}

// Create 创建
func ( *TpCmsBannerService) Create(p *TpCmsBannerCreateParam) (*TpCmsBannerCreateBack, error) {
	dao := &dao.TpCmsBannerDao{}
	 
	data, err := dao.Create(&p.TpCmsBanner)

	if err != nil {
		return nil, err

	}

	var back = TpCmsBannerCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpCmsBannerService) Delete(id int) error {
	dao := &dao.TpCmsBannerDao{}
	return dao.Delete(id)
}

 

// TpCmsBannerSelectBack  返回参数
type TpCmsBannerSelectBack struct {
	model.TpCmsBanner
}

// Select ...
func (*TpCmsBannerService) Select(id int) (*TpCmsBannerSelectBack, error) {
	dao := &dao.TpCmsBannerDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpCmsBannerSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpCmsBannerUpdateParam   参数 
type TpCmsBannerUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpCmsBannerUpdateBack  返回参数
type TpCmsBannerUpdateBack struct {
	model.TpCmsBanner
}

// Update ...
func (*TpCmsBannerService) Update(p *TpCmsBannerUpdateParam) (*TpCmsBannerUpdateBack, error) {
	dao := &dao.TpCmsBannerDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpCmsBannerUpdateBack {
		*data,
	}
	 
	return &back, nil
}

