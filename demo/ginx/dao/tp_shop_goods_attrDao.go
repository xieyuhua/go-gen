
package dao
import (
	"ginx/model"
)

// TpShopGoodsAttrDao ...
type TpShopGoodsAttrDao struct {
}

// Create 增
func (*TpShopGoodsAttrDao) Create(m *model.TpShopGoodsAttr) (*model.TpShopGoodsAttr, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopGoodsAttrDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopGoodsAttr{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopGoodsAttrDao) SelectByID(id int) (*model.TpShopGoodsAttr, error) {

	var m model.TpShopGoodsAttr
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopGoodsAttrDao) Update(id int, update map[string]interface{}) (*model.TpShopGoodsAttr, error) {

	var m model.TpShopGoodsAttr
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
