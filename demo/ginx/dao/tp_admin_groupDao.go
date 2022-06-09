
package dao
import (
	"ginx/model"
)

// TpAdminGroupDao ...
type TpAdminGroupDao struct {
}

// Create 增
func (*TpAdminGroupDao) Create(m *model.TpAdminGroup) (*model.TpAdminGroup, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAdminGroupDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAdminGroup{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAdminGroupDao) SelectByID(id int) (*model.TpAdminGroup, error) {

	var m model.TpAdminGroup
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAdminGroupDao) Update(id int, update map[string]interface{}) (*model.TpAdminGroup, error) {

	var m model.TpAdminGroup
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
