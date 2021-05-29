do $$
begin
	if (select system__exists_table('role_to_privilege') is false) then
		create table role_to_privilege
		(
			role_to_privilege_id serial not null,
			role_id int not null,
			privilege_id int not null
		);

		alter table role_to_privilege add constraint pk$role_to_privilege primary key (role_to_privilege_id);
		alter table role_to_privilege add constraint fk$role_to_privilege$role_id foreign key (role_id) references "role" (role_id);
		alter table role_to_privilege add constraint fk$role_to_privilege$privilege_id foreign key (privilege_id) references privilege (privilege_id);

		create index i$role_to_privilege$role_id$privilege_id on role_to_privilege (role_id, privilege_id);
	end if;
end $$
