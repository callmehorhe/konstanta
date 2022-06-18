CREATE TABLE IF NOT EXISTS transactions(
    t_id serial NOT NULL,
    user_id integer NOT NULL,
    email varchar(255) NOT NULL,
    amount double precision NOT NULL,
    currency varchar(255) NOT NULL,
    create_time timestamp NOT NULL,
    update_time timestamp NOT NULL,
    status varchar(255) NOT NULL,
    PRIMARY KEY (t_id)
);