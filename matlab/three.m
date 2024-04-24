m = 2000; % 1000
n = 2000; % 1000
A = randn(m, n);

tic;
S = svd(A, 0);
sprintf('%0.3f', toc)