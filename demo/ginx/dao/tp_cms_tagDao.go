
package dao
import (
	"ginx/model"
)

// TpCmsTagDao ...
type TpCmsTagDao struct {
}

// Create 增
func (*TpCmsTagDao) Create(m *model.TpCmsTag) (*model.TpCmsTag, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpCmsTagDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpCmsTag{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpCmsTagDao) SelectByID(id int) (*model.TpCmsTag, error) {

	var m model.TpCmsTag
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpCmsTagDao) Update(id int, update map[string]interface{}) (*model.TpCmsTag, error) {

	var m model.TpCmsTag
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
