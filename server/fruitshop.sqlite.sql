--SQLite Maestro 16.11.0.11
------------------------------------------
--Host     : localhost
--Database : C:\Users\Rockey\VSCodeProjects\fruitshop\fruitshop.sqlite


CREATE TABLE applied_dual_item_discounts (
  id                       integer PRIMARY KEY AUTOINCREMENT,
  created_at               datetime,
  updated_at               datetime,
  deleted_at               datetime,
  cart_id                  integer NOT NULL,
  dual_item_discount_id    integer,
  dual_item_discount_name  varchar(255),
  savings                  real
);

CREATE INDEX idx_applied_dual_item_discounts_deleted_at
  ON applied_dual_item_discounts
  (deleted_at);

CREATE TABLE applied_single_item_coupons (
  id                       integer PRIMARY KEY AUTOINCREMENT,
  created_at               datetime,
  updated_at               datetime,
  deleted_at               datetime,
  cart_id                  integer NOT NULL,
  single_item_coupon_id    integer,
  single_item_coupon_name  varchar(255),
  savings                  real
);

CREATE INDEX idx_applied_single_item_coupons_deleted_at
  ON applied_single_item_coupons
  (deleted_at);

CREATE TABLE applied_single_item_discounts (
  id                         integer PRIMARY KEY AUTOINCREMENT,
  created_at                 datetime,
  updated_at                 datetime,
  deleted_at                 datetime,
  cart_id                    integer NOT NULL,
  single_item_discount_id    integer,
  single_item_discount_name  varchar(255),
  savings                    real
);

CREATE INDEX idx_applied_single_item_discounts_deleted_at
  ON applied_single_item_discounts
  (deleted_at);

CREATE TABLE cart_items (
  id                     integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  cart_id                integer NOT NULL,
  fruit_id               integer NOT NULL,
  name                   varchar(255) NOT NULL,
  quantity               integer,
  item_total             real,
  item_discounted_total  real
);

CREATE TABLE carts (
  id             integer PRIMARY KEY AUTOINCREMENT,
  created_at     datetime,
  updated_at     datetime,
  deleted_at     datetime,
  customer_id    integer NOT NULL,
  "total"        real,
  total_savings  real,
  status         varchar(255)
);

CREATE INDEX idx_carts_deleted_at
  ON carts
  (deleted_at);

CREATE TABLE customers (
  login_id    varchar(255),
  first_name  varchar(255),
  last_name   varchar(255),
  id          integer PRIMARY KEY AUTOINCREMENT,
  created_at  datetime,
  updated_at  datetime,
  deleted_at  datetime
);

CREATE INDEX idx_customers_deleted_at
  ON customers
  (deleted_at);

CREATE UNIQUE INDEX uix_customers_login_id
  ON customers
  (login_id);

CREATE TABLE dual_item_discounts (
  id          integer PRIMARY KEY AUTOINCREMENT,
  created_at  datetime,
  updated_at  datetime,
  deleted_at  datetime,
  name        varchar(255),
  fruit_id    integer,
  fruit_id_1  integer,
  fruit_id_2  integer,
  count_1     integer,
  count_2     integer,
  discount    integer
);

CREATE INDEX idx_dual_item_discounts_deleted_at
  ON dual_item_discounts
  (deleted_at);

CREATE TABLE fruits (
  id          integer PRIMARY KEY AUTOINCREMENT,
  created_at  datetime,
  updated_at  datetime,
  deleted_at  datetime,
  name        varchar(255),
  price       real
);

CREATE INDEX idx_fruits_deleted_at
  ON fruits
  (deleted_at);

CREATE UNIQUE INDEX uix_fruits_name
  ON fruits
  (name);

CREATE TABLE payments (
  id           integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  customer_id  integer,
  cart_id      integer NOT NULL,
  amount       real,
  status       varchar(255)
);

CREATE TABLE single_item_coupons (
  id          integer PRIMARY KEY AUTOINCREMENT,
  created_at  datetime,
  updated_at  datetime,
  deleted_at  datetime,
  name        varchar(255),
  fruit_id    integer,
  discount    integer,
  duration    integer
);

CREATE INDEX idx_single_item_coupons_deleted_at
  ON single_item_coupons
  (deleted_at);

