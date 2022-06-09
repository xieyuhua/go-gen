

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpExtdemoService 服务
type TpExtdemoService struct {
}

// TpExtdemoCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpExtdemoCreateParam struct {
	model.TpExtdemo
}

// TpExtdemoCreateBack  返回参数
type TpExtdemoCreateBack struct {
	model.TpExtdemo
}

// Create 创建
func ( *TpExtdemoService) Create(p *TpExtdemoCreateParam) (*TpExtdemoCreateBack, error) {
	dao := &dao.TpExtdemoDao{}
	 
	data, err := dao.Create(&p.TpExtdemo)

	if err != nil {
		return nil, err

	}

	var back = TpExtdemoCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpExtdemoService) Delete(id int) error {
	dao := &dao.TpExtdemoDao{}
	return dao.Delete(id)
}

 

// TpExtdemoSelectBack  返回参数
type TpExtdemoSelectBack struct {
	model.TpExtdemo
}

// Select ...
func (*TpExtdemoService) Select(id int) (*TpExtdemoSelectBack, error) {
	dao := &dao.TpExtdemoDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpExtdemoSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpExtdemoUpdateParam   参数 
type TpExtdemoUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpExtdemoUpdateBack  返回参数
type TpExtdemoUpdateBack struct {
	model.TpExtdemo
}

// Update ...
func (*TpExtdemoService) Update(p *TpExtdemoUpdateParam) (*TpExtdemoUpdateBack, error) {
	dao := &dao.TpExtdemoDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpExtdemoUpdateBack {
		*data,
	}
	 
	return &back, nil
}

