
package dao
import (
	"xieyuhuas/model"
)

// ZoomServerDao ...
type ZoomServerDao struct {
}

// Create 增
func (*ZoomServerDao) Create(m *model.ZoomServer) (*model.ZoomServer, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomServerDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomServer{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomServerDao) SelectByID(id int) (*model.ZoomServer, error) {

	var m model.ZoomServer
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomServerDao) All() (*[]model.ZoomServer, error) {
	var ms []model.ZoomServer
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomServerDao) Update(id int, update map[string]interface{}) (*model.ZoomServer, error) {

	var m model.ZoomServer
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
