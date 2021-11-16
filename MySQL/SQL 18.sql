/*Foreign Key*/

USE db1;
CREATE TABLE courses(
course_id VARCHAR(5) NOT NULL,
course_name VARCHAR(50) NOT NULL,
PRIMARY KEY (course_id)
);

INSERT INTO courses
VALUES
('c1', 'Computer Hardware'),
('c2', 'Networking'),
('c3', 'Web Designing'),
('c4', 'Graphic Designing');

SELECT * FROM courses;

