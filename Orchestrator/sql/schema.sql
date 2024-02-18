CREATE TYPE expression_status as ENUM('wait','processing','invalid','done');

CREATE TABLE expressions (
    id BIGSERIAL PRIMARY KEY,
    expression text NOT NULL,
    status expression_status NOT NULL,
    result double precision ,
    created_at timestamptz NOT NULL DEFAULT now(),
    calculated_at timestamptz
);

CREATE TYPE operation_symbol as ENUM ('+','-','*','/');

CREATE TABLE operations (
    id BIGSERIAL PRIMARY KEY,
    operation operation_symbol NOT NULL,
    cost int NOT NULL DEFAULT 1
);

CREATE TABLE sub_expressions (
    id BIGSERIAL PRIMARY KEY,
    expression_id bigint NOT NULL REFERENCES expressions(id),
    expression text NOT NULL,
    exec_order int NOT NULL,
    result double precision
);

CREATE TABLE taken_expressions (
    expression_id bigint NOT NULL REFERENCES expressions(id),
    agent text NOT NULL,
    calculator int NOT NULL
);
