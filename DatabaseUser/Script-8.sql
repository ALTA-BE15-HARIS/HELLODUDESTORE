use project1;

CREATE TABLE user (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  name varchar(100),
  password longtext,
  no_hp varchar(15),
  saldo bigint,
  PRIMARY KEY (`id`),
  UNIQUE KEY no_hp (`no_hp`)
) ENGINE=InnoDB AUTO_INCREMENT=1;

select * from user;