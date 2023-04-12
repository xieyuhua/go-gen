
package dao
import (
	"xieyuhuas/model"
)

// ZoomDeployTaskDao ...
type ZoomDeployTaskDao struct {
}

// Create 增
func (*ZoomDeployTaskDao) Create(m *model.ZoomDeployTask) (*model.ZoomDeployTask, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomDeployTaskDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomDeployTask{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomDeployTaskDao) SelectByID(id int) (*model.ZoomDeployTask, error) {

	var m model.ZoomDeployTask
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomDeployTaskDao) All() (*[]model.ZoomDeployTask, error) {
	var ms []model.ZoomDeployTask
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomDeployTaskDao) Update(id int, update map[string]interface{}) (*model.ZoomDeployTask, error) {

	var m model.ZoomDeployTask
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
