\copy input (line) from 'input.txt';

SELECT sum(i)
FROM (SELECT (SELECT distinct l.item
              FROM contents l
                       JOIN contents r ON r.bag = b.id AND r.compartment = 1 AND r.item = l.item
              WHERE l.bag = b.id
                AND l.compartment = 0) i
      FROM bags b) AS x;

SELECT sum(item)
FROM (SELECT (bag - 1) / 3 AS g, item, COUNT(DISTINCT bag) AS numBags FROM contents GROUP BY g, item) AS i
WHERE i.numBags = 3;