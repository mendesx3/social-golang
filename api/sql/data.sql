#CREATE DATABASE IF NOT EXISTS db
#USER db

DROP TABLE IF EXISTS db.user;

CREATE TABLE db.user (
    id int auto_increment primary key,
    name varchar(50) not null,
    password varchar(20) not null
) ENGINE=INNODB;

insert into user (name, password) values ('andre','golag');

select * from db.user

select * from dual