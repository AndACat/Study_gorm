package entity

import "gorm.io/gorm"

/*
create table class(

	id bigint PRIMARY KEY AUTO_INCREMENT comment '主键',
	class_no int(11) not null comment '班级号',
	class_name varchar(50) not null comment '班级名',
	college_no varchar(50) not null comment '学院号',
	unique index cno_uniq (`class_no`)

)engine = INNODB;
*/
type ClassEntity struct {
	gorm.Model
	ID        uint
	ClassNo   string
	ClassName string
	CollegeNo string
}
