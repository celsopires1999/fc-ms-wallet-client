START TRANSACTION;
Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date);
Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date);
Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);

insert into clients (id, name, email, created_at) values ('235769c3-963f-4b55-b2c8-dcba08810b16', 'John Doe', 'john@j.com', NOW());
insert into clients (id, name, email, created_at) values ('f2be64d6-18dc-44bd-879e-d7b30d8cb6bd', 'Jane Doe', 'jane@j.com', NOW());

insert into accounts (id, client_id, balance, created_at) values ('4926e6cf-9d18-49d6-ae4e-32c9ddcb9b81', '235769c3-963f-4b55-b2c8-dcba08810b16', 999, NOW());
insert into accounts (id, client_id, balance, created_at) values ('bf3a2451-f5cd-463a-845f-10cb9ee46d4f', 'f2be64d6-18dc-44bd-879e-d7b30d8cb6bd', 1, NOW());

COMMIT;