

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopOrderGoodService 服务
type TpShopOrderGoodService struct {
}

// TpShopOrderGoodCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopOrderGoodCreateParam struct {
	model.TpShopOrderGood
}

// TpShopOrderGoodCreateBack  返回参数
type TpShopOrderGoodCreateBack struct {
	model.TpShopOrderGood
}

// Create 创建
func ( *TpShopOrderGoodService) Create(p *TpShopOrderGoodCreateParam) (*TpShopOrderGoodCreateBack, error) {
	dao := &dao.TpShopOrderGoodDao{}
	 
	data, err := dao.Create(&p.TpShopOrderGood)

	if err != nil {
		return nil, err

	}

	var back = TpShopOrderGoodCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopOrderGoodService) Delete(id int) error {
	dao := &dao.TpShopOrderGoodDao{}
	return dao.Delete(id)
}

 

// TpShopOrderGoodSelectBack  返回参数
type TpShopOrderGoodSelectBack struct {
	model.TpShopOrderGood
}

// Select ...
func (*TpShopOrderGoodService) Select(id int) (*TpShopOrderGoodSelectBack, error) {
	dao := &dao.TpShopOrderGoodDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopOrderGoodSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopOrderGoodUpdateParam   参数 
type TpShopOrderGoodUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopOrderGoodUpdateBack  返回参数
type TpShopOrderGoodUpdateBack struct {
	model.TpShopOrderGood
}

// Update ...
func (*TpShopOrderGoodService) Update(p *TpShopOrderGoodUpdateParam) (*TpShopOrderGoodUpdateBack, error) {
	dao := &dao.TpShopOrderGoodDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopOrderGoodUpdateBack {
		*data,
	}
	 
	return &back, nil
}

