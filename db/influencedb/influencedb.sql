CREATE SCHEMA IF NOT EXISTS influence;

CREATE TABLE influence.users(
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX users_id_index ON influence.users (id);

CREATE TABLE influence.strategies(
    id UUID DEFAULT gen_random_uuid() UNIQUE,
    uid BIGINT,
    code TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    rate NUMBER DEFAULT NULL
);

CREATE INDEX strategies_uid_created_at_index ON influence.strategies (uid, created_at);
CREATE INDEX strategies_id_index ON influence.strategies (id);

CREATE TYPE influence.simulation_state AS ENUM (
    'Queued',
    'Compiling',
    'Running',
    'Success',
    'Failed'
);

CREATE TABLE influence.users_simulations(
    uid BIGINT NOT NULL REFERENCES influence.users(id),
    sid UUID NOT NULL REFERENCES influence.simulations(id),
    order INT NOT NULL
)

CREATE INDEX users_simulations_uid_index ON influence.users_simulations (uid);
CREATE INDEX users_simulations_sid_index ON influence.users_simulations (sid);

CREATE TABLE influence.simulations(
    id UUID UNIQUE DEFAULT gen_random_uuid(),
    map_id UUID NOT NULL REFERENCES influence.maps(id),
    data JSONB,
    queued_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP WITH TIME ZONE,
    finished_at TIMESTAMP WITH TIME ZONE,
    state influence.simulation_state
);

CREATE INDEX simulations_id_index ON influence.replays (id);
CREATE INDEX simulations_created_at_index ON influence.replays (created_at);

CREATE TABLE influence.maps(
    id UUID UNIQUE DEFAULT gen_random_uuid(),
    data JSONB,
    name TEXT NOT NULL,
    meta TEXT
)

CREATE INDEX influence.maps_id_index ON influence.maps (id);
