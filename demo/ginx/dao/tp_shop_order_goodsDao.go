
package dao
import (
	"ginx/model"
)

// TpShopOrderGoodDao ...
type TpShopOrderGoodDao struct {
}

// Create 增
func (*TpShopOrderGoodDao) Create(m *model.TpShopOrderGood) (*model.TpShopOrderGood, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopOrderGoodDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopOrderGood{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopOrderGoodDao) SelectByID(id int) (*model.TpShopOrderGood, error) {

	var m model.TpShopOrderGood
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopOrderGoodDao) Update(id int, update map[string]interface{}) (*model.TpShopOrderGood, error) {

	var m model.TpShopOrderGood
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
