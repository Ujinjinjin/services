do $$
begin
	if (select system__exists_column('user', 'is_deleted') is false) then
		alter table "user" add column "is_deleted" bool not null default false;
		create index i$user$is_deleted on "user" (is_deleted);
	end if;
end $$
