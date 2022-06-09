
package dao
import (
	"ginx/model"
)

// TpShippingAreaItemDao ...
type TpShippingAreaItemDao struct {
}

// Create 增
func (*TpShippingAreaItemDao) Create(m *model.TpShippingAreaItem) (*model.TpShippingAreaItem, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShippingAreaItemDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShippingAreaItem{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShippingAreaItemDao) SelectByID(id int) (*model.TpShippingAreaItem, error) {

	var m model.TpShippingAreaItem
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShippingAreaItemDao) Update(id int, update map[string]interface{}) (*model.TpShippingAreaItem, error) {

	var m model.TpShippingAreaItem
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
