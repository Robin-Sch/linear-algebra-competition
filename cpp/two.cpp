#include <iostream>
#include <Eigen/Dense>
#include <chrono>

int main() {
	int n = 1000;

	Eigen::MatrixXd A = Eigen::MatrixXd::Random(n, n);

	auto start = std::chrono::steady_clock::now();
	Eigen::SelfAdjointEigenSolver<Eigen::MatrixXd> eigensolver(A);
	auto end = std::chrono::steady_clock::now();

	std::chrono::duration<double> time = end - start;
	std::cout << time.count();

	return 0;
}