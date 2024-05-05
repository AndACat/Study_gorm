package dao

import (
	"Study_gorm/config/db"
	"Study_gorm/model/entity"
)

// dao层代表了能够访问数据库
type ClassDao struct{}

// 在方法前面绑定ClassDao是为了防止方法名冲突

// CreateStudent 创建学生信息
func (ClassDao) CreateStudent(student *entity.StudentEntity) (int64, error) {
	db := db.DB.Table("student").Create(student)
	return db.RowsAffected, db.Error
}

// GetStudent 查询学生信息
func (ClassDao) GetStudent() (*entity.StudentEntity, error) {
	student := &entity.StudentEntity{}
	// First方法默认按照主键升序排列
	err := db.DB.Table("student").First(student).Error
	return student, err
}
