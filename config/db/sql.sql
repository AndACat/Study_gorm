# username: golang
# password: 4_wvcLDCfhw.U4P12345

create database golang;
use golang;

create table student
(
    id         bigint PRIMARY KEY AUTO_INCREMENT comment '����',
    sno        int(11)     not null comment 'ѧ��',
    sname      varchar(50) not null comment '����',
    age        tinyint(4)  null comment '����',
    class_no   varchar(20) not null comment '�༶��',
    college_no varchar(20) not null comment 'ѧԺ��',
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    UNIQUE INDEX sno_uniq (`sno`)
) ENGINE = INNODB;

create table class
(
    id         bigint PRIMARY KEY AUTO_INCREMENT comment '����',
    class_no   int(11)     not null comment '�༶��',
    class_name varchar(50) not null comment '�༶��',
    college_no varchar(50) not null comment 'ѧԺ��',
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    unique index cno_uniq (`class_no`)
) engine = INNODB;

create table college
(
    id           bigint PRIMARY KEY AUTO_INCREMENT comment '����',
    college_no   int(11)     not null comment 'ѧԺ��',
    college_name varchar(50) not null comment 'ѧԺ��',
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null,
    unique index college_no_uniq (`college_no`)
) engine = INNODB;
