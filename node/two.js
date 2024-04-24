const math = require('mathjs');

const n = 1000;

const A = math.random([n, n]);

const start = Date.now();
const x = math.eigs(A).eigenvectors;
const end = Date.now();

console.log(((end - start) / 1000).toFixed(3));