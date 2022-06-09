
package dao
import (
	"ginx/model"
)

// TpShopBrandDao ...
type TpShopBrandDao struct {
}

// Create 增
func (*TpShopBrandDao) Create(m *model.TpShopBrand) (*model.TpShopBrand, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopBrandDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopBrand{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopBrandDao) SelectByID(id int) (*model.TpShopBrand, error) {

	var m model.TpShopBrand
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopBrandDao) Update(id int, update map[string]interface{}) (*model.TpShopBrand, error) {

	var m model.TpShopBrand
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
