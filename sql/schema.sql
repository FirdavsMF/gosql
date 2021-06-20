-- таблица товаров
CREATE TABLE products 
(
    id      BIGSERIAL PRIMARY KEY,
    name    TEXT      NOT NULL,
    price   INTEGER   NOT NULL CHECK (price > 0),
    qty     INTEGER   NOT NULL DEFAULT 0 CHECK (qty >= 0),
    active  BOOLEAN   NOT NULL DEFAULT TRUE, 
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-----------------------------------------------------
-- SELECT id, name, price FROM products 
-- WHERE active AND qty > 0     // активный товар и который есть на складе
-- ORDER BY price DESC          //( DESC - сортировка по убыванию, т.е идёт сначало самый большой)
-- LIMIT 3;                     // Количество(лимит) 3 товара

--таблица сотрудников
CREATE TABLE managers
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT      NOT NULL,
    login      TEXT      NOT NULL UNIQUE,
    password   TEXT      NOT NULL,
    salary     INTEGER   NOT NULL CHECK (salary > 0),
    plan       INTEGER   NOT NULL CHECK (plan > 0),
    boss_id    BIGINT    REFERENCES managers,
    department TEXT,
    active     BOOLEAN   NOT NULL DEFAULT TRUE, 
    created    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--таблица зарегестрированных покупателей
CREATE TABLE customers
(
    id       BIGSERIAL PRIMARY KEY,
    name     TEXT      NOT NULL,
    phone    TEXT      NOT NULL UNIQUE,
    password TEXT      NOT NULL,
    active   BOOLEAN   NOT NULL DEFAULT TRUE, 
    created  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--таблица токенов
CREATE TABLE customers_tokens 
(
    token        TEXT      NOT NULL,
    customer_id  BIGINT    NOT NULL REFERENCES customers,
    expire       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 hour',
    created      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);


--таблица продаж
CREATE TABLE sales
(
    id          BIGSERIAL PRIMARY KEY,
    manager_id  BIGINT    NOT NULL REFERENCES managers,
    customer_id BIGINT    REFERENCES customers,
    created     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--конкретные позиции в продаже (чек)
CREATE TABLE sale_positions
(
    id          BIGSERIAL PRIMARY KEY,
    sale_id     BIGINT    NOT NULL REFERENCES sales,
    product_id  BIGINT    NOT NULL REFERENCES products,
    name        TEXT      NOT NULL,
    price       INTEGER   NOT NULL CHECK (price >= 0),
    qty         INTEGER   NOT NULL DEFAULT 1 CHECK (qty > 0),
    created     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--DROP TABLE <name_of_table> CASCADE

--DROP TABLE products;
--DROP TABLE managers;
--DROP TABLE customers;
--DROP TABLE sales;
--DROP TABLE sale_positions;
