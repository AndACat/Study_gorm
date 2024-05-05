package main

import (
	"Study_gorm/config/dao"
	"Study_gorm/config/db"
	"Study_gorm/model/entity"
	"fmt"
	"strconv"
)

var (
	// 定义mapper对象
	ClassDao = dao.ClassDao{}
)

// main： 新增学生信息
func main() {
	// 先让gorm连接数据库，将连接好的信息放置到db.DB中
	// 下面这种写法, 就不会产生err命名的冲突, 而且go中很多人都是这样写代码的
	if _, err := db.InitMySql(); err != nil {
		panic("mysql连接错误")
	}
	// 1. 测试新增学生代码
	studentEntity := &entity.StudentEntity{Sno: "123", Sname: "zhangsan", Age: 18, CollegeNo: "101", ClassNo: "2401"}
	rowsAffected, err := ClassDao.CreateStudent(studentEntity)
	if err != nil {
		return
	}
	// 使用strconv库，将一个int的数字转化为10进制的字符串
	// 在go中无法直接将string和int类型的数据进行拼接成字符串
	fmt.Println("受影响的行数：" + strconv.FormatInt(rowsAffected, 10))
}
