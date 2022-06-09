
package dao
import (
	"ginx/model"
)

// TpMemberLevelDao ...
type TpMemberLevelDao struct {
}

// Create 增
func (*TpMemberLevelDao) Create(m *model.TpMemberLevel) (*model.TpMemberLevel, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpMemberLevelDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpMemberLevel{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpMemberLevelDao) SelectByID(id int) (*model.TpMemberLevel, error) {

	var m model.TpMemberLevel
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpMemberLevelDao) Update(id int, update map[string]interface{}) (*model.TpMemberLevel, error) {

	var m model.TpMemberLevel
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
