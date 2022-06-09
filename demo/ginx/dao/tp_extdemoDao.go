
package dao
import (
	"ginx/model"
)

// TpExtdemoDao ...
type TpExtdemoDao struct {
}

// Create 增
func (*TpExtdemoDao) Create(m *model.TpExtdemo) (*model.TpExtdemo, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpExtdemoDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpExtdemo{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpExtdemoDao) SelectByID(id int) (*model.TpExtdemo, error) {

	var m model.TpExtdemo
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpExtdemoDao) Update(id int, update map[string]interface{}) (*model.TpExtdemo, error) {

	var m model.TpExtdemo
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
