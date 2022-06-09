
package dao
import (
	"ginx/model"
)

// TpAdminRoleDao ...
type TpAdminRoleDao struct {
}

// Create 增
func (*TpAdminRoleDao) Create(m *model.TpAdminRole) (*model.TpAdminRole, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAdminRoleDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAdminRole{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAdminRoleDao) SelectByID(id int) (*model.TpAdminRole, error) {

	var m model.TpAdminRole
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAdminRoleDao) Update(id int, update map[string]interface{}) (*model.TpAdminRole, error) {

	var m model.TpAdminRole
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
