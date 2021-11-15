SELECT * FROM student;

ALTER TABLE student
RENAME COLUMN contact TO student_contact;

ALTER TABLE student
DROP COLUMN student_contact;

TRUNCATE student1;
DROP TABLE student1;

SELECT * FROM student
ORDER BY student_id ASC 
LIMIT 1;

SELECT * FROM student
ORDER BY student_id DESC 
LIMIT 2;

SELECT * FROM student
ORDER BY rand() 
LIMIT 1;

SELECT * FROM student;
SELECT student_name AS 'first_name',
age, place
FROM student;