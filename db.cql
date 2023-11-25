-- Create keyspace and table
CREATE KEYSPACE IF NOT EXISTS recruitment WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
USE recruitment;

CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY,
    title TEXT,
    company TEXT,
    location TEXT
);

