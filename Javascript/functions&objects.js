//JS Functions
//============

function myFunction(p1, p2) {
    return p1 * p2;   // The function returns the product of p1 and p2
}

let x = myFunction(4, 3);   // Function is called, return value will end up in x
function myFunction(a, b) {
  return a * b;             // Function returns the product of a and b
}

//JS Objects
//==========

const person = {
    firstName: "John",
    lastName: "Doe",
    age: 50,
    eyeColor: "blue"
  };

const person = {
    firstName: "John",
    lastName : "Doe",
    id       : 5566,
    fullName : function() {
      return this.firstName + " " + this.lastName;
    }
};
