CREATE TYPE activity_type AS ENUM ('Work', 'Positive', 'People', 'Neutral', 'University', 'Sleep', 'Rest', 'Useless', 'Mandatory', 'Unknown');

CREATE TABLE periods (
    id bigserial PRIMARY KEY,
    date_stamp integer NOT NULL,
    location smallint NOT NULL,
    duration smallint NOT NULL,
    type activity_type NOT NULL,
    title VARCHAR(64) NOT NULL,
    description VARCHAR(512) NOT NULL
);