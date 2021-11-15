const cars = ["Saab", "Volvo", "BMW"];

const cars = [];
cars[0]= "Saab";
cars[1]= "Volvo";
cars[2]= "BMW";

const fruits = new Array("Banana", "Orange", "Apple"); //creates new array

const fruits = ["Banana", "Orange", "Apple"];
fruits[fruits.length] = "Lemon";  // Adds "Lemon" to fruits

const fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits.pop();  //removes last element

const fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits.push("Kiwi") //adds new element to the end

const fruits = ["Banana", "Orange", "Apple", "Mango"];
let fruit = fruits.shift();

const fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits[0] = "Kiwi"; //changing elements with index number

const fruits = ["Banana", "Orange", "Apple", "Mango"];
delete fruits[0];  //delete

const fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits.splice(2, 0, "Lemon", "Kiwi");  //adds new items

const fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits.splice(0, 1);  //removes items

const myGirls = ["Cecilie", "Lone"];
const myBoys = ["Emil", "Tobias", "Linus"];
const myChildren = myGirls.concat(myBoys);  //merging 

const fruits = ["Banana", "Orange", "Lemon", "Apple", "Mango"];
const citrus = fruits.slice(1); //array slice

//Sorting

const fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits.sort();
fruits.reverse();

const points = [40, 100, 1, 5, 25, 10];
points.sort(function(a, b){return a - b});  //ascending

const points = [40, 100, 1, 5, 25, 10];
points.sort(function(a, b){return b - a}); //descending

const cars = [
    {type:"Volvo", year:2016},
    {type:"Saab", year:2001},
    {type:"BMW", year:2010}
];
cars.sort(function(a, b){return a.year - b.year}); //sorting objects

//array const

const cars = ["Saab", "Volvo", "BMW"];
cars = ["Toyota", "Volvo", "Audi"];    // ERROR