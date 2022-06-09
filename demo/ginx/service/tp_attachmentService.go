

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAttachmentService 服务
type TpAttachmentService struct {
}

// TpAttachmentCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAttachmentCreateParam struct {
	model.TpAttachment
}

// TpAttachmentCreateBack  返回参数
type TpAttachmentCreateBack struct {
	model.TpAttachment
}

// Create 创建
func ( *TpAttachmentService) Create(p *TpAttachmentCreateParam) (*TpAttachmentCreateBack, error) {
	dao := &dao.TpAttachmentDao{}
	 
	data, err := dao.Create(&p.TpAttachment)

	if err != nil {
		return nil, err

	}

	var back = TpAttachmentCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAttachmentService) Delete(id int) error {
	dao := &dao.TpAttachmentDao{}
	return dao.Delete(id)
}

 

// TpAttachmentSelectBack  返回参数
type TpAttachmentSelectBack struct {
	model.TpAttachment
}

// Select ...
func (*TpAttachmentService) Select(id int) (*TpAttachmentSelectBack, error) {
	dao := &dao.TpAttachmentDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAttachmentSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAttachmentUpdateParam   参数 
type TpAttachmentUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAttachmentUpdateBack  返回参数
type TpAttachmentUpdateBack struct {
	model.TpAttachment
}

// Update ...
func (*TpAttachmentService) Update(p *TpAttachmentUpdateParam) (*TpAttachmentUpdateBack, error) {
	dao := &dao.TpAttachmentDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAttachmentUpdateBack {
		*data,
	}
	 
	return &back, nil
}

