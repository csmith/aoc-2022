CREATE TABLE input
(
    line varchar
);

CREATE TABLE elves
(
    id serial primary key
);

CREATE TABLE calories
(
    elf     int,
    calorie int
);

CREATE OR REPLACE FUNCTION on_input() RETURNS trigger AS
$$
DECLARE
    elf integer;
BEGIN
    if NEW.line = '' then
        INSERT INTO elves DEFAULT VALUES RETURNING id INTO elf;
        INSERT INTO calories select elf, line::int FROM input;
        DELETE FROM input;
        return null;
    end if;
    return NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER onInput BEFORE INSERT ON input FOR EACH ROW EXECUTE PROCEDURE on_input();