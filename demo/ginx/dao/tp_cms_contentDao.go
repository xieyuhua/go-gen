
package dao
import (
	"ginx/model"
)

// TpCmsContentDao ...
type TpCmsContentDao struct {
}

// Create 增
func (*TpCmsContentDao) Create(m *model.TpCmsContent) (*model.TpCmsContent, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpCmsContentDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpCmsContent{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpCmsContentDao) SelectByID(id int) (*model.TpCmsContent, error) {

	var m model.TpCmsContent
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpCmsContentDao) Update(id int, update map[string]interface{}) (*model.TpCmsContent, error) {

	var m model.TpCmsContent
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
