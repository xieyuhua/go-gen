
package dao
import (
	"ginx/model"
)

// TpShopTagDao ...
type TpShopTagDao struct {
}

// Create 增
func (*TpShopTagDao) Create(m *model.TpShopTag) (*model.TpShopTag, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpShopTagDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpShopTag{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpShopTagDao) SelectByID(id int) (*model.TpShopTag, error) {

	var m model.TpShopTag
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpShopTagDao) Update(id int, update map[string]interface{}) (*model.TpShopTag, error) {

	var m model.TpShopTag
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
