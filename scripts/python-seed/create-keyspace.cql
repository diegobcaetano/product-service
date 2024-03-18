CREATE KEYSPACE IF NOT EXISTS product_keyspace WITH REPLICATION = {'class': 'NetworkTopologyStrategy', 'datacenter1': 1};
USE product_keyspace;
CREATE TABLE IF NOT EXISTS products (
    id text PRIMARY KEY,
    name text,
    categories list<text>,
    custom_attributes text, -- JSON-serialized map
    created_at timestamp,
    is_blocked boolean,
    images list<text>
);
CREATE TABLE IF NOT EXISTS buy_options (
    product_id text,
    seller_id text,
    price int,
    promotional_discount tinyint,
    stock smallint,
    PRIMARY KEY (product_id, seller_id)
);