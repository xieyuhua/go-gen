
package dao
import (
	"ginx/model"
)

// TpCmsCategoryDao ...
type TpCmsCategoryDao struct {
}

// Create 增
func (*TpCmsCategoryDao) Create(m *model.TpCmsCategory) (*model.TpCmsCategory, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpCmsCategoryDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpCmsCategory{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpCmsCategoryDao) SelectByID(id int) (*model.TpCmsCategory, error) {

	var m model.TpCmsCategory
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpCmsCategoryDao) Update(id int, update map[string]interface{}) (*model.TpCmsCategory, error) {

	var m model.TpCmsCategory
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
