
package dao
import (
	"xieyuhuas/model"
)

// ZoomServerGroupDao ...
type ZoomServerGroupDao struct {
}

// Create 增
func (*ZoomServerGroupDao) Create(m *model.ZoomServerGroup) (*model.ZoomServerGroup, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomServerGroupDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomServerGroup{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomServerGroupDao) SelectByID(id int) (*model.ZoomServerGroup, error) {

	var m model.ZoomServerGroup
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomServerGroupDao) All() (*[]model.ZoomServerGroup, error) {
	var ms []model.ZoomServerGroup
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomServerGroupDao) Update(id int, update map[string]interface{}) (*model.ZoomServerGroup, error) {

	var m model.ZoomServerGroup
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
