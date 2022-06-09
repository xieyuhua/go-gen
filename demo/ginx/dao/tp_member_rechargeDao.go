
package dao
import (
	"ginx/model"
)

// TpMemberRechargeDao ...
type TpMemberRechargeDao struct {
}

// Create 增
func (*TpMemberRechargeDao) Create(m *model.TpMemberRecharge) (*model.TpMemberRecharge, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpMemberRechargeDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpMemberRecharge{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpMemberRechargeDao) SelectByID(id int) (*model.TpMemberRecharge, error) {

	var m model.TpMemberRecharge
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpMemberRechargeDao) Update(id int, update map[string]interface{}) (*model.TpMemberRecharge, error) {

	var m model.TpMemberRecharge
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
