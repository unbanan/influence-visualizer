CREATE SCHEMA IF NOT EXISTS influence;

CREATE TABLE influence.users(
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX users_id_index ON influence.users (id);

CREATE TABLE influence.strategies(
    uid BIGINT,
    sid UUID DEFAULT gen_random_uuid() UNIQUE,
    code TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    rate NUMBER DEFAULT NULL
);

CREATE INDEX strategies_uid_created_at_index ON influence.strategies (uid, created_at);
CREATE INDEX strategies_sid_index ON influence.strategies (sid);

CREATE TABLE influence.replays(
    id UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    data JSONB
);

CREATE INDEX replays_id_index ON influence.replays (id);

CREATE TABLE influence.maps(
    id UUID UNIQUE DEFAULT gen_random_uuid(),
    data JSONB,
    meta TEXT
)

CREATE INDEX influence.maps_id_index ON influence.maps (id);
