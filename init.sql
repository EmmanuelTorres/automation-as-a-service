CREATE TABLE person(
	id bigserial primary key,
	username varchar not null unique,
	email varchar not null unique,
	password varchar not null,
	role varchar not null
);

CREATE TABLE country(
	id bigserial primary key,
	name varchar not null unique
);

CREATE TABLE designer(
	id bigserial primary key,
	name varchar not null,
	country_id bigint not null,
	CONSTRAINT fk_country
		FOREIGN KEY(country_id)
			REFERENCES country(id)
);

CREATE TABLE brand(
	id bigserial primary key,
	name varchar not null,
	country_id bigint not null,
	CONSTRAINT fk_country
		FOREIGN KEY(country_id)
			REFERENCES country(id)
);

CREATE TABLE garment(
	id bigserial primary key,
	code varchar not null,
	designer_id bigint not null,
	brand_id bigint not null,
	CONSTRAINT fk_designer
		FOREIGN KEY(designer_id)
			REFERENCES designer(id),
	CONSTRAINT fk_brand
		FOREIGN KEY(brand_id)
			REFERENCES brand(id)
);