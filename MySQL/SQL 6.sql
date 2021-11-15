SELECT * FROM student;
CREATE TABLE student1 LIKE student;
INSERT student1 SELECT * FROM student;
SELECT * FROM student1;

DELETE FROM student1
WHERE student_id = 's2';

SELECT student_id, place FROM student;