create table time_variants (
  id bigint not null auto_increment,

  type_dateTime datetime,
  type_timestamp timestamp,

  primary key (id)
);
