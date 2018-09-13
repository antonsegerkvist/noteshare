create database if not exists `noteshare`;
use `noteshare`;

--
-- Table structure for table `t_account`
--

drop table if exists `t_account`;
create table `t_account` (
  `c_id`   bigint  (10) unsigned not null auto_increment,
  `c_name` varchar (256) not null,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_file`
--

drop table if exists `t_file`;
create table `t_file` (
  `c_id`                   bigint (10) unsigned not null auto_increment,
  `c_type`                 bigint (10) unsigned not null,
  `c_file_reference_id`    bigint (10) unsigned default null,
  `c_file_reference_count` bigint (10) unsigned not null default '1',
  `c_parent`               bigint (10) unsigned default null,
  `c_name`                 varchar (256) not null,
  `c_is_processed`         tinyint (3) unsigned not null default '0',
  `c_is_uploaded`          tinyint (3) unsigned not null default '0',
  `c_modified_by_user_id`  bigint (10) unsigned not null,
  `c_modification_date`    datetime not null default CURRENT_TIMESTAMP,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_file_belongs_to_user`
--

drop table if exists `t_file_belongs_to_user`;
create table `t_file_belongs_to_user` (
  `c_user_id` bigint(10) unsigned not null,
  `c_file_id` bigint(10) unsigned not null,
  primary key (`c_user_id`,`c_file_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_group`
--

drop table if exists `t_group`;
create table `t_group` (
  `c_id`   bigint(10) unsigned not null auto_increment,
  `c_name` varchar(256) not null,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_user`
--

drop table if exists `t_user`;
create table `t_user` (
  `c_id`            bigint(10) unsigned not null auto_increment,
  `c_account_id`    bigint(10) unsigned not null,
  `c_email`         varchar(320) not null,
  `c_username`      varchar(256) not null,
  `c_password_hash` varchar(64) not null,
  `c_password_salt` varchar(64) not null,
  `c_created`       datetime not null default CURRENT_TIMESTAMP,
  `c_activated`     datetime default null,
  primary key (`c_id`),
  unique key (`c_email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_user_belongs_to_group`
--

drop table if exists `t_user_belongs_to_group`;
create table `t_user_belongs_to_group` (
  `c_id`       bigint(10) unsigned not null auto_increment,
  `c_user_id`  bigint(10) unsigned not null,
  `c_group_id` bigint(10) unsigned not null,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
