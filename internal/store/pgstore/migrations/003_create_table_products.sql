CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    seller_id UUID NOT NULL REFERENCES users (id),

    product_name TEXT NOT NULL,
    description TEXT,

    baseprice DECIMAL NOT NULL,
    auction_end TIMESTAMPTZ NOT NULL,

    is_sold BOOLEAN NOT NULL DEFAULT false,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

---- create above / drop below ----

DROP TABLE IF EXISTS products;