

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopTagService 服务
type TpShopTagService struct {
}

// TpShopTagCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopTagCreateParam struct {
	model.TpShopTag
}

// TpShopTagCreateBack  返回参数
type TpShopTagCreateBack struct {
	model.TpShopTag
}

// Create 创建
func ( *TpShopTagService) Create(p *TpShopTagCreateParam) (*TpShopTagCreateBack, error) {
	dao := &dao.TpShopTagDao{}
	 
	data, err := dao.Create(&p.TpShopTag)

	if err != nil {
		return nil, err

	}

	var back = TpShopTagCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopTagService) Delete(id int) error {
	dao := &dao.TpShopTagDao{}
	return dao.Delete(id)
}

 

// TpShopTagSelectBack  返回参数
type TpShopTagSelectBack struct {
	model.TpShopTag
}

// Select ...
func (*TpShopTagService) Select(id int) (*TpShopTagSelectBack, error) {
	dao := &dao.TpShopTagDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopTagSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopTagUpdateParam   参数 
type TpShopTagUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopTagUpdateBack  返回参数
type TpShopTagUpdateBack struct {
	model.TpShopTag
}

// Update ...
func (*TpShopTagService) Update(p *TpShopTagUpdateParam) (*TpShopTagUpdateBack, error) {
	dao := &dao.TpShopTagDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopTagUpdateBack {
		*data,
	}
	 
	return &back, nil
}

