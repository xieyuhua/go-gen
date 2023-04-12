
package dao
import (
	"yuhua/model"
)

// ZoomDeployApplyDao ...
type ZoomDeployApplyDao struct {
}

// Create 增
func (*ZoomDeployApplyDao) Create(m *model.ZoomDeployApply) (*model.ZoomDeployApply, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomDeployApplyDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomDeployApply{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomDeployApplyDao) SelectByID(id int) (*model.ZoomDeployApply, error) {

	var m model.ZoomDeployApply
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomDeployApplyDao) Update(id int, update map[string]interface{}) (*model.ZoomDeployApply, error) {

	var m model.ZoomDeployApply
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
