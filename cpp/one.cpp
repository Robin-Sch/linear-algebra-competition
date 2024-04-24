#include <iostream>
#include <Eigen/Dense>
#include <chrono>

int main() {
	int n = 5000;

	Eigen::MatrixXd A = Eigen::MatrixXd::Random(n, n);
	Eigen::VectorXd b = Eigen::VectorXd::Random(n);

	auto start = std::chrono::steady_clock::now();
	Eigen::VectorXd x = A.lu().solve(b);
	auto end = std::chrono::steady_clock::now();

	std::chrono::duration<double> time = end - start;
	std::cout << time.count();

	return 0;
}