

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpTableRelationService 服务
type TpTableRelationService struct {
}

// TpTableRelationCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpTableRelationCreateParam struct {
	model.TpTableRelation
}

// TpTableRelationCreateBack  返回参数
type TpTableRelationCreateBack struct {
	model.TpTableRelation
}

// Create 创建
func ( *TpTableRelationService) Create(p *TpTableRelationCreateParam) (*TpTableRelationCreateBack, error) {
	dao := &dao.TpTableRelationDao{}
	 
	data, err := dao.Create(&p.TpTableRelation)

	if err != nil {
		return nil, err

	}

	var back = TpTableRelationCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpTableRelationService) Delete(id int) error {
	dao := &dao.TpTableRelationDao{}
	return dao.Delete(id)
}

 

// TpTableRelationSelectBack  返回参数
type TpTableRelationSelectBack struct {
	model.TpTableRelation
}

// Select ...
func (*TpTableRelationService) Select(id int) (*TpTableRelationSelectBack, error) {
	dao := &dao.TpTableRelationDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpTableRelationSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpTableRelationUpdateParam   参数 
type TpTableRelationUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpTableRelationUpdateBack  返回参数
type TpTableRelationUpdateBack struct {
	model.TpTableRelation
}

// Update ...
func (*TpTableRelationService) Update(p *TpTableRelationUpdateParam) (*TpTableRelationUpdateBack, error) {
	dao := &dao.TpTableRelationDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpTableRelationUpdateBack {
		*data,
	}
	 
	return &back, nil
}

