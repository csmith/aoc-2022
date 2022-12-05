\copy input (line) from 'input.txt';

SELECT string_agg(v, '')
FROM (SELECT (select value from crates_part1 where id = max(c.id)) v
      FROM crates_part1 c
      group by stack
      order by stack) AS n;

SELECT string_agg(v, '')
FROM (SELECT (select value from crates_part2 where id = max(c.id)) v
      FROM crates_part2 c
      group by stack
      order by stack) AS n;