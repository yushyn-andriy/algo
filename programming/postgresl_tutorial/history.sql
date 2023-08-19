-- SELECT

select first_name, last_name, email from customer;
select first_name || ' ' || last_name || ' ' || email from customer;
select 5*9;
select current_date;
select current_timestamp;
select first_name || ' ' || last_name || ' ' || email as full_info from customer; 

-- ORDER BY

select first_name as fn, last_name as ls, length(first_name) as len, first_name ||'<>'|| last_name as full_name  from customer
order by len, fn asc,  ls asc;

create table sort_demo (
	num int
);

insert into sort_demo(num)
values
(1),
(2),
(3),
(null);

select num from sort_demo order by num asc nulls first;
select num from sort_demo order by num asc nulls last;

-- DISTINCT

create table if not exists demo_distinct (
	id serial not null primary key,
	bcolor varchar,
	fcolor varchar
);

insert into demo_distinct (bcolor, fcolor)
values
('red', 'green'),
('green', 'red'),
('red', 'red'),
('red', null),
(null, 'red'),
(null, null);


select distinct on(bcolor) bcolor, fcolor from demo_distinct;

-- WHERE

select * from customer
where first_name = 'Jimmy';
select * from customer
where first_name = 'Jamie' and last_name = 'Rice';

select * from customer 
where first_name = 'Jamie' or last_name = 'Rice';

select * from customer 
where first_name IN ('Ann','Anne','Annie');


select * from customer
where first_name LIKE '%ann%';


select 
first_name,
LENGTH(first_name) flen
from customer 
where
first_name <> 'Amy' and
first_name != 'Sue' and
LENGTH(first_name) BETWEEN 3 and 6
order by flen;

-- LIMIT, OFFSET

select film_id, title, rental_rate, rental_rate from
film
order by rental_rate desc
limit 10 offset 5;


select * from film
order by title
fetch first 3 row only;

select * from information_schema.tables
where 
	table_catalog = 'dvdrental' and
	table_schema != 'pg_catalog' and
	table_type = 'BASE TABLE' and
	table_schema='public';

select * from film limit 10 offset 10;
select count(film_id) from film; 

select * from film
order by title
offset 2 row
fetch first 3 row only;


-- IN, SUBQUERY

select * from customer
where customer_id in (
	select customer_id from customer 
	where first_name in ('Mary', 'Linda')
) and not customer_id not in (1);
SELECT customer_id, return_date,
CAST (return_date AS DATE) as dd
FROM rental
WHERE CAST (return_date AS DATE) = '2005-05-27';

select * from customer
where customer_id in (
SELECT customer_id
FROM rental
WHERE CAST (return_date AS DATE) = '2005-05-27');

-- LIKE, ILIKE

SELECT
	customer_id,
	payment_id,
	amount,
 payment_date
FROM
	payment
WHERE
	payment_date BETWEEN '2007-02-07' AND '2007-02-15';


select * from customer limit 100;


select * from customer where
first_name not ILIKE 'm___';


-- IS NULL, IS NOT NULL

create table if not exists contacts (
	id serial not null,
	name varchar(50) not null,
	phone varchar(12) default null,
	primary key (id)
);

insert into contacts (name, phone)
values
('a', null),
('a', '709823'),
('b', '12e21e'),
('a', null);

select * from contacts where
phone is null;
select * from contacts where
phone is not null;


-- JOINs

create table if not exists basket_a (
	id serial primary key,
	fruit_a varchar
);

create table if not exists basket_b (
	id serial primary key,
	fruit_b varchar
);

INSERT INTO basket_a (id, fruit_a)
VALUES
    (1, 'Apple'),
    (2, 'Orange'),
    (3, 'Banana'),
    (4, 'Cucumber');

INSERT INTO basket_b (id, fruit_b)
VALUES
    (1, 'Orange'),
    (2, 'Apple'),
    (3, 'Watermelon'),
    (4, 'Pear');

select * from basket_a inner join
basket_b on basket_a.fruit_a = basket_b.fruit_b
order by basket_a.id;


select * from basket_a left outer join 
basket_b on basket_a.fruit_a = basket_b.fruit_b
order by basket_a.id;

select * from basket_a right outer join 
basket_b on basket_a.fruit_a = basket_b.fruit_b
order by basket_a.id;

select * from basket_a full outer join 
basket_b on basket_a.fruit_a = basket_b.fruit_b
order by basket_a.id;

-- AS

select * from payment;


select * from customer as c
inner join payment as p on c.customer_id = p.customer_id
where c.customer_id = 342;
