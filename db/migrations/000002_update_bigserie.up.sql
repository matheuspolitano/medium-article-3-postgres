ALTER TABLE users ALTER COLUMN ID TYPE bigint;

CREATE SEQUENCE users_ID_seq;

-- Optional: Set the starting point of the sequence
SELECT setval('users_ID_seq', COALESCE((SELECT MAX(ID)+1 FROM users), 1), false);

-- Set the column to use the sequence
ALTER TABLE users ALTER COLUMN ID SET DEFAULT nextval('users_ID_seq'::regclass);


ALTER SEQUENCE users_ID_seq OWNED BY users.ID;
