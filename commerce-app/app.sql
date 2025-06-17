DROP TABLE product;

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    price NUMERIC(10,2)
);

INSERT INTO products (name, price) VALUES
  ('Apple MacBook Pro 14″',       1999.99),
  ('Logitech MX Master 3 Mouse',    99.99),
  ('Anker USB-C to HDMI Adapter',   19.99),
  ('Samsung T7 Portable SSD 1 TB', 129.99),
  ('Sony WH-1000XM5 Headphones',   349.95);
