
package dao
import (
	"ginx/model"
)

// TpMemberAddresDao ...
type TpMemberAddresDao struct {
}

// Create 增
func (*TpMemberAddresDao) Create(m *model.TpMemberAddres) (*model.TpMemberAddres, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpMemberAddresDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpMemberAddres{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpMemberAddresDao) SelectByID(id int) (*model.TpMemberAddres, error) {

	var m model.TpMemberAddres
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpMemberAddresDao) Update(id int, update map[string]interface{}) (*model.TpMemberAddres, error) {

	var m model.TpMemberAddres
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
