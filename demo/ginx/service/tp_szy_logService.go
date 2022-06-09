

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpSzyLogService 服务
type TpSzyLogService struct {
}

// TpSzyLogCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpSzyLogCreateParam struct {
	model.TpSzyLog
}

// TpSzyLogCreateBack  返回参数
type TpSzyLogCreateBack struct {
	model.TpSzyLog
}

// Create 创建
func ( *TpSzyLogService) Create(p *TpSzyLogCreateParam) (*TpSzyLogCreateBack, error) {
	dao := &dao.TpSzyLogDao{}
	 
	data, err := dao.Create(&p.TpSzyLog)

	if err != nil {
		return nil, err

	}

	var back = TpSzyLogCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpSzyLogService) Delete(id int) error {
	dao := &dao.TpSzyLogDao{}
	return dao.Delete(id)
}

 

// TpSzyLogSelectBack  返回参数
type TpSzyLogSelectBack struct {
	model.TpSzyLog
}

// Select ...
func (*TpSzyLogService) Select(id int) (*TpSzyLogSelectBack, error) {
	dao := &dao.TpSzyLogDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpSzyLogSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpSzyLogUpdateParam   参数 
type TpSzyLogUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpSzyLogUpdateBack  返回参数
type TpSzyLogUpdateBack struct {
	model.TpSzyLog
}

// Update ...
func (*TpSzyLogService) Update(p *TpSzyLogUpdateParam) (*TpSzyLogUpdateBack, error) {
	dao := &dao.TpSzyLogDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpSzyLogUpdateBack {
		*data,
	}
	 
	return &back, nil
}

