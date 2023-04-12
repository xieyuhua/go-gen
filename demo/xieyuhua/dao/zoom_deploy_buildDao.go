
package dao
import (
	"xieyuhua/model"
)

// ZoomDeployBuildDao ...
type ZoomDeployBuildDao struct {
}

// Create 增
func (*ZoomDeployBuildDao) Create(m *model.ZoomDeployBuild) (*model.ZoomDeployBuild, error) {
	err := model.DB.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete 删
func (*ZoomDeployBuildDao) Delete(id int) error {
	err := model.DB.Delete(&model.ZoomDeployBuild{ID: id}).Error
	return err
}

// SelectByID 查
func (*ZoomDeployBuildDao) SelectByID(id int) (*model.ZoomDeployBuild, error) {

	var m model.ZoomDeployBuild
	err := model.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// All 查
func (*ZoomDeployBuildDao) All() (*[]model.ZoomDeployBuild, error) {
	var ms []model.ZoomDeployBuild
	err := model.DB.Find(&ms).Error
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*ZoomDeployBuildDao) Update(id int, update map[string]interface{}) (*model.ZoomDeployBuild, error) {

	var m model.ZoomDeployBuild
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
