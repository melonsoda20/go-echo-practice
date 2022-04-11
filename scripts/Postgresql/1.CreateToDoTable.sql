CREATE TABLE IF NOT EXISTS todo (
	id UUID PRIMARY KEY  DEFAULT gen_random_uuid (),
    name VARCHAR NULL,
	age int NULL,
    created_by VARCHAR NOT NULL,
    created_date timestamp NOT NULL,
	updated_by VARCHAR NULL,
	updated_date timestamp NULL
);