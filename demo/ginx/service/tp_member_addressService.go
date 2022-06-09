

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpMemberAddresService 服务
type TpMemberAddresService struct {
}

// TpMemberAddresCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpMemberAddresCreateParam struct {
	model.TpMemberAddres
}

// TpMemberAddresCreateBack  返回参数
type TpMemberAddresCreateBack struct {
	model.TpMemberAddres
}

// Create 创建
func ( *TpMemberAddresService) Create(p *TpMemberAddresCreateParam) (*TpMemberAddresCreateBack, error) {
	dao := &dao.TpMemberAddresDao{}
	 
	data, err := dao.Create(&p.TpMemberAddres)

	if err != nil {
		return nil, err

	}

	var back = TpMemberAddresCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpMemberAddresService) Delete(id int) error {
	dao := &dao.TpMemberAddresDao{}
	return dao.Delete(id)
}

 

// TpMemberAddresSelectBack  返回参数
type TpMemberAddresSelectBack struct {
	model.TpMemberAddres
}

// Select ...
func (*TpMemberAddresService) Select(id int) (*TpMemberAddresSelectBack, error) {
	dao := &dao.TpMemberAddresDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpMemberAddresSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpMemberAddresUpdateParam   参数 
type TpMemberAddresUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpMemberAddresUpdateBack  返回参数
type TpMemberAddresUpdateBack struct {
	model.TpMemberAddres
}

// Update ...
func (*TpMemberAddresService) Update(p *TpMemberAddresUpdateParam) (*TpMemberAddresUpdateBack, error) {
	dao := &dao.TpMemberAddresDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpMemberAddresUpdateBack {
		*data,
	}
	 
	return &back, nil
}

