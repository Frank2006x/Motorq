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