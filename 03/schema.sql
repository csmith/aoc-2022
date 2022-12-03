CREATE TABLE input
(
    line varchar
);

CREATE TABLE bags
(
    id serial primary key
);

CREATE TABLE contents
(
    bag         int,
    compartment int,
    item        int
);

CREATE OR REPLACE FUNCTION on_input() RETURNS trigger AS
$$
DECLARE
    x           integer;
    bag         integer;
    compartment integer;
    code        integer;
BEGIN
    INSERT INTO bags DEFAULT VALUES RETURNING id INTO bag;
    for x in 0..length(NEW.line)-1
        loop
            if x >= length(NEW.line) / 2 then
                compartment = 1;
            else
                compartment = 0;
            end if;

            code = ascii(substring(NEW.line from x+1 for 1));
            if code > 96 then
                code = code - 96;
            else
                code = code - 38;
            end if;

            INSERT INTO contents VALUES (bag, compartment, code);
        end loop;
    return null;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER onInput
    BEFORE INSERT
    ON input
    FOR EACH ROW
EXECUTE PROCEDURE on_input();