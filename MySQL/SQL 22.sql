/*IF()*/

SELECT if(10>20, "Value1", "Value2") AS result;

USE db1;
SELECT student_name, age, 
if(age>=18, "Adult", "Minor") AS student_type
FROM student;

/*IFNULL()*/

SELECT ifnull(null, "Hello") as result;
SELECT ifnull(2, "Hello") as result;

USE db2;
SELECT country.country_name, 
ifnull(capital.capital_name, "Not Defined") AS capital_name
FROM country LEFT JOIN capital
ON country.country_code = capital.country_code;

/*ISNULL()*/

SELECT isnull(null) as result;
SELECT isnull(2) as result;

USE db2;
SELECT country.country_name, capital.capital_name,
isnull(capital.capital_name) AS capital_name
FROM country LEFT JOIN capital
ON country.country_code = capital.country_code;

