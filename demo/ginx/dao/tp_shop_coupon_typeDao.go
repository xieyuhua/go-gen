
package dao
import (
	"ginx/model"
)

// TpShopCouponTypeDao ...
type TpShopCouponTypeDao struct {
}

// Create 增
func (*TpShopCouponTypeDao) Create(m *model.TpShopCouponType) (*model.TpShopCouponType, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopCouponTypeDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopCouponType{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopCouponTypeDao) SelectByID(id int) (*model.TpShopCouponType, error) {

	var m model.TpShopCouponType
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopCouponTypeDao) Update(id int, update map[string]interface{}) (*model.TpShopCouponType, error) {

	var m model.TpShopCouponType
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
