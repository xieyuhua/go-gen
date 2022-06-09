
package dao
import (
	"ginx/model"
)

// TpAdminRolePermissionDao ...
type TpAdminRolePermissionDao struct {
}

// Create 增
func (*TpAdminRolePermissionDao) Create(m *model.TpAdminRolePermission) (*model.TpAdminRolePermission, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAdminRolePermissionDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAdminRolePermission{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAdminRolePermissionDao) SelectByID(id int) (*model.TpAdminRolePermission, error) {

	var m model.TpAdminRolePermission
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAdminRolePermissionDao) Update(id int, update map[string]interface{}) (*model.TpAdminRolePermission, error) {

	var m model.TpAdminRolePermission
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
