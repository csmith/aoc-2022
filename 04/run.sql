\copy input (elf1, elf2) from 'input.txt' with delimiter ',';

SELECT COUNT(*) FROM pairs WHERE (elf1start >= elf2start AND elf1end <= elf2end) OR (elf2start >= elf1start AND elf2end <= elf1end);
SELECT COUNT(*) FROM pairs WHERE (elf1start >= elf2start AND elf1start <= elf2end) OR (elf2start >= elf1start AND elf2start <= elf1end);