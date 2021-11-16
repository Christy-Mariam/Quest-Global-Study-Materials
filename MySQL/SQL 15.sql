/*Built-in Functions*/

/*Char_Length()*/

USE db1;
SELECT * FROM student;
SELECT place, char_length(place) AS Length
FROM student;

/*Concat()*/

SELECT * FROM student;
SELECT concat(student_name, " ",place) AS New_String
FROM student;

/*Format()*/

SELECT format(250500.564, 2) AS new_number;

/*Insert()*/

SELECT insert("Google", 1,1, "f");
SELECT insert("Google", 1,3, "fff");

/*upper()*/
SELECT upper("hello world") AS new_string;

/*lower()*/
SELECT lower("HELLO WORLD") AS new_string;

/*reverse()*/
SELECT reverse("HELLO WORLD") AS new_string;

/*repeat()*/
SELECT repeat("HELLO",5) AS new_string;

/*left()*/
SELECT left("HELLO",2) AS new_string;

/*right()*/
SELECT right("HELLO",3) AS new_string;

/*length()*/
SELECT length("HELLO") AS new_string;