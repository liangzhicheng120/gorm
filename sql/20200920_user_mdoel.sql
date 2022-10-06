CREATE SCHEMA `learn_gorm_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

create table learn_gorm_db.user_tab
(
    id            bigint auto_increment
        primary key,
    name          varchar(45)         null,
    age           int unsigned        null,
    email         varchar(45)         null,
    is_del        tinyint(1) unsigned null,
    create_time   int unsigned        null,
    update_time   int unsigned        null,
    birthday      timestamp(6) default CURRENT_TIMESTAMP(6) on update CURRENT_TIMESTAMP(6),
    member_number varchar(100),
    delete_at     timestamp(6)        null

);