\copy input (line) from 'input.txt';
-- Just in case there's not a blank line at the end...
INSERT INTO input VALUES ('');

SELECT SUM(calorie) m FROM calories GROUP BY elf ORDER BY m DESC LIMIT 1;

SELECT SUM(m) FROM (SELECT SUM(calorie) m FROM calories GROUP BY elf ORDER BY m DESC LIMIT 3) AS c;