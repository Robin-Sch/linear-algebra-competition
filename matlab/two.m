n = 2000; % 1000
A = randn(n, n);

tic;
x = eig(A);
sprintf('%0.3f', toc)