
package dao
import (
	"ddada/model"
)

// ZoomUserTokenDao ...
type ZoomUserTokenDao struct {
}

// Create 增
func (*ZoomUserTokenDao) Create(m *model.ZoomUserToken) (*model.ZoomUserToken, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomUserTokenDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomUserToken{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomUserTokenDao) SelectByID(id int) (*model.ZoomUserToken, error) {

	var m model.ZoomUserToken
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomUserTokenDao) All() (*[]model.ZoomUserToken, error) {
	var ms []model.ZoomUserToken
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomUserTokenDao) Update(id int, update map[string]interface{}) (*model.ZoomUserToken, error) {

	var m model.ZoomUserToken
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
