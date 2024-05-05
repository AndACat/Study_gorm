package main

import (
	"Study_gorm/config/dao"
	"Study_gorm/config/db"
	"Study_gorm/model/entity"
	"encoding/json"
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

	if false {
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
	if true {
		// 2. 测试查询学生代码
		student2, _ := ClassDao.GetStudent()
		printObjToJson(2, student2)

		// 2.1 新增一条数据
		if false {
			ClassDao.CreateStudent(&entity.StudentEntity{Sno: "456", Sname: "王五", Age: 23, CollegeNo: "102", ClassNo: "2402"})
		}

		if false {
			// 2.2 批量新增数据
			for i := 0; i < 100; i++ {
				db.DB.Table("student").Create(&entity.StudentEntity{
					Sno:       strconv.Itoa(i),
					Sname:     strconv.Itoa(i),
					Age:       uint8(i),
					CollegeNo: strconv.Itoa(i),
					ClassNo:   strconv.Itoa(i),
				})
			}
		}

		// 从下面开始 直接操作gorm.DB, 不再封装函数
		// 3. First: 使用First方法查询记录，First会默认按照主键升序排列，取第一条记录
		student3 := &entity.StudentEntity{}
		// Table()方法是用于起别名
		db.DB.First(student3)
		printObjToJson(3, student3)

		// 4. Take: 获取一条记录，没有指定排序字段
		student4 := &entity.StudentEntity{}
		db.DB.Table("student").Take(student4)
		printObjToJson(4, student4)

		// 5. Last: 获取最后一条记录，按照主键降序排列
		student5 := &entity.StudentEntity{}
		db.DB.Table("student").Last(student5)
		printObjToJson(5, student5)
	}
	if true {
		// 3.
	}

}

// 用于将obj转为json并打印，入参为多个any对象，第一个是obj, 第二个是err
func printObjToJson(row int64, obj ...any) {
	bytes, _ := json.Marshal(obj[0])
	fmt.Println(strconv.FormatInt(row, 10) + ": " + string(bytes))
	fmt.Println("--------------------------")
}
