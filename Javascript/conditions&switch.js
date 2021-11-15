var age = 15;
if (age > 18) {
    console.log("Qualifies for driving");
} else {
    console.log("Does not qualify for driving");
}


var book = "maths";
if (book == "history") {
    document.write("<b>History Book</b>");
} else if (book == "maths") {
    document.write("<b>Maths Book</b>");
} else if (book == "economics") {
    document.write("<b>Economics Book</b>");
} else {
    document.write("<b>Unknown Book</b>");
}


switch (new Date().getDay()) {
    case 6:
      text = "Today is Saturday";
      break;
    case 0:
      text = "Today is Sunday";
      break;
    default:
      text = "Looking forward to the Weekend";
}

let x = "0";
switch (x) {
  case 0:
    text = "Off";
    break;
  case 1:
    text = "On";
    break;
  default:
    text = "No value found";
}