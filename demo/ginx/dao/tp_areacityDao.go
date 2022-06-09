
package dao
import (
	"ginx/model"
)

// TpAreacityDao ...
type TpAreacityDao struct {
}

// Create 增
func (*TpAreacityDao) Create(m *model.TpAreacity) (*model.TpAreacity, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAreacityDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAreacity{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAreacityDao) SelectByID(id int) (*model.TpAreacity, error) {

	var m model.TpAreacity
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAreacityDao) Update(id int, update map[string]interface{}) (*model.TpAreacity, error) {

	var m model.TpAreacity
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
