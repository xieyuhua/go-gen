

package service
import (
	"xieyuhuas/dao"
	"xieyuhuas/model"
)

// ZoomProjectMemberService 服务
type ZoomProjectMemberService struct {
}

// ZoomProjectMemberCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type ZoomProjectMemberCreateParam struct {
	model.ZoomProjectMember
}

// ZoomProjectMemberCreateBack  返回参数
type ZoomProjectMemberCreateBack struct {
	model.ZoomProjectMember
}

// Create 创建
func ( *ZoomProjectMemberService) Create(p *ZoomProjectMemberCreateParam) (*ZoomProjectMemberCreateBack, error) {
	dao := &dao.ZoomProjectMemberDao{}
	 
	data, err := dao.Create(&p.ZoomProjectMember)

	if err != nil {
		return nil, err

	}

	var back = ZoomProjectMemberCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *ZoomProjectMemberService) Delete(id int) error {
	dao := &dao.ZoomProjectMemberDao{}
	return dao.Delete(id)
}

 

// ZoomProjectMemberSelectBack  返回参数
type ZoomProjectMemberSelectBack struct {
	model.ZoomProjectMember
}

// All ...
func (*ZoomProjectMemberService) All() (*[]model.ZoomProjectMember, error) {
	dao := &dao.ZoomProjectMemberDao{}
	data, err := dao.All()
	if err != nil {
		return nil, err
	}
	return data, nil
}


// Select ...
func (*ZoomProjectMemberService) Select(id int) (*ZoomProjectMemberSelectBack, error) {
	dao := &dao.ZoomProjectMemberDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = ZoomProjectMemberSelectBack{
		*data,
	}
	  
	return &back, nil
}

// ZoomProjectMemberUpdateParam   参数 
type ZoomProjectMemberUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// ZoomProjectMemberUpdateBack  返回参数
type ZoomProjectMemberUpdateBack struct {
	model.ZoomProjectMember
}

// Update ...
func (*ZoomProjectMemberService) Update(p *ZoomProjectMemberUpdateParam) (*ZoomProjectMemberUpdateBack, error) {
	dao := &dao.ZoomProjectMemberDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = ZoomProjectMemberUpdateBack {
		*data,
	}
	 
	return &back, nil
}

