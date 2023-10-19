CREATE TABLE IF NOT EXISTS tenants (
    id int GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY(id)
);
ALTER SEQUENCE tenants_id_seq RESTART WITH 1000;

CREATE TABLE IF NOT EXISTS orders (
    id bigint GENERATED ALWAYS AS IDENTITY,
    tenant_id int NOT NULL,
    title VARCHAR(50) NOT NULL,
    order_nr VARCHAR(50),
    status INT NOT NULL DEFAULT 1,
    note TEXT,
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
    note TEXT,
    status INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
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
