CREATE DATABASE IF NOT EXISTS db;
USE db;


DROP TABLE IF EXISTS db.user;

create table db.user (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null,
    email varchar(50) not null,
    password varchar(20) not null,
    created timestamp default current_timestamp()
) ENGINE=INNODB;

insert into db.user
        (name, nick, email, password, created)
    values
        ('Fulano', 'fulano', '@gmail.com', '123456', now());

