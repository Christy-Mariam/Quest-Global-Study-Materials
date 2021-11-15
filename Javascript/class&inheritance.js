class Car {
    constructor(name, year) {
        this.name = name;
        this.year = year;
    }
}

class Car {
    constructor(name, year) {
        this.name = name;
        this.year = year;
    }
    age() {
        let date = new Date();
        return date.getFullYear() - this.year;
    }
}
let myCar = new Car("Ford", 2014);
console.log("My car is " + myCar.age() + " years old.")


class Car {
    constructor(brand) {
        this.carname = brand;
    }
    present() {
        return 'I have a ' + this.carname;
    }
}
class Model extends Car {
    constructor(brand, mod) {
        super(brand);
        this.model = mod;
    }
    show() {
        return this.present() + ', it is a ' + this.model;
    }
}
let myCar = new Model("Ford", "Mustang");
let x = myCar.show();
console.log(x);