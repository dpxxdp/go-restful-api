# go-restful-api
serves json over http
learning go



useful

export GOPATH=$HOME/Developer/goworkspace
export PATH=$PATH:$GOPATH/bin



postgres

pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start
pg_ctl -D /usr/local/var/postgres stop -s -m fast


CREATE TABLE users
(	
	user_id integer CONSTRAINT user_id_pk PRIMARY KEY,
	name varchar[20],
	email varchar[20],
	created timestamp,
	hash varchar[50],
	salt varchar[50],
	active bool
);

CREATE TABLE channels
(	
	channel_id integer CONSTRAINT channel_id_pk PRIMARY KEY,
	name varchar[20],
	created timestamp,
	active bool
);

CREATE TABLE mod_channel_association
(
	user_id integer REFERENCES users (user_id),
	channel_id integer REFERENCES channels (channel_id)
);

CREATE TABLE user_channel_association
(
	user_id integer REFERENCES users (user_id),
	channel_id integer REFERENCES channels (channel_id)
);

CREATE TABLE medias
(	
	media_id integer CONSTRAINT media_id_pk PRIMARY KEY,
	title varchar[20],
	content text,
	user_id integer REFERENCES users (user_id),
	channel_id integer REFERENCES channels (channel_id),
	created timestamp,
	points integer,
	active bool
);