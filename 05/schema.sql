CREATE TABLE input
(
    line varchar
);

-- If crates are always inserted from the bottom up, and removed/reinserted when moved,
-- then the serial will allow us to order the ones in each stack (highest on top) regardless
-- of what happens.
CREATE TABLE crates_part1
(
    id    serial primary key,
    stack int,
    value char
);

CREATE TABLE crates_part2
(
    id    serial primary key,
    stack int,
    value char
);

CREATE OR REPLACE FUNCTION move(number int, source int, dest int) returns void as
$$
BEGIN
    -- Part 1
    WITH targets AS (DELETE FROM crates_part1 WHERE id IN
                                              (SELECT id FROM crates_part1 WHERE stack = source ORDER BY id DESC LIMIT number) RETURNING id, value)
    INSERT
    INTO crates_part1 (stack, value)
    SELECT dest, value FROM targets order by id desc;

    -- Part 2
    WITH targets AS (DELETE FROM crates_part2 WHERE id IN
                                                    (SELECT id FROM crates_part2 WHERE stack = source ORDER BY id DESC LIMIT number) RETURNING id, value)
    INSERT
    INTO crates_part2 (stack, value)
    SELECT dest, value FROM targets order by id asc;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION on_input() RETURNS trigger AS
$$
DECLARE
    l       varchar;
    rows    int;
    i       int;
    j       int;
    matches varchar[];
BEGIN
    IF NEW.line = '' THEN
        SELECT COUNT(*) - 1 FROM input INTO rows;
        for i in reverse rows - 1..0
            loop
                select line from input offset i limit 1 into l;
                for j in 0..length(l) / 4
                    loop
                        INSERT INTO crates_part1(stack, value) VALUES (j + 1, substr(l, 2 + 4 * j, 1)::char);
                        INSERT INTO crates_part2(stack, value) VALUES (j + 1, substr(l, 2 + 4 * j, 1)::char);
                    end loop;
            end loop;

        DELETE FROM input;
        DELETE FROM crates_part1 WHERE value = ' ';
        DELETE FROM crates_part2 WHERE value = ' ';
        return null;
    END IF;
    IF NEW.line ~ '^move' THEN
        SELECT regexp_match(NEW.line, 'move (\d+) from (\d+) to (\d+)') INTO matches;
        perform move(matches[1]::int, matches[2]::int, matches[3]::int);
        return null;
    end if;
    return NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER onInput
    BEFORE INSERT
    ON input
    FOR EACH ROW
EXECUTE PROCEDURE on_input();