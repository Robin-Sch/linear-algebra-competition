const math = require('mathjs');

const m = 1000;
const n = 1000;

const A = math.random([m, n]);

const start = Date.now();
// const S = math.svd(A);
const end = Date.now();

console.log(((end - start) / 1000).toFixed(3));