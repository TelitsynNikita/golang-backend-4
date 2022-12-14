create table users
(
    id        serial unique,
    full_name varchar(255) not null,
    email     varchar(255) not null unique,
    password  varchar(255) not null,
    role      varchar(255) default 'USER'
);

create table deals
(
    id serial unique,
    purpose varchar(255) not null,
    description varchar(255),
    count integer not null,
    product_link varchar(255) not null,
    amount integer not null,
    status varchar(255) default 'NEW',
    user_id integer not null references users (id) on delete cascade,
    bookkeeper_id integer  default 0,
    created_at    timestamp    default now()
);