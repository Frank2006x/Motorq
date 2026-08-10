-- Create fleet table
CREATE TABLE fleet (
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- Create vehicle table
CREATE TABLE vehicle (
    id        BIGSERIAL PRIMARY KEY,
    fleet_id  BIGINT NOT NULL REFERENCES fleet(id),
    model     VARCHAR(255) NOT NULL,
    status    VARCHAR(50) NOT NULL DEFAULT 'active'
);

-- Create telemetry history table
CREATE TABLE telemetry_history (
    id               BIGSERIAL PRIMARY KEY,
    vehicle_id       BIGINT NOT NULL REFERENCES vehicle(id),
    timestamp        TIMESTAMP NOT NULL DEFAULT NOW(),
    speed_mph        DECIMAL(10, 2) NOT NULL,
    fuel_level_percent DECIMAL(5, 2) NOT NULL,
    engine_state     VARCHAR(50) NOT NULL,
    odometer_miles   DECIMAL(12, 2) NOT NULL,
    status           VARCHAR(20) NOT NULL DEFAULT 'pending'
);

-- Create priority enum
CREATE TYPE priority AS ENUM ('low', 'mid', 'high');

-- Create operator enum
CREATE TYPE rule_operator AS ENUM ('>', '<');

-- Create simple rule fleet table
CREATE TABLE simple_rule_fleet (
    id              BIGSERIAL PRIMARY KEY,
    fleet_id        BIGINT NOT NULL REFERENCES fleet(id),
    target_field    VARCHAR(100) NOT NULL,
    operator        rule_operator NOT NULL,
    threshold_value DECIMAL(12, 2) NOT NULL,
    priority        priority NOT NULL DEFAULT 'low'
);

-- Create simple rule vehicle table
CREATE TABLE simple_rule_vehicle (
    id              BIGSERIAL PRIMARY KEY,
    vehicle_id      BIGINT NOT NULL REFERENCES vehicle(id),
    target_field    VARCHAR(100) NOT NULL,
    operator        rule_operator NOT NULL,
    threshold_value DECIMAL(12, 2) NOT NULL,
    priority        priority NOT NULL DEFAULT 'low'
);