

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAgentLevelService 服务
type TpAgentLevelService struct {
}

// TpAgentLevelCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAgentLevelCreateParam struct {
	model.TpAgentLevel
}

// TpAgentLevelCreateBack  返回参数
type TpAgentLevelCreateBack struct {
	model.TpAgentLevel
}

// Create 创建
func ( *TpAgentLevelService) Create(p *TpAgentLevelCreateParam) (*TpAgentLevelCreateBack, error) {
	dao := &dao.TpAgentLevelDao{}
	 
	data, err := dao.Create(&p.TpAgentLevel)

	if err != nil {
		return nil, err

	}

	var back = TpAgentLevelCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAgentLevelService) Delete(id int) error {
	dao := &dao.TpAgentLevelDao{}
	return dao.Delete(id)
}

 

// TpAgentLevelSelectBack  返回参数
type TpAgentLevelSelectBack struct {
	model.TpAgentLevel
}

// Select ...
func (*TpAgentLevelService) Select(id int) (*TpAgentLevelSelectBack, error) {
	dao := &dao.TpAgentLevelDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAgentLevelSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAgentLevelUpdateParam   参数 
type TpAgentLevelUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAgentLevelUpdateBack  返回参数
type TpAgentLevelUpdateBack struct {
	model.TpAgentLevel
}

// Update ...
func (*TpAgentLevelService) Update(p *TpAgentLevelUpdateParam) (*TpAgentLevelUpdateBack, error) {
	dao := &dao.TpAgentLevelDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAgentLevelUpdateBack {
		*data,
	}
	 
	return &back, nil
}

