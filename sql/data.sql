INSERT INTO products (name, price, qty)
VALUES ('Pizza', 200, 10),
       ('Burger', 150, 20),
       ('Free', 120, 15),
       ('Tea', 100, 50),
       ('Cola', 100, 50),
       ('Coffee', 100, 50);
       


INSERT INTO managers (name, salary, plan, boss_id, department)
VALUES ('Vasya', 100, 0, NULL, NULL),
       ('Petya', 80, 80, 1, 'boys'),
       ('Vanya', 60, 60, 2, 'boys'),
       ('Dasha', 90, 90, 1, 'girls'),
       ('Sasha', 70, 70, 4, 'girls'),
       ('Masha', 50, 50, 5, 'girls');



INSERT INTO customers (name, phone)
VALUES ('Zheyna', '+998941001010');

INSERT INTO sales (manager_id, customer_id)
VALUES (1, DEFAULT),
       (2, DEFAULT),
       (3, DEFAULT),
       (4, 1),
       (4, 1),
       (5, DEFAULT),
       (5, DEFAULT);

INSERT INTO sale_positions (sale_id, product_id, name, qty, price)
VALUES
-- Vasya, Pizza, 5 шт по 200
(1, 1, 'Pizza', 5, 200),
-- Vasya, Burger, 5 шт по 200
(1, 2, 'Burger', 5, 200),

--Petya, Free, 10 шт по 120
(2, 3, 'Free', 10, 120),

--Vanya, Free, 10 шт по 120
(3, 3, 'Free', 10, 120),

--Dasha, Coffee, 20 шт по 150
(4, 6, 'Coffee', 20, 150),

--Dasha, Coffee, 20 шт по 150
(5, 6, 'Coffee', 20, 150),

--Masha, Coffee, 20 шт по 150
(6, 6, 'Coffee', 20, 150),

--Masha, Coffee, 10 шт по 100
(7, 5, 'Cola', 10, 100);

/*  
*/
ALTER SEQUENCE customers_id_seq RESTART WITH 1;
ALTER SEQUENCE sale_positions_id_seq RESTART WITH 1;
ALTER SEQUENCE managers_id_seq RESTART WITH 1;
DELETE FROM products;
DELETE FROM managers;
DELETE FROM customers;
DELETE FROM sales;
DELETE FROM sale_positions;


INSERT INTO managers (id, name, login, password, salary, plan)
VALUES (1, 'abror', 'mike', '123', 1000, 100);




--                 1й вариант
INSERT INTO products VALUES (DEFAULT, 'iPhone', 100, 10, DEFAULT, DEFAULT);

--                 2й - вариант
INSERT INTO products (name, price, qty) VALUES ('iMac', 300, DEFAULT);

--                 3й - вариант
INSERT INTO products (name, price, qty)
VALUES ('MacBook', 200, 10),
       ('Mi Notebook', 120, 20),
       ('Mi Notebook Pro', 150, 15);

--                4й - вариант - возвращение по id и времени создания
INSERT INTO products (name, price, qty) VALUES ('Redmi Note 10', 40, 10) RETURNING id, created;



--                1й - вариант - конфликты(разные имена, но одинаковые телефоны)
-- 1. DO NOTHING - ничего не будет(просто ошибка) 
INSERT INTO customers (name, phone) 
VALUES ('Petya', '+998901001010')
ON CONFLICT DO NOTHING;

-- 2. Обновить запись с которой конфликтуем(DO UPDATE)
INSERT INTO customers (name, phone) 
VALUES ('Petya', '+998901001010')
ON CONFLICT (phone) DO UPDATE SET name = excluded.name;



--мы в продажу с id=1 добавили товар с id=1 в количестве 1 (по умолчанию), 
--выбрав название и цену из таблицы продуктов.
INSERT INTO sale_positions(sale_id, product_id, name, price)
SELECT 1, 1, name, price FROM products WHERE id = 1;

UPDATE sale_positions
SET (name, price) = (SELECT name, price FROM products WHERE id = 1)
WHERE id = 1;