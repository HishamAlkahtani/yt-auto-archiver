CREATE TABLE
    channels (
        channel_id TEXT PRIMARY KEY,
        status TEXT CHECK (status IN ('new', 'archiving', 'archived')),
        last_check INTEGER
    );

CREATE TABLE
    videos (
        video_id TEXT PRIMARY KEY,
        channel_id REFERENCES channels (channel_id) ON DELETE CASCADE,
        status TEXT CHECK (status IN ('downloaded', 'pending'))
    );

CREATE INDEX pending_vids ON videos (channel_id)
WHERE
    status = 'pending';