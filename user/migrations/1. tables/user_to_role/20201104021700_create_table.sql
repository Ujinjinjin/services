create table user_to_role
(
	user_to_role_id serial not null,
	user_id int not null,
	role_id int not null
);

alter table user_to_role add constraint pk$user_to_role primary key (user_to_role_id);
alter table user_to_role add constraint fk$user_to_role$user_id foreign key (user_id) references "user" (user_id);
alter table user_to_role add constraint fk$user_to_role$role_id foreign key (role_id) references "role" (role_id);

create index i$user_to_role$user_id$role_id on user_to_role (user_id, role_id);
