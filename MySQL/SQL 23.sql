/*CASE Function*/

USE db2;
SELECT country_name,
CASE 
	WHEN country_name = "India" THEN "Hindi"
    WHEN country_name = "Pakistan" THEN "Urdu"
    WHEN country_name = "Sri Lanka" THEN "Sinhala"
    ELSE "Bengali"
END as Official_Language
FROM country;


