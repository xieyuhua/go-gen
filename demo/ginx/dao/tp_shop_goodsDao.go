
package dao
import (
	"ginx/model"
)

// TpShopGoodDao ...
type TpShopGoodDao struct {
}

// Create 增
func (*TpShopGoodDao) Create(m *model.TpShopGood) (*model.TpShopGood, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopGoodDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopGood{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopGoodDao) SelectByID(id int) (*model.TpShopGood, error) {

	var m model.TpShopGood
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopGoodDao) Update(id int, update map[string]interface{}) (*model.TpShopGood, error) {

	var m model.TpShopGood
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
