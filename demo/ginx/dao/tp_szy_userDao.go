
package dao
import (
	"ginx/model"
)

// TpSzyUserDao ...
type TpSzyUserDao struct {
}

// Create 增
func (*TpSzyUserDao) Create(m *model.TpSzyUser) (*model.TpSzyUser, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpSzyUserDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpSzyUser{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpSzyUserDao) SelectByID(id int) (*model.TpSzyUser, error) {

	var m model.TpSzyUser
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpSzyUserDao) Update(id int, update map[string]interface{}) (*model.TpSzyUser, error) {

	var m model.TpSzyUser
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
