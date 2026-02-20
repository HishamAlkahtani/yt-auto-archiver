-- Up Migration

CREATE TABLE channels (
    channel_id VARCHAR(255) PRIMARY KEY,
    status VARCHAR(20) CHECK (status IN ('archiving', 'monitoring')),
    last_check TIMESTAMPTZ 
);

CREATE TABLE videos (
    video_id VARCHAR(255) PRIMARY KEY,
    channel_id VARCHAR(255) NOT NULL REFERENCES channels (channel_id) ON DELETE CASCADE,
    file_name TEXT,
    status VARCHAR(20) CHECK (status IN ('downloaded', 'pending'))
);

CREATE INDEX idx_pending_vids ON videos (channel_id)
WHERE status = 'pending';

CREATE INDEX idx_monitoring_channels ON channels (status, last_check)
WHERE status = 'monitoring';