while (i < 10) {
    text += "The number is " + i;
    i++;
}


do {
    text += "The number is " + i;
    i++;
}
while (i < 10);


for (let i = 0; i < 5; i++) {
    text += "The number is " + i + "<br>";
}


const numbers = [45, 4, 9, 16, 25];
let txt = "";
numbers.forEach(myFunction);
function myFunction(value, index, array) {
  txt += value;
}