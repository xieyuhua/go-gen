
package dao
import (
	"ginx/model"
)

// TpShopCouponListDao ...
type TpShopCouponListDao struct {
}

// Create 增
func (*TpShopCouponListDao) Create(m *model.TpShopCouponList) (*model.TpShopCouponList, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopCouponListDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopCouponList{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopCouponListDao) SelectByID(id int) (*model.TpShopCouponList, error) {

	var m model.TpShopCouponList
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopCouponListDao) Update(id int, update map[string]interface{}) (*model.TpShopCouponList, error) {

	var m model.TpShopCouponList
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
