
package dao
import (
	"ginx/model"
)

// TpShopCategoryDao ...
type TpShopCategoryDao struct {
}

// Create 增
func (*TpShopCategoryDao) Create(m *model.TpShopCategory) (*model.TpShopCategory, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopCategoryDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopCategory{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopCategoryDao) SelectByID(id int) (*model.TpShopCategory, error) {

	var m model.TpShopCategory
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopCategoryDao) Update(id int, update map[string]interface{}) (*model.TpShopCategory, error) {

	var m model.TpShopCategory
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
