

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpCmsTagService 服务
type TpCmsTagService struct {
}

// TpCmsTagCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpCmsTagCreateParam struct {
	model.TpCmsTag
}

// TpCmsTagCreateBack  返回参数
type TpCmsTagCreateBack struct {
	model.TpCmsTag
}

// Create 创建
func ( *TpCmsTagService) Create(p *TpCmsTagCreateParam) (*TpCmsTagCreateBack, error) {
	dao := &dao.TpCmsTagDao{}
	 
	data, err := dao.Create(&p.TpCmsTag)

	if err != nil {
		return nil, err

	}

	var back = TpCmsTagCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpCmsTagService) Delete(id int) error {
	dao := &dao.TpCmsTagDao{}
	return dao.Delete(id)
}

 

// TpCmsTagSelectBack  返回参数
type TpCmsTagSelectBack struct {
	model.TpCmsTag
}

// Select ...
func (*TpCmsTagService) Select(id int) (*TpCmsTagSelectBack, error) {
	dao := &dao.TpCmsTagDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpCmsTagSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpCmsTagUpdateParam   参数 
type TpCmsTagUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpCmsTagUpdateBack  返回参数
type TpCmsTagUpdateBack struct {
	model.TpCmsTag
}

// Update ...
func (*TpCmsTagService) Update(p *TpCmsTagUpdateParam) (*TpCmsTagUpdateBack, error) {
	dao := &dao.TpCmsTagDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpCmsTagUpdateBack {
		*data,
	}
	 
	return &back, nil
}

