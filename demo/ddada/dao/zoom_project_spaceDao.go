
package dao
import (
	"ddada/model"
)

// ZoomProjectSpaceDao ...
type ZoomProjectSpaceDao struct {
}

// Create 增
func (*ZoomProjectSpaceDao) Create(m *model.ZoomProjectSpace) (*model.ZoomProjectSpace, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomProjectSpaceDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomProjectSpace{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomProjectSpaceDao) SelectByID(id int) (*model.ZoomProjectSpace, error) {

	var m model.ZoomProjectSpace
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomProjectSpaceDao) All() (*[]model.ZoomProjectSpace, error) {
	var ms []model.ZoomProjectSpace
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomProjectSpaceDao) Update(id int, update map[string]interface{}) (*model.ZoomProjectSpace, error) {

	var m model.ZoomProjectSpace
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
