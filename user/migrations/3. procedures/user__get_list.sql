-- Drop function
select system__drop_routine('user__get_list');

-- Crate function
create or replace function user__get_list(
	p_user_id_list integer[],
	p_username varchar(128),
	p_email varchar(128),
	p_is_deleted integer
)
returns table (
	user_id integer,
	username varchar(128),
	email varchar(128),
	first_name varchar(256),
	last_name varchar(256),
	middle_name varchar(256),
	is_deleted bool
) as
$$
declare v_is_deleted bool;
begin
	----------------------------------------
	case
		when p_is_deleted = 1 then v_is_deleted = true;
		when p_is_deleted = 2 then v_is_deleted = false;
		else v_is_deleted = null;
	end case;
	----------------------------------------
	return query select
		"user".user_id,
		"user".username,
		"user".email,
		"user".first_name,
		"user".last_name,
		"user".middle_name,
		"user".is_deleted
	from "user"
	where 1 = 1
		and(1 = 0
			or array_length(p_user_id_list, 1) = 0
			or "user".user_id = any(p_user_id_list)
		)
		and(1 = 0
			or p_username is null
			or "user".username like '%' || p_username || '%'
		)
		and(1 = 0
			or p_email is null
			or "user".email like '%' || p_email || '%'
		)
		and(1 = 0
			or v_is_deleted is null
			or "user".is_deleted = v_is_deleted
		)
	order by "user".user_id
	;
	----------------------------------------
end;
$$
language plpgsql;
