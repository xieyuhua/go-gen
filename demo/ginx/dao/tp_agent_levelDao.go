
package dao
import (
	"ginx/model"
)

// TpAgentLevelDao ...
type TpAgentLevelDao struct {
}

// Create 增
func (*TpAgentLevelDao) Create(m *model.TpAgentLevel) (*model.TpAgentLevel, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpAgentLevelDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpAgentLevel{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpAgentLevelDao) SelectByID(id int) (*model.TpAgentLevel, error) {

	var m model.TpAgentLevel
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpAgentLevelDao) Update(id int, update map[string]interface{}) (*model.TpAgentLevel, error) {

	var m model.TpAgentLevel
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
