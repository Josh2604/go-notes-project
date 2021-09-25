DROP DATABASE IF EXISTS NotesProject;
CREATE DATABASE NotesProject;

CREATE TABLE Notes(
  id SERIAL NOT NULL PRIMARY KEY,
  name varchar(30) DEFAULT '',
  description varchar(100) DEFAULT '',
  deleted boolean DEFAULT false,
  date_created timestamp DEFAULT null,
  date_updated timestamp DEFAULT null,
  date_deleted timestamp DEFAULT null
);

INSERT INTO
  Notes (name,description,date_created)
VALUES
  ('Test name 1', 'Nota de test 1', NOW()),
  ('Test name 2', 'Nota de test 2', NOW());