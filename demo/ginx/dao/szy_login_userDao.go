
package dao
import (
	"ginx/model"
)

// SzyLoginUserDao ...
type SzyLoginUserDao struct {
}

// Create 增
func (*SzyLoginUserDao) Create(m *model.SzyLoginUser) (*model.SzyLoginUser, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*SzyLoginUserDao) Delete(id int) error {
	err := model.DB.Delete(&model.SzyLoginUser{ID: id}).Error
	return err
}

// SelectByID 查
func (*SzyLoginUserDao) SelectByID(id int) (*model.SzyLoginUser, error) {

	var m model.SzyLoginUser
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*SzyLoginUserDao) Update(id int, update map[string]interface{}) (*model.SzyLoginUser, error) {

	var m model.SzyLoginUser
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
