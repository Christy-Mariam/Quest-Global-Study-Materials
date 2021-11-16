/*Mathematical Functions*/

SELECT abs(-123) as new_number;

USE db1;
SELECT avg(age) as new_number FROM student;

SELECT ceiling(25.3) as new_number;

SELECT floor(25.8) as new_number;

SELECT round(25.4) as new_number;
SELECT round(25.43434, 2) as new_number;

SELECT count(student_id) as total_student
FROM student;

SELECT max(age) as max_age FROM student;

SELECT min(age) as min_age FROM student;

SELECT pi() AS pi_value;

SELECT rand() AS random_number;
SELECT floor(rand()*11) AS random_number;
SELECT floor(rand()*101) AS random_number;

SELECT sqrt(9) AS new_number;

SELECT sum(age) AS new_number FROM student;
