CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    hash_password TEXT NOT NULL,
    balance DECIMAL DEFAULT 0,
    items_list UUID[] DEFAULT '{}'
);

CREATE TABLE items (
    id UUID PRIMARY KEY UNIQUE NOT NULL,
    seller_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    quantity INTEGER NOT NULL
);

CREATE TABLE purchases (
    id UUID PRIMARY KEY UNIQUE NOT NULL,
    buyer_id INTEGER NOT NULL REFERENCES users(id),
    seller_id INTEGER NOT NULL REFERENCES users(id),
    item_id UUID NOT NULL REFERENCES items(id),
    quantity INTEGER NOT NULL,
    final_price DECIMAL NOT NULL,
    purchased_at TIMESTAMP NOT NULL
);

CREATE TABLE blacklist (
    token TEXT NOT NULL
);

