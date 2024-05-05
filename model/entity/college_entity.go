package entity

import "gorm.io/gorm"

/*
create table college(

	id bigint PRIMARY KEY AUTO_INCREMENT comment '主键',
	college_no int(11) not null comment '学院号',
	college_name varchar(50) not null comment '学院名',
	unique index college_no_uniq (`college_no`)

)engine = INNODB;
*/
type CollegeEntity struct {
	gorm.Model
	ID          uint
	CollegeNo   string
	CollegeName string
}
