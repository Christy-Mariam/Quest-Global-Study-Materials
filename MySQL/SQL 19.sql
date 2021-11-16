USE db1;

CREATE TABLE enrolment(
enrolment_id VARCHAR(5) NOT NULL,
student_id VARCHAR(10),
course_id VARCHAR(10),
PRIMARY KEY (enrolment_id),
FOREIGN KEY (student_id) REFERENCES student(student_id),
FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

INSERT INTO enrolment
VALUES
('e1', 's1', 'c1'),
('e2', 's4', 'c4'),
('e3', 's2', 'c3'),
('e4', 's3', 'c2');

SELECT * FROM enrolment;