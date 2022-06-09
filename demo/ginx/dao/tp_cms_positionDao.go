
package dao
import (
	"ginx/model"
)

// TpCmsPositionDao ...
type TpCmsPositionDao struct {
}

// Create 增
func (*TpCmsPositionDao) Create(m *model.TpCmsPosition) (*model.TpCmsPosition, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpCmsPositionDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpCmsPosition{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpCmsPositionDao) SelectByID(id int) (*model.TpCmsPosition, error) {

	var m model.TpCmsPosition
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpCmsPositionDao) Update(id int, update map[string]interface{}) (*model.TpCmsPosition, error) {

	var m model.TpCmsPosition
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
