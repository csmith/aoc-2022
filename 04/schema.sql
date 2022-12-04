CREATE TABLE input
(
    elf1 varchar,
    elf2 varchar
);

CREATE TABLE pairs
(
    id        serial primary key,
    elf1start int,
    elf1end   int,
    elf2start int,
    elf2end   int
);

CREATE OR REPLACE FUNCTION on_input() RETURNS trigger AS
$$
BEGIN
    INSERT INTO pairs (elf1start, elf1end, elf2start, elf2end)
    VALUES (split_part(NEW.elf1, '-', 1)::int,
            split_part(NEW.elf1, '-', 2)::int,
            split_part(NEW.elf2, '-', 1)::int,
            split_part(NEW.elf2, '-', 2)::int);
    return null;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER onInput
    BEFORE INSERT
    ON input
    FOR EACH ROW
EXECUTE PROCEDURE on_input();