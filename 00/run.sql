\copy depths (depth) from '/code/00/input.txt';

WITH leads AS (SELECT depth < LEAD(depth, 1) OVER (order by id) increasing
               FROM depths)
SELECT COUNT(*)
FROM leads
WHERE increasing;

WITH windows AS (SELECT id, depth + LEAD(depth, 1) OVER (order by id) + LEAD(depth, 2) OVER (order by id) w
                 FROM depths),
     increasing AS (SELECT w < LEAD(w, 1) OVER (order by id) i
                    FROM windows)
SELECT COUNT(*)
FROM increasing
WHERE i;