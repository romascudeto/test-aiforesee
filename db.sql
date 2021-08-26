CREATE TABLE fuels (
	id serial NOT NULL,
	fuel_liter int4 NULL,
	fuel_type varchar NULL,
	fuel_price float8 NULL,
	CONSTRAINT fuels_pk PRIMARY KEY (id)
);