CREATE TABLE moves
(
    name  varchar primary key,
    them  varchar unique,
    me    varchar unique,
    score int
);

INSERT INTO moves
VALUES ('Rock', 'A', 'X', 1),
       ('Paper', 'B', 'Y', 2),
       ('Scissors', 'C', 'Z', 3);

CREATE TABLE outcomes
(
    me    varchar,
    them  varchar,
    score int
);

INSERT INTO outcomes
VALUES ('Rock', 'Rock', 3),
       ('Rock', 'Paper', 0),
       ('Rock', 'Scissors', 6),
       ('Paper', 'Rock', 6),
       ('Paper', 'Paper', 3),
       ('Paper', 'Scissors', 0),
       ('Scissors', 'Rock', 0),
       ('Scissors', 'Paper', 6),
       ('Scissors', 'Scissors', 3);

CREATE TABLE targets
(
    code  varchar unique,
    score int unique
);

INSERT INTO targets
VALUES ('X', 0),
       ('Y', 3),
       ('Z', 6);

CREATE TABLE matches
(
    l varchar,
    r varchar
);

CREATE VIEW p1results AS
SELECT lm.name AS theirmove, rm.name AS mymove, o.score AS matchscore, o.score + rm.score AS totalscore
FROM matches
         JOIN moves lm ON lm.them = l
         JOIN moves rm ON rm.me = r
         JOIN outcomes o ON o.me = rm.name AND o.them = lm.name;

CREATE VIEW p2results AS
SELECT lm.name AS theirmove, rm.name AS mymove, o.score AS matchscore, o.score + rm.score AS totalscore
FROM matches
         JOIN moves lm ON lm.them = l
         JOIN targets t ON t.code = r
         JOIN outcomes o ON o.score = t.score AND o.them = lm.name
         JOIN moves rm ON rm.name = o.me;