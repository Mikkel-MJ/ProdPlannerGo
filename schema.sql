CREATE TABLE IF NOT EXISTS tenants (
    id int GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS orders (
    id bigint GENERATED ALWAYS AS IDENTITY,
    tenant_id int NOT NULL,
    title VARCHAR(50) NOT NULL,
    order_nr VARCHAR(50),
    status INT NOT NULL DEFAULT 0,
    note TEXT Default '',
    deadline_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(tenant_id) REFERENCES tenants(id)
);

CREATE TABLE IF NOT EXISTS tasks (
    id bigint GENERATED ALWAYS AS IDENTITY,
    tenant_id int NOT NULL,
    order_id bigint NOT NULL,
    title VARCHAR(50) NOT NULL,
    note TEXT Default '',
    status INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(tenant_id) REFERENCES tenants(id),
    FOREIGN KEY(order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS task_rankings (
    id bigint GENERATED ALWAYS AS IDENTITY,
    tenant_id int NOT NULL,
    order_id bigint NOT NULL,
    rank int[] NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(tenant_id) REFERENCES tenants(id),
    FOREIGN KEY(order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS order_templates (
    id bigint GENERATED ALWAYS AS IDENTITY,
    tenant_id int NOT NULL,
    title VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(tenant_id) REFERENCES tenants(id)
);

CREATE TABLE IF NOT EXISTS task_templates (
    id bigint GENERATED ALWAYS AS IDENTITY,
    tenant_id int NOT NULL,
    title VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(tenant_id) REFERENCES tenants(id)
);

CREATE TABLE IF NOT EXISTS order_task_templates (
    order_template_id bigint NOT NULL,
    task_template_id bigint NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY(order_template_id,task_template_id),
    FOREIGN KEY(order_template_id) REFERENCES order_templates(id),
    FOREIGN KEY(task_template_id) REFERENCES task_templates(id)
);

ALTER SEQUENCE tenants_id_seq RESTART WITH 1000;

DROP FUNCTION IF EXISTS before_update_updated_at() CASCADE;

CREATE OR REPLACE FUNCTION before_update_updated_at() RETURNS trigger AS
$BODY$
BEGIN
    IF row(NEW.*::text) IS DISTINCT FROM row(OLD.*::text) THEN
        NEW.updated_at = now();
    END IF;

    RETURN NEW;
END;
$BODY$

LANGUAGE plpgsql;

-- Update nonce code
-- CREATE FUNCTION update_nonce()
--   RETURNS TRIGGER
--   LANGUAGE plpgsql AS
-- $func$
-- BEGIN
--    UPDATE games AS g
--       SET    nonce = g.nonce + 1
--       WHERE  g.id = NEW.game_id;  -- fixed
   
--    RETURN NULL;      -- for AFTER trigger this can be NULL
-- END
-- $func$;

DO $BODY$

    

DECLARE t text;
BEGIN
-- EXECUTE format('
--                 CREATE TRIGGER log_nonce_count_update
--     SELECT ON server_seeds
--     FOR EACH ROW EXECUTE PROCEDURE update_nonce();
--             ');
    FOR t IN
        SELECT table_name
        FROM information_schema.columns
        WHERE (
            column_name = 'updated_at'
            AND (
                SELECT 1
                FROM information_schema.triggers
                WHERE trigger_name = 'before_update_updated_at_' || table_name
            ) IS NULL
        )
    LOOP
        EXECUTE format('
            CREATE TRIGGER before_update_updated_at_%s
            BEFORE UPDATE ON %I
            FOR EACH ROW EXECUTE PROCEDURE before_update_updated_at();
        ', t, t);
    END loop;
END;
$BODY$

LANGUAGE plpgsql;
