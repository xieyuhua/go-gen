
package dao
import (
	"ginx/model"
)

// TpShopGoodsSpecDao ...
type TpShopGoodsSpecDao struct {
}

// Create 增
func (*TpShopGoodsSpecDao) Create(m *model.TpShopGoodsSpec) (*model.TpShopGoodsSpec, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopGoodsSpecDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopGoodsSpec{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopGoodsSpecDao) SelectByID(id int) (*model.TpShopGoodsSpec, error) {

	var m model.TpShopGoodsSpec
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopGoodsSpecDao) Update(id int, update map[string]interface{}) (*model.TpShopGoodsSpec, error) {

	var m model.TpShopGoodsSpec
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
