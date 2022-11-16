CREATE DATABASE IF NOT EXISTS devgo
USER devgo

DROP TABLE IF EXISTS user

CREATE TABLE USER(
    id int auto_increment primary key,
    name varchar(50) not null,
    password varchar(20) not null,
    active boolean DEFAULT true,
    created timestamp DEFAULT current_timestamp()
) ENGINE=INNODB;