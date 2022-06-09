
package dao
import (
	"ginx/model"
)

// TpShopGoodsSpecValueDao ...
type TpShopGoodsSpecValueDao struct {
}

// Create 增
func (*TpShopGoodsSpecValueDao) Create(m *model.TpShopGoodsSpecValue) (*model.TpShopGoodsSpecValue, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopGoodsSpecValueDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopGoodsSpecValue{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopGoodsSpecValueDao) SelectByID(id int) (*model.TpShopGoodsSpecValue, error) {

	var m model.TpShopGoodsSpecValue
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopGoodsSpecValueDao) Update(id int, update map[string]interface{}) (*model.TpShopGoodsSpecValue, error) {

	var m model.TpShopGoodsSpecValue
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
