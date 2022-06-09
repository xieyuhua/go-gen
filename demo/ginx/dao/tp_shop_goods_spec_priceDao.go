
package dao
import (
	"ginx/model"
)

// TpShopGoodsSpecPriceDao ...
type TpShopGoodsSpecPriceDao struct {
}

// Create 增
func (*TpShopGoodsSpecPriceDao) Create(m *model.TpShopGoodsSpecPrice) (*model.TpShopGoodsSpecPrice, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopGoodsSpecPriceDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopGoodsSpecPrice{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopGoodsSpecPriceDao) SelectByID(id int) (*model.TpShopGoodsSpecPrice, error) {

	var m model.TpShopGoodsSpecPrice
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopGoodsSpecPriceDao) Update(id int, update map[string]interface{}) (*model.TpShopGoodsSpecPrice, error) {

	var m model.TpShopGoodsSpecPrice
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
