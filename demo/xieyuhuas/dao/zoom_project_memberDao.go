
package dao
import (
	"xieyuhuas/model"
)

// ZoomProjectMemberDao ...
type ZoomProjectMemberDao struct {
}

// Create 增
func (*ZoomProjectMemberDao) Create(m *model.ZoomProjectMember) (*model.ZoomProjectMember, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomProjectMemberDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomProjectMember{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomProjectMemberDao) SelectByID(id int) (*model.ZoomProjectMember, error) {

	var m model.ZoomProjectMember
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomProjectMemberDao) All() (*[]model.ZoomProjectMember, error) {
	var ms []model.ZoomProjectMember
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomProjectMemberDao) Update(id int, update map[string]interface{}) (*model.ZoomProjectMember, error) {

	var m model.ZoomProjectMember
	err := model.DB.Model(&m).Where("id = ?", id).Updates(update).Error

	if err != nil {
		return nil, err
	}
	err = model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}
