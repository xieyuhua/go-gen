
package dao
import (
	"ginx/model"
)

// TpMemberAccountDao ...
type TpMemberAccountDao struct {
}

// Create 增
func (*TpMemberAccountDao) Create(m *model.TpMemberAccount) (*model.TpMemberAccount, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpMemberAccountDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpMemberAccount{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpMemberAccountDao) SelectByID(id int) (*model.TpMemberAccount, error) {

	var m model.TpMemberAccount
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpMemberAccountDao) Update(id int, update map[string]interface{}) (*model.TpMemberAccount, error) {

	var m model.TpMemberAccount
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
