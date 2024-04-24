#include <iostream>
#include <Eigen/Dense>
#include <chrono>

int main() {
	int m = 1000;
	int n = 1000;

	Eigen::MatrixXd A = Eigen::MatrixXd::Random(m, n);

	auto start = std::chrono::steady_clock::now();
	Eigen::JacobiSVD<Eigen::MatrixXd> svd(A, Eigen::ComputeThinU | Eigen::ComputeThinV);
	auto end = std::chrono::steady_clock::now();

	std::chrono::duration<double> time = end - start;
	std::cout << time.count();

	return 0;
}