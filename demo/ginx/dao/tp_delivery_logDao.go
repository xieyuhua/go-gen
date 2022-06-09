
package dao
import (
	"ginx/model"
)

// TpDeliveryLogDao ...
type TpDeliveryLogDao struct {
}

// Create 增
func (*TpDeliveryLogDao) Create(m *model.TpDeliveryLog) (*model.TpDeliveryLog, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpDeliveryLogDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpDeliveryLog{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpDeliveryLogDao) SelectByID(id int) (*model.TpDeliveryLog, error) {

	var m model.TpDeliveryLog
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpDeliveryLogDao) Update(id int, update map[string]interface{}) (*model.TpDeliveryLog, error) {

	var m model.TpDeliveryLog
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
