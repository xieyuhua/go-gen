
package dao
import (
	"ginx/model"
)

// TpTableRelationDao ...
type TpTableRelationDao struct {
}

// Create 增
func (*TpTableRelationDao) Create(m *model.TpTableRelation) (*model.TpTableRelation, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpTableRelationDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpTableRelation{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpTableRelationDao) SelectByID(id int) (*model.TpTableRelation, error) {

	var m model.TpTableRelation
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpTableRelationDao) Update(id int, update map[string]interface{}) (*model.TpTableRelation, error) {

	var m model.TpTableRelation
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
