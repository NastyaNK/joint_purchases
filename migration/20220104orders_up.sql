CREATE TABLE IF NOT EXISTS orders (
	id 		serial NOT NULL,
	product_id int NOT NULL,
	user_id int NOT NULL,
	count      integer not null,
	CONSTRAINT order_pkey PRIMARY KEY (id),
	CONSTRAINT orders_product_id_fkey FOREIGN KEY (product_id) REFERENCES products(id),
	CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);