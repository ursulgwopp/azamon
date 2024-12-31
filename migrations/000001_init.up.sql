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
    seller TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL NOT NULL
);

CREATE TABLE purchases (
    id UUID PRIMARY KEY UNIQUE NOT NULL,
    buyer TEXT NOT NULL REFERENCES users(username),
    seller TEXT NOT NULL REFERENCES users(username),
    name TEXT NOT NULL REFERENCES items(name),
    quantity INTEGER NOT NULL,
    final_price DECIMAL NOT NULL,
    purchased_at TIMESTAMP NOT NULL
);

CREATE TABLE blacklist (
    token TEXT NOT NULL
);

