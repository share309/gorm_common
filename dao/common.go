package dao

import "gorm_common/global"

func Create[T any](data *T) bool {

	if err := global.DB.Create(data).Error; err != nil {
		return false
	}

	return true
}

func Delete[T any](data *T) bool {
	if err := global.DB.Delete(data).Error; err != nil {
		return false
	}
	return true
}

// save 结构体全量更新
// updates 有值的更新，不更新

func Update[T any](data *T) bool {
	if err := global.DB.Updates(data).Error; err != nil {
		return false
	}
	return true
}

func GetOneById[T any](id uint, data *T) bool {

	if err := global.DB.Where("id=?", id).Find(data).Limit(1).Error; err != nil {
		return false
	}
	return true
}

func GetOneByFields[T any](where *T, data *T) bool {

	if err := global.DB.Where(where).Find(data).Limit(1).Error; err != nil {
		return false
	}
	return true

}

func GetListByFields[T any](where *T, data *[]T) bool {

	if err := global.DB.Where(where).Find(data).Error; err != nil {
		return false
	}
	return true

}
