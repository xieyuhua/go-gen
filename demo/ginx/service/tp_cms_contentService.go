

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpCmsContentService 服务
type TpCmsContentService struct {
}

// TpCmsContentCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpCmsContentCreateParam struct {
	model.TpCmsContent
}

// TpCmsContentCreateBack  返回参数
type TpCmsContentCreateBack struct {
	model.TpCmsContent
}

// Create 创建
func ( *TpCmsContentService) Create(p *TpCmsContentCreateParam) (*TpCmsContentCreateBack, error) {
	dao := &dao.TpCmsContentDao{}
	 
	data, err := dao.Create(&p.TpCmsContent)

	if err != nil {
		return nil, err

	}

	var back = TpCmsContentCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpCmsContentService) Delete(id int) error {
	dao := &dao.TpCmsContentDao{}
	return dao.Delete(id)
}

 

// TpCmsContentSelectBack  返回参数
type TpCmsContentSelectBack struct {
	model.TpCmsContent
}

// Select ...
func (*TpCmsContentService) Select(id int) (*TpCmsContentSelectBack, error) {
	dao := &dao.TpCmsContentDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpCmsContentSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpCmsContentUpdateParam   参数 
type TpCmsContentUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpCmsContentUpdateBack  返回参数
type TpCmsContentUpdateBack struct {
	model.TpCmsContent
}

// Update ...
func (*TpCmsContentService) Update(p *TpCmsContentUpdateParam) (*TpCmsContentUpdateBack, error) {
	dao := &dao.TpCmsContentDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpCmsContentUpdateBack {
		*data,
	}
	 
	return &back, nil
}

