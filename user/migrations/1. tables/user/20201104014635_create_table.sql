create table "user"
(
	user_id serial not null,
	username varchar(128) not null,
	email varchar(128) not null,
	first_name varchar(256) null,
	last_name varchar(256) null,
	middle_name varchar(256) null
);

alter table "user" add constraint pk$user primary key (user_id);

create index i$user$username on "user" (username);
create index i$user$email on "user" (email);
