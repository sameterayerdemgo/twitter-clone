CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT uuid_generate_v4();


create table post
(
    id        uuid DEFAULT uuid_generate_v4 () ,
    message   text not null,
    timestamp timestamp not null,
    CONSTRAINT post_pkey PRIMARY KEY (id)

)