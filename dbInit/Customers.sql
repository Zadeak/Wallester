DROP TABLE IF EXISTS customer;

CREATE TABLE customer
(
    id         serial NOT NULL PRIMARY KEY,
    first_name varchar(100),
    last_name  varchar(100),
    email varchar (100)
);

insert into customer values (1, 'VVasja', 'Pupkin', 'emailtest@gmail.com');
insert into customer values (2, 'Vladimir', 'Putin','emailtest2@Gmail.com');

SELECT setval('customer_id_seq', 100, true);