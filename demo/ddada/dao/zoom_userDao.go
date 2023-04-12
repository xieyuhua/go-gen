
package dao
import (
	"ddada/model"
)

// ZoomUserDao ...
type ZoomUserDao struct {
}

// Create 增
func (*ZoomUserDao) Create(m *model.ZoomUser) (*model.ZoomUser, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomUserDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomUser{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomUserDao) SelectByID(id int) (*model.ZoomUser, error) {

	var m model.ZoomUser
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomUserDao) All() (*[]model.ZoomUser, error) {
	var ms []model.ZoomUser
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomUserDao) Update(id int, update map[string]interface{}) (*model.ZoomUser, error) {

	var m model.ZoomUser
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
