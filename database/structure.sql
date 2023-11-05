CREATE DATABASE event_sourcing_example;
GRANT ALL PRIVILEGES ON DATABASE event_sourcing_example TO postgres;

\c event_sourcing_example;

CREATE TABLE order_views (
	id uuid NOT NULL,
	amount integer DEFAULT 0 NOT NULL,
	paid_amount integer DEFAULT 0 NOT NULL,
	status character varying
);

CREATE TABLE order_events (
  aggregate_id uuid NOT NULL,
	version integer DEFAULT 0 NOT NULL,
	event_type character varying,
	event_payload jsonb
);

CREATE UNIQUE INDEX idx_order_events_on_view_id_version ON order_events USING btree (view_id, version);