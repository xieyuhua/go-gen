
package dao
import (
	"ginx/model"
)

// TpWebConfigDao ...
type TpWebConfigDao struct {
}

// Create 增
func (*TpWebConfigDao) Create(m *model.TpWebConfig) (*model.TpWebConfig, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpWebConfigDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpWebConfig{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpWebConfigDao) SelectByID(id int) (*model.TpWebConfig, error) {

	var m model.TpWebConfig
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpWebConfigDao) Update(id int, update map[string]interface{}) (*model.TpWebConfig, error) {

	var m model.TpWebConfig
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
