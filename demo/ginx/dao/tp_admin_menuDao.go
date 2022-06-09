
package dao
import (
	"ginx/model"
)

// TpAdminMenuDao ...
type TpAdminMenuDao struct {
}

// Create 增
func (*TpAdminMenuDao) Create(m *model.TpAdminMenu) (*model.TpAdminMenu, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAdminMenuDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAdminMenu{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAdminMenuDao) SelectByID(id int) (*model.TpAdminMenu, error) {

	var m model.TpAdminMenu
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAdminMenuDao) Update(id int, update map[string]interface{}) (*model.TpAdminMenu, error) {

	var m model.TpAdminMenu
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
