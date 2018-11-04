create database if not exists `noteshare`;
use `noteshare`;

--
-- Table structure for table `t_account`
--

drop table if exists `t_account`;
create table `t_account` (
  `c_id`     bigint  (10)  unsigned not null auto_increment,
  `c_name`   varchar (256) not null,
  `c_layout` json          default null,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_user`
--

drop table if exists `t_user`;
create table `t_user` (
  `c_id`            bigint  (10)  unsigned not null auto_increment,
  `c_account_id`    bigint  (10)  unsigned not null,
  `c_is_admin`      tinyint (8)   not null default 0,
  `c_email`         varchar (320) not null,
  `c_username`      varchar (256) not null,
  `c_password_hash` varchar (64)  not null,
  `c_password_salt` varchar (64)  not null,
  `c_created`       datetime not null default CURRENT_TIMESTAMP,
  `c_activated`     datetime default null,
  primary key (`c_id`),
  unique key (`c_email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_group`
--

drop table if exists `t_group`;
create table `t_group` (
  `c_id`         bigint  (10)  unsigned not null auto_increment,
  `c_account_id` bigint  (10)  unsigned not null,
  `c_name`       varchar (256) not null,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_user_belongs_to_group`
--

drop table if exists `t_user_belongs_to_group`;
create table `t_user_belongs_to_group` (
  `c_id`       bigint (10) unsigned not null auto_increment,
  `c_user_id`  bigint (10) unsigned not null,
  `c_group_id` bigint (10) unsigned not null,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_group_permission`
--

drop table if exists `t_group_permission`;
create table `t_group_permission` (
  `c_account_id` bigint (10) unsigned not null,
  `c_group_id`   bigint (10) unsigned not null,
  `c_key`        int    (10) unsigned not null,
  `c_value`      int    (10) unsigned not null,
  primary key (`c_group_id`, `c_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_folder_group_permission`
--

drop table if exists `t_folder_group_permission`;
create table `t_folder_group_permission` (
  `c_account_id` bigint (10) unsigned not null,
  `c_group_id`   bigint (10) unsigned not null,
  `c_folder_id`  bigint (10) unsigned not null,
  `c_key`        int    (10) unsigned not null,
  `c_value`      int    (10) unsigned not null,
  primary key (`c_group_id`, `c_folder_id`, `c_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_folder`
--

drop table if exists `t_folder`;
create table `t_folder` (
  `c_id`                  bigint  (10)  unsigned not null auto_increment,
  `c_account_id`          bigint  (10)  unsigned not null,
  `c_parent`              bigint  (10)  unsigned default 0,
  `c_name`                varchar (256) not null,
  `c_created_by_user_id`  bigint  (10)  unsigned not null,
  `c_modified_by_user_id` bigint  (10)  unsigned not null,
  `c_created_date`        datetime not null default CURRENT_TIMESTAMP,
  `c_modified_date`       datetime not null default CURRENT_TIMESTAMP,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `t_file`
--

drop table if exists `t_file`;
create table `t_file` (
  `c_id`                  bigint  (10)  unsigned not null auto_increment,
  `c_account_id`          bigint  (10)  unsigned not null,
  `c_is_uploaded`         tinyint (8)   unsigned not null default 0,
  `c_is_processed`        tinyint (8)   unsigned not null default 0,
  `c_has_preview`         tinyint (8)   unsigned not null default 0,
  `c_type`                bigint  (10)  unsigned not null,
  `c_name`                varchar (256) not null,
  `c_filename`            varchar (256) not null,
  `c_filesize`            bigint  (10)  unsigned not null,
  `c_checksum`            bigint  (10)  unsigned not null,
  `c_created_by_user_id`  bigint  (10)  unsigned not null,
  `c_modified_by_user_id` bigint  (10)  unsigned not null,
  `c_created_date`        datetime not null default CURRENT_TIMESTAMP,
  `c_modified_date`       datetime not null default CURRENT_TIMESTAMP,
  primary key (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;