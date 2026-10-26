CREATE TABLE sites (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    domain TEXT NOT NULL UNIQUE,
    first_seen_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_scanned_at TIMESTAMPTZ
);

CREATE TABLE scans (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    site_id BIGINT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    scanned_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    status_code INT,
    title TEXT,
    detections JSONB NOT NULL DEFAULT '[]',
    error TEXT,
    body_size INT,
    duration_ms INT
);

CREATE INDEX scans_site_scanned_idx ON scans (site_id, scanned_at DESC);
CREATE INDEX scans_detections_gin ON scans USING GIN (detections);
