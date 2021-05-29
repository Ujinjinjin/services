do $$
begin
	if (select system__exists_column('user', 'date_created') is false) then
		alter table "user" add column "date_created" timestamp not null default current_timestamp;
	end if;
end $$
