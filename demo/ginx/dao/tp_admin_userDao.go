
package dao
import (
	"ginx/model"
)

// TpAdminUserDao ...
type TpAdminUserDao struct {
}

// Create 增
func (*TpAdminUserDao) Create(m *model.TpAdminUser) (*model.TpAdminUser, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAdminUserDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAdminUser{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAdminUserDao) SelectByID(id int) (*model.TpAdminUser, error) {

	var m model.TpAdminUser
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAdminUserDao) Update(id int, update map[string]interface{}) (*model.TpAdminUser, error) {

	var m model.TpAdminUser
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
