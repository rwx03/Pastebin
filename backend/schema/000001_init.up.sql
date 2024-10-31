create table users(
	id serial primary key,
	email varchar(100),
	password varchar(100),
	created_at timestamp,
	updated_at timestamp
);

create table tokens(
	id serial primary key,
	user_id int references users(id) on delete cascade,
	refresh_token varchar(255)
);

create table pastes(
	id serial primary key,
	paste_id varchar(20) not null unique,
	title varchar(100),
	content varchar(5000),
	creator_id int,
	created_at timestamp,
	views int
);