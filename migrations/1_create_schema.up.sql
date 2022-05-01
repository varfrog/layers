create table items
(
    id bigint unsigned auto_increment,
    k varchar(255) not null,
    v mediumtext not null,
    constraint items_pk primary key (id),
    constraint items_key unique (k)
);
