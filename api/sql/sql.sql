create database if not exists devbook;

use devbook;

create table if not exists users(
	id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    createdIn timestamp default current_timestamp()
) engine=InnoDB;