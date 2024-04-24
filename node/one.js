const math = require('mathjs');

const n = 5000;

const A = math.random([n, n]);
const b = math.random([n]);

const start = Date.now();
const x = math.lusolve(A, b);
const end = Date.now();

console.log(((end - start) / 1000).toFixed(3));