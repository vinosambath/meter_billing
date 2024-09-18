create user 'vinxsam'@'localhost' identified by 'Password@1';

create database readings_01;
create database readings_02;

--> in both readings_01 db and readings_02 db
create table meter_readings (
    id varchar(36) not null,
    `nmi` varchar(10) not null,
    `timestamp` timestamp not null,
    `consumption` numeric not null,
    constraint meter_readings_pk primary key (id),
    constraint meter_readings_unique_consumption unique (`nmi`, `timestamp`)
);

