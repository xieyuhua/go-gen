
package dao
import (
	"ginx/model"
)

// TpExtensionDao ...
type TpExtensionDao struct {
}

// Create 增
func (*TpExtensionDao) Create(m *model.TpExtension) (*model.TpExtension, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpExtensionDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpExtension{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpExtensionDao) SelectByID(id int) (*model.TpExtension, error) {

	var m model.TpExtension
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpExtensionDao) Update(id int, update map[string]interface{}) (*model.TpExtension, error) {

	var m model.TpExtension
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
