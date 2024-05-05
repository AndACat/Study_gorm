# username: golang
# password: 4_wvcLDCfhw.U4P12345

create database golang;
use golang;

create table student
(
    id         bigint PRIMARY KEY AUTO_INCREMENT comment '主键',
    sno        int(11)     not null comment '学号',
    sname      varchar(50) not null comment '姓名',
    age        tinyint(4)  null comment '年龄',
    class_no   varchar(20) not null comment '班级号',
    college_no varchar(20) not null comment '学院号',
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    UNIQUE INDEX sno_uniq (`sno`)
) ENGINE = INNODB;

create table class
(
    id         bigint PRIMARY KEY AUTO_INCREMENT comment '主键',
    class_no   int(11)     not null comment '班级号',
    class_name varchar(50) not null comment '班级名',
    college_no varchar(50) not null comment '学院号',
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    unique index cno_uniq (`class_no`)
) engine = INNODB;

create table college
(
    id           bigint PRIMARY KEY AUTO_INCREMENT comment '主键',
    college_no   int(11)     not null comment '学院号',
    college_name varchar(50) not null comment '学院名',
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    unique index college_no_uniq (`college_no`)
) engine = INNODB;

