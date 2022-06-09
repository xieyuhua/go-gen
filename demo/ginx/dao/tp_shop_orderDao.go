
package dao
import (
	"ginx/model"
)

// TpShopOrderDao ...
type TpShopOrderDao struct {
}

// Create 增
func (*TpShopOrderDao) Create(m *model.TpShopOrder) (*model.TpShopOrder, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopOrderDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopOrder{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopOrderDao) SelectByID(id int) (*model.TpShopOrder, error) {

	var m model.TpShopOrder
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopOrderDao) Update(id int, update map[string]interface{}) (*model.TpShopOrder, error) {

	var m model.TpShopOrder
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
