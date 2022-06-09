

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpMemberLevelService 服务
type TpMemberLevelService struct {
}

// TpMemberLevelCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpMemberLevelCreateParam struct {
	model.TpMemberLevel
}

// TpMemberLevelCreateBack  返回参数
type TpMemberLevelCreateBack struct {
	model.TpMemberLevel
}

// Create 创建
func ( *TpMemberLevelService) Create(p *TpMemberLevelCreateParam) (*TpMemberLevelCreateBack, error) {
	dao := &dao.TpMemberLevelDao{}
	 
	data, err := dao.Create(&p.TpMemberLevel)

	if err != nil {
		return nil, err

	}

	var back = TpMemberLevelCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpMemberLevelService) Delete(id int) error {
	dao := &dao.TpMemberLevelDao{}
	return dao.Delete(id)
}

 

// TpMemberLevelSelectBack  返回参数
type TpMemberLevelSelectBack struct {
	model.TpMemberLevel
}

// Select ...
func (*TpMemberLevelService) Select(id int) (*TpMemberLevelSelectBack, error) {
	dao := &dao.TpMemberLevelDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpMemberLevelSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpMemberLevelUpdateParam   参数 
type TpMemberLevelUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpMemberLevelUpdateBack  返回参数
type TpMemberLevelUpdateBack struct {
	model.TpMemberLevel
}

// Update ...
func (*TpMemberLevelService) Update(p *TpMemberLevelUpdateParam) (*TpMemberLevelUpdateBack, error) {
	dao := &dao.TpMemberLevelDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpMemberLevelUpdateBack {
		*data,
	}
	 
	return &back, nil
}

