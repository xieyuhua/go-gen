

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpMemberService 服务
type TpMemberService struct {
}

// TpMemberCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpMemberCreateParam struct {
	model.TpMember
}

// TpMemberCreateBack  返回参数
type TpMemberCreateBack struct {
	model.TpMember
}

// Create 创建
func ( *TpMemberService) Create(p *TpMemberCreateParam) (*TpMemberCreateBack, error) {
	dao := &dao.TpMemberDao{}
	 
	data, err := dao.Create(&p.TpMember)

	if err != nil {
		return nil, err

	}

	var back = TpMemberCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpMemberService) Delete(id int) error {
	dao := &dao.TpMemberDao{}
	return dao.Delete(id)
}

 

// TpMemberSelectBack  返回参数
type TpMemberSelectBack struct {
	model.TpMember
}

// Select ...
func (*TpMemberService) Select(id int) (*TpMemberSelectBack, error) {
	dao := &dao.TpMemberDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpMemberSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpMemberUpdateParam   参数 
type TpMemberUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpMemberUpdateBack  返回参数
type TpMemberUpdateBack struct {
	model.TpMember
}

// Update ...
func (*TpMemberService) Update(p *TpMemberUpdateParam) (*TpMemberUpdateBack, error) {
	dao := &dao.TpMemberDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpMemberUpdateBack {
		*data,
	}
	 
	return &back, nil
}

