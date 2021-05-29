do $$
begin
	if (select system__exists_table('role') is false) then
		create table "role"
		(
			role_id serial not null,
			name varchar(128) not null
		);

		alter table "role" add constraint pk$role primary key (role_id);
	end if;
end $$
