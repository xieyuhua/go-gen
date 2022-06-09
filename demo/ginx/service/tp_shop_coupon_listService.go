

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopCouponListService 服务
type TpShopCouponListService struct {
}

// TpShopCouponListCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopCouponListCreateParam struct {
	model.TpShopCouponList
}

// TpShopCouponListCreateBack  返回参数
type TpShopCouponListCreateBack struct {
	model.TpShopCouponList
}

// Create 创建
func ( *TpShopCouponListService) Create(p *TpShopCouponListCreateParam) (*TpShopCouponListCreateBack, error) {
	dao := &dao.TpShopCouponListDao{}
	 
	data, err := dao.Create(&p.TpShopCouponList)

	if err != nil {
		return nil, err

	}

	var back = TpShopCouponListCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopCouponListService) Delete(id int) error {
	dao := &dao.TpShopCouponListDao{}
	return dao.Delete(id)
}

 

// TpShopCouponListSelectBack  返回参数
type TpShopCouponListSelectBack struct {
	model.TpShopCouponList
}

// Select ...
func (*TpShopCouponListService) Select(id int) (*TpShopCouponListSelectBack, error) {
	dao := &dao.TpShopCouponListDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopCouponListSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopCouponListUpdateParam   参数 
type TpShopCouponListUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopCouponListUpdateBack  返回参数
type TpShopCouponListUpdateBack struct {
	model.TpShopCouponList
}

// Update ...
func (*TpShopCouponListService) Update(p *TpShopCouponListUpdateParam) (*TpShopCouponListUpdateBack, error) {
	dao := &dao.TpShopCouponListDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopCouponListUpdateBack {
		*data,
	}
	 
	return &back, nil
}

