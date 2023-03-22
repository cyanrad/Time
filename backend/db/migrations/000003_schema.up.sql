CREATE TYPE task_status AS ENUM ('to-do', 'doing', 'done', 'blocked');
CREATE TYPE task_type AS ENUM ('global', 'revolver', 'daily');

CREATE TABLE tasks (
    id bigserial PRIMARY KEY,
    date_stamp integer NOT NULL,
    text VARCHAR(64) NOT NULL,
    status task_status NOT NULL,
    type task_type NOT NULL,
    parent_id BIGINT
);