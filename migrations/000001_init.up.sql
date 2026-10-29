CREATE TABLE threats (
    md5 VARCHAR(32) PRIMARY KEY,
    threat_name TEXT NOT NULL,
    threat_description TEXT NOT NULL,
    severity            TEXT NOT NULL,
    threat_type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,

    CONSTRAINT threats_md5_format
        CHECK (md5 ~ '^[a-fA-F0-9]{32}$')
)