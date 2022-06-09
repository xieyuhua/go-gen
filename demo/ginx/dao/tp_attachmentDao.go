
package dao
import (
	"ginx/model"
)

// TpAttachmentDao ...
type TpAttachmentDao struct {
}

// Create 增
func (*TpAttachmentDao) Create(m *model.TpAttachment) (*model.TpAttachment, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAttachmentDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAttachment{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAttachmentDao) SelectByID(id int) (*model.TpAttachment, error) {

	var m model.TpAttachment
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAttachmentDao) Update(id int, update map[string]interface{}) (*model.TpAttachment, error) {

	var m model.TpAttachment
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
