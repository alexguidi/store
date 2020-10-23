create table products(id serial primary key,name varchar,description varchar,price decimal,quantity integer)

insert into products(name, description, price, quantity) values('T-Shirt', 'Black', 19, 10),('Earpod', 'Very Cool', 99, 5);