-- Drop function
select system__drop_routine('user__add');

-- Crate function
create or replace function user__add(
	p_username varchar(128),
	p_email varchar(128),
	p_first_name varchar(256),
	p_last_name varchar(256),
	p_middle_name varchar(256),
	p_change_user_id integer
)
returns integer as
$$
declare r_user_id integer;
begin
	----------------------------------------
	if not exists(
		select 1
		from "user"
		where 1 = 0
			or "user".email = p_email
			or "user".username = p_username
	) then
		insert into "user" (
			username,
			email,
			first_name,
			last_name,
			middle_name
		) values (
			p_username,
			p_email,
			p_first_name,
			p_last_name,
			p_middle_name
		)
		returning user_id into r_user_id;
	else
		r_user_id := -1;
	end if;
	----------------------------------------
	perform user__fill_history(
		p_user_id_list := array(select r_user_id),
		p_change_user_id := p_change_user_id
	);
	----------------------------------------
	return r_user_id;
	----------------------------------------
end;
$$
language plpgsql;
