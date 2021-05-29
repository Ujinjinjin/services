-- Drop function
select system__drop_routine('user__fill_history');

-- Crate function
create or replace function user__fill_history(
	p_user_id_list integer[],
	p_change_user_id integer
)
returns void as
$$
begin
	----------------------------------------
	insert into "user_history" (
		user_id,
		username,
		email,
		first_name,
		last_name,
		middle_name,
		date_created,
		is_deleted,
		change_user_id,
		date_changed
	) select
		user_id,
		username,
		email,
		first_name,
		last_name,
		middle_name,
		date_created,
		is_deleted,
		p_change_user_id,
		current_timestamp
	from "user"
	where 1 = 1
		and "user".user_id = any(p_user_id_list)
	;
	----------------------------------------
end;
$$
language plpgsql;
