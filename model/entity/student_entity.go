package entity

import "gorm.io/gorm"

/*
	create table student(
		id bigint PRIMARY KEY AUTO_INCREMENT comment '主键',
		sno int(11) not null comment '学号',
		sname varchar(50) not null comment '姓名',
		age tinyint(4) null comment '年龄',
		class_no varchar(20) not null comment '班级号',
		college_no varchar(20) not null comment '学院号',
		UNIQUE INDEX sno_uniq(`sno`)
	)ENGINE=INNODB;
*/
// 视图层 定义和数据库映射的实体类
type StudentEntity struct {
	gorm.Model
	// 这个可以覆盖gorm.Model里面ID属性, 用于指定实体对应表的别名
	ID        uint `gorm:"primaryKey; tableName: student"`
	Sno       string
	Sname     string
	Age       uint8
	ClassNo   string
	CollegeNo string
}
