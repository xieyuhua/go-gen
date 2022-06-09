
package dao
import (
	"ginx/model"
)

// TpShopOrderActionDao ...
type TpShopOrderActionDao struct {
}

// Create 增
func (*TpShopOrderActionDao) Create(m *model.TpShopOrderAction) (*model.TpShopOrderAction, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopOrderActionDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopOrderAction{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopOrderActionDao) SelectByID(id int) (*model.TpShopOrderAction, error) {

	var m model.TpShopOrderAction
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopOrderActionDao) Update(id int, update map[string]interface{}) (*model.TpShopOrderAction, error) {

	var m model.TpShopOrderAction
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
