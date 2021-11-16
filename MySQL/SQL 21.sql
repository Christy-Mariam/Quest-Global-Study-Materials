/*Inner Join*/

SELECT *
FROM country INNER JOIN capital
ON country.country_code = capital.country_code;

SELECT country.country_name, capital.capital_name
FROM country INNER JOIN capital
ON country.country_code = capital.country_code;

/*Left Join*/

SELECT *
FROM country LEFT JOIN capital
ON country.country_code = capital.country_code;

/*Right Join*/

SELECT *
FROM country RIGHT JOIN capital
ON country.country_code = capital.country_code;

/*Full Join*/

SELECT *
FROM country LEFT JOIN capital
ON country.country_code = capital.country_code
UNION
SELECT *
FROM country RIGHT JOIN capital
ON country.country_code = capital.country_code;
