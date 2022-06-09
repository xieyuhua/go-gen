
package dao
import (
	"ginx/model"
)

// TpCmsBannerDao ...
type TpCmsBannerDao struct {
}

// Create 增
func (*TpCmsBannerDao) Create(m *model.TpCmsBanner) (*model.TpCmsBanner, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*TpCmsBannerDao) Delete(id int) error {
	err := model.DB.Delete(&model.TpCmsBanner{ID: id}).Error
	return err
}

// SelectByID 查
func (*TpCmsBannerDao) SelectByID(id int) (*model.TpCmsBanner, error) {

	var m model.TpCmsBanner
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*TpCmsBannerDao) Update(id int, update map[string]interface{}) (*model.TpCmsBanner, error) {

	var m model.TpCmsBanner
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
