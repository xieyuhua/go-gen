
package dao
import (
	"ginx/model"
)

// TpShippingAreaDao ...
type TpShippingAreaDao struct {
}

// Create 增
func (*TpShippingAreaDao) Create(m *model.TpShippingArea) (*model.TpShippingArea, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShippingAreaDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShippingArea{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShippingAreaDao) SelectByID(id int) (*model.TpShippingArea, error) {

	var m model.TpShippingArea
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShippingAreaDao) Update(id int, update map[string]interface{}) (*model.TpShippingArea, error) {

	var m model.TpShippingArea
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
