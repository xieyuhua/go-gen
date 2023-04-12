
package dao
import (
	"xieyuhua/model"
)

// ZoomUserRoleDao ...
type ZoomUserRoleDao struct {
}

// Create 增
func (*ZoomUserRoleDao) Create(m *model.ZoomUserRole) (*model.ZoomUserRole, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomUserRoleDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomUserRole{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomUserRoleDao) SelectByID(id int) (*model.ZoomUserRole, error) {

	var m model.ZoomUserRole
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomUserRoleDao) All() (*[]model.ZoomUserRole, error) {
	var ms []model.ZoomUserRole
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomUserRoleDao) Update(id int, update map[string]interface{}) (*model.ZoomUserRole, error) {

	var m model.ZoomUserRole
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
