do $$
begin
	if (select system__exists_table('user_history') is false) then
		create table "user_history"
		(
			user_history_id serial not null,
			user_id int not null,
			username varchar(128) not null,
			email varchar(128) not null,
			first_name varchar(256) null,
			last_name varchar(256) null,
			middle_name varchar(256) null,
			date_created timestamp not null,
			is_deleted bool not null,
			change_user_id int not null,
			date_changed timestamp not null
		);

		alter table "user_history" add constraint pk$user_history primary key (user_history_id);

		create index i$user_history$user_id on "user_history" (user_id);
		create index i$user_history$change_user_id on "user_history" (change_user_id);
	end if;
end $$
