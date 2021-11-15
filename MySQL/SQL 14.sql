/*Logical Operators*/
USE db1;
SELECT * FROM student
WHERE age=18 AND place= 'Palakkad';

/*Between Operator*/
SELECT * FROM student
WHERE age BETWEEN 10 AND 15;

/*Exists Operator*/
SELECT * FROM student
WHERE EXISTS(SELECT age FROM student WHERE age<20);

/*IN Operator*/
SELECT * FROM student
WHERE place IN('Malappuram', 'Trivandrum');

/*LIKE Operator*/

SELECT * FROM student
WHERE student_name LIKE 'akhil';

SELECT * FROM student
WHERE student_name LIKE 'a%';

SELECT * FROM student
WHERE student_name LIKE '%u';

SELECT * FROM student
WHERE student_name LIKE '%o%';

/*NOT LIKE Operator*/
SELECT * FROM student
WHERE student_name NOT LIKE 'Akhil';

/*OR Operator*/
SELECT * FROM student
WHERE place ='Thrissur' OR age= 20;


