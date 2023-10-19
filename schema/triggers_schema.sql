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
