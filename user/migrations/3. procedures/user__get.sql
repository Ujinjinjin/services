create or replace function user__get(
	p_user_id integer
)
returns table (
	user_id integer,
	username varchar(128),
	email varchar(128),
	first_name varchar(256),
	last_name varchar(256),
	middle_name varchar(256)
) as
$$
begin
	----------------------------------------
	return query select
		"user".user_id,
		"user".username,
		"user".email,
		"user".first_name,
		"user".last_name,
		"user".middle_name
	from "user"
	where 1 = 1
		and "user".user_id = p_user_id
	;
	----------------------------------------
end;
$$
language plpgsql;
