n = 7500; % 5000
A = randn(n, n);
b = randn(n, 1);

tic;
X = linsolve(A, b);
sprintf('%0.3f', toc)

tic;
X = A \ b;
sprintf('%0.3f', toc)