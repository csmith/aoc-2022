CREATE TABLE input
(
    line varchar
);

CREATE OR REPLACE FUNCTION isUnique(str varchar) returns bool as
$$
DECLARE
    i int;
BEGIN
    for i in reverse length(str)..1
        loop
            if position(substr(str, i, 1) in str) != i then
                return false;
            end if;
        end loop;
    return true;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION findMarker(str varchar, size int) returns int as
$$
DECLARE
    i     int;
    chars varchar;
BEGIN
    chars = '';

    for i in 1..length(str)
        loop
            if length(chars) = size then
                chars = substr(chars, 2) || substr(str, i, 1);
            else
                chars = chars || substr(str, i, 1);
            end if;

            if length(chars) = size and isUnique(chars) then
                return i;
            end if;
        end loop;
    return -1;
END;
$$ LANGUAGE plpgsql;