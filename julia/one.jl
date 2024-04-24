using LinearAlgebra
using Random
using BenchmarkTools

n = 7500; # 5000
A = randn(n, n);
b = randn(n);

@btime x = A \ b;