
//https://www.freecodecamp.org/news/higher-order-functions-in-javascript-explained/
function callbackFunc() {
    console.log('I am a callback func')
}

function ravikiran() {
    console.log('I am ravikiran')
}

function higherOrderFunc(func) {
    console.log('I am higher order function')
    func()
}

higherOrderFunc(callbackFunc)
higherOrderFunc(ravikiran)

const radius = [1, 2, 3];

const calculateArea = function (radius) {
    const output = [];
    for (let i = 0; i < radius.length; i++) {
        output.push(Math.PI * radius[i] * radius[i]);
    }
    return output;
}

const calculateDiameter = function (radius) {
    const output = [];

    for (let i = 0; i < radius.length; i++) {
        output.push(2 * radius[i]);
    }

    return output;
}

console.log(calculateArea(radius));
console.log(calculateDiameter(radius))

const arr = [1, 2, 3, 4, 5];
const output = arr.map((num) => num + 10)
console.log(arr);
console.log(output);