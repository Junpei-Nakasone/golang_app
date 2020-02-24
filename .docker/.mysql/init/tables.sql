DROP TABLE IF EXISTS `test_table`;

create table IF not exists `test_table`
(
    `id` INT(20) AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL
)