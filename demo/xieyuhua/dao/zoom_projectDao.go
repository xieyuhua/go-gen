
package dao
import (
	"xieyuhua/model"
)

// ZoomProjectDao ...
type ZoomProjectDao struct {
}

// Create 增
func (*ZoomProjectDao) Create(m *model.ZoomProject) (*model.ZoomProject, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomProjectDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomProject{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomProjectDao) SelectByID(id int) (*model.ZoomProject, error) {

	var m model.ZoomProject
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomProjectDao) All() (*[]model.ZoomProject, error) {
	var ms []model.ZoomProject
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomProjectDao) Update(id int, update map[string]interface{}) (*model.ZoomProject, error) {

	var m model.ZoomProject
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
