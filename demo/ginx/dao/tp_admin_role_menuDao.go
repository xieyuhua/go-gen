
package dao
import (
	"ginx/model"
)

// TpAdminRoleMenuDao ...
type TpAdminRoleMenuDao struct {
}

// Create 增
func (*TpAdminRoleMenuDao) Create(m *model.TpAdminRoleMenu) (*model.TpAdminRoleMenu, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAdminRoleMenuDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAdminRoleMenu{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAdminRoleMenuDao) SelectByID(id int) (*model.TpAdminRoleMenu, error) {

	var m model.TpAdminRoleMenu
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAdminRoleMenuDao) Update(id int, update map[string]interface{}) (*model.TpAdminRoleMenu, error) {

	var m model.TpAdminRoleMenu
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
