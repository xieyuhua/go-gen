

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopCouponTypeService 服务
type TpShopCouponTypeService struct {
}

// TpShopCouponTypeCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopCouponTypeCreateParam struct {
	model.TpShopCouponType
}

// TpShopCouponTypeCreateBack  返回参数
type TpShopCouponTypeCreateBack struct {
	model.TpShopCouponType
}

// Create 创建
func ( *TpShopCouponTypeService) Create(p *TpShopCouponTypeCreateParam) (*TpShopCouponTypeCreateBack, error) {
	dao := &dao.TpShopCouponTypeDao{}
	 
	data, err := dao.Create(&p.TpShopCouponType)

	if err != nil {
		return nil, err

	}

	var back = TpShopCouponTypeCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopCouponTypeService) Delete(id int) error {
	dao := &dao.TpShopCouponTypeDao{}
	return dao.Delete(id)
}

 

// TpShopCouponTypeSelectBack  返回参数
type TpShopCouponTypeSelectBack struct {
	model.TpShopCouponType
}

// Select ...
func (*TpShopCouponTypeService) Select(id int) (*TpShopCouponTypeSelectBack, error) {
	dao := &dao.TpShopCouponTypeDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopCouponTypeSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopCouponTypeUpdateParam   参数 
type TpShopCouponTypeUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopCouponTypeUpdateBack  返回参数
type TpShopCouponTypeUpdateBack struct {
	model.TpShopCouponType
}

// Update ...
func (*TpShopCouponTypeService) Update(p *TpShopCouponTypeUpdateParam) (*TpShopCouponTypeUpdateBack, error) {
	dao := &dao.TpShopCouponTypeDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopCouponTypeUpdateBack {
		*data,
	}
	 
	return &back, nil
}

