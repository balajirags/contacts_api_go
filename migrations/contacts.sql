CREATE SEQUENCE contacts_seq;
CREATE TABLE IF NOT EXISTS "contacts" (
  id integer PRIMARY KEY default nextval('contacts_seq'),
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  email varchar(50) NOT NULL,
  address varchar(100) NOT NULL,
  phone_number integer NOT NULL
);