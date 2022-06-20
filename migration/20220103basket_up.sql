CREATE TABLE IF NOT EXISTS baskets (
	id 		    serial NOT NULL,
	product_id  int NOT NULL,
	user_id     int NOT NULL,
	count       integer NOT NULL,
    added_time  timestamp default CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT basket_pkey PRIMARY KEY (id),
	CONSTRAINT baskets_product_id_fkey FOREIGN KEY (product_id) REFERENCES products(id),
	CONSTRAINT baskets_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);