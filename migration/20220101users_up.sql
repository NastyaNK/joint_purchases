CREATE TABLE IF NOT EXISTS users (
	id       SERIAL NOT NULL,
	"name"   varchar(255) NOT NULL,
	email    varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	role     varchar(25) default 'buyer' not null,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	UNIQUE ("name", email)
);