CREATE TABLE single_item_discounts (
  id          integer PRIMARY KEY AUTOINCREMENT,
  created_at  datetime,
  updated_at  datetime,
  deleted_at  datetime,
  name        varchar(255),
  fruit_id    integer,
  "count"     integer,
  discount    integer
);

CREATE INDEX idx_single_item_discounts_deleted_at
  ON single_item_discounts
  (deleted_at);

/* Data for table applied_dual_item_discounts */
INSERT INTO applied_dual_item_discounts (id, created_at, updated_at, deleted_at, cart_id, dual_item_discount_id, dual_item_discount_name, savings) VALUES (1, '2020-08-06 20:12:52', '2020-08-06 20:12:52', NULL, 1, 1, 'PEARBANANA30', 1.8);



/* Data for table applied_single_item_coupons */




/* Data for table applied_single_item_discounts */
INSERT INTO applied_single_item_discounts (id, created_at, updated_at, deleted_at, cart_id, single_item_discount_id, single_item_discount_name, savings) VALUES (1, '2020-08-06 20:12:54', '2020-08-06 20:12:54', NULL, 1, 1, 'APPLE10', 0.8);



/* Data for table cart_items */
INSERT INTO cart_items (id, cart_id, fruit_id, name, quantity, item_total, item_discounted_total) VALUES (1, 1, 2, '', 2, 1.4, 0.6);
INSERT INTO cart_items (id, cart_id, fruit_id, name, quantity, item_total, item_discounted_total) VALUES (2, 1, 3, '', 4, 2.8, 1.2);
INSERT INTO cart_items (id, cart_id, fruit_id, name, quantity, item_total, item_discounted_total) VALUES (3, 1, 1, '', 8, 7.2, 0.8);
INSERT INTO cart_items (id, cart_id, fruit_id, name, quantity, item_total, item_discounted_total) VALUES (4, 1, 4, '', 10, 10, 0);



/* Data for table carts */
INSERT INTO carts (id, created_at, updated_at, deleted_at, customer_id, "total", total_savings, status) VALUES (1, '2020-08-06 20:12:46', '2020-08-06 20:14:03', NULL, 1, 0, 2.6, 'CLOSED');
INSERT INTO carts (id, created_at, updated_at, deleted_at, customer_id, "total", total_savings, status) VALUES (2, '2020-08-06 20:14:03', '2020-08-06 20:14:03', NULL, 1, 0, 0, 'OPEN');



/* Data for table customers */
INSERT INTO customers (login_id, first_name, last_name, id, created_at, updated_at, deleted_at) VALUES ('a', 'a', 'a', 1, '2020-08-06 20:12:46', '2020-08-06 20:12:46', NULL);



/* Data for table dual_item_discounts */
INSERT INTO dual_item_discounts (id, created_at, updated_at, deleted_at, name, fruit_id, fruit_id_1, fruit_id_2, count_1, count_2, discount) VALUES (1, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'PEARBANANA30', 3, 3, 2, 4, 2, 30);



/* Data for table fruits */
INSERT INTO fruits (id, created_at, updated_at, deleted_at, name, price) VALUES (1, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'Apple', 1);
INSERT INTO fruits (id, created_at, updated_at, deleted_at, name, price) VALUES (2, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'Banana', 1);
INSERT INTO fruits (id, created_at, updated_at, deleted_at, name, price) VALUES (3, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'Pear', 1);
INSERT INTO fruits (id, created_at, updated_at, deleted_at, name, price) VALUES (4, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'Orange', 1);



/* Data for table payments */
INSERT INTO payments (id, customer_id, cart_id, amount, status) VALUES (1, 1, 1, 21.4, 'PAID');



/* Data for table single_item_coupons */
INSERT INTO single_item_coupons (id, created_at, updated_at, deleted_at, name, fruit_id, discount, duration) VALUES (1, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'ORANGE30', 4, 30, 10);



/* Data for table single_item_discounts */
INSERT INTO single_item_discounts (id, created_at, updated_at, deleted_at, name, fruit_id, "count", discount) VALUES (1, '2020-08-06 20:12:41', '2020-08-06 20:12:41', NULL, 'APPLE10', 1, 7, 10);

