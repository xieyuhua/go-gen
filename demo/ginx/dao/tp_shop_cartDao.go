
package dao
import (
	"ginx/model"
)

// TpShopCartDao ...
type TpShopCartDao struct {
}

// Create 增
func (*TpShopCartDao) Create(m *model.TpShopCart) (*model.TpShopCart, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopCartDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopCart{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopCartDao) SelectByID(id int) (*model.TpShopCart, error) {

	var m model.TpShopCart
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopCartDao) Update(id int, update map[string]interface{}) (*model.TpShopCart, error) {

	var m model.TpShopCart
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
