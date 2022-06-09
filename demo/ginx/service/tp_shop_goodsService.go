

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopGoodService 服务
type TpShopGoodService struct {
}

// TpShopGoodCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopGoodCreateParam struct {
	model.TpShopGood
}

// TpShopGoodCreateBack  返回参数
type TpShopGoodCreateBack struct {
	model.TpShopGood
}

// Create 创建
func ( *TpShopGoodService) Create(p *TpShopGoodCreateParam) (*TpShopGoodCreateBack, error) {
	dao := &dao.TpShopGoodDao{}
	 
	data, err := dao.Create(&p.TpShopGood)

	if err != nil {
		return nil, err

	}

	var back = TpShopGoodCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopGoodService) Delete(id int) error {
	dao := &dao.TpShopGoodDao{}
	return dao.Delete(id)
}

 

// TpShopGoodSelectBack  返回参数
type TpShopGoodSelectBack struct {
	model.TpShopGood
}

// Select ...
func (*TpShopGoodService) Select(id int) (*TpShopGoodSelectBack, error) {
	dao := &dao.TpShopGoodDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopGoodUpdateParam   参数 
type TpShopGoodUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopGoodUpdateBack  返回参数
type TpShopGoodUpdateBack struct {
	model.TpShopGood
}

// Update ...
func (*TpShopGoodService) Update(p *TpShopGoodUpdateParam) (*TpShopGoodUpdateBack, error) {
	dao := &dao.TpShopGoodDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodUpdateBack {
		*data,
	}
	 
	return &back, nil
}

