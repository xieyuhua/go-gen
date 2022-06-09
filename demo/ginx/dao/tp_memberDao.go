
package dao
import (
	"ginx/model"
)

// TpMemberDao ...
type TpMemberDao struct {
}

// Create 增
func (*TpMemberDao) Create(m *model.TpMember) (*model.TpMember, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpMemberDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpMember{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpMemberDao) SelectByID(id int) (*model.TpMember, error) {

	var m model.TpMember
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpMemberDao) Update(id int, update map[string]interface{}) (*model.TpMember, error) {

	var m model.TpMember
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
