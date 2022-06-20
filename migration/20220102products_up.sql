CREATE TABLE IF NOT EXISTS products (
	id 			  SERIAL NOT NULL,
	"name" 		  varchar(255) NOT NULL,
    "required"	  int NOT NULL,
	"description" text NOT NULL,
	price         float NOT NULL,
	start_time	  timestamp default CURRENT_TIMESTAMP NOT NULL,
    end_time	  timestamp default CURRENT_TIMESTAMP NOT NULL,
	organizer     integer NOT NULL,
	image         varchar(1000),
	CONSTRAINT product_pkey PRIMARY KEY (id),
	CONSTRAINT product_organizer_fkey FOREIGN KEY (organizer) REFERENCES users(id)
);