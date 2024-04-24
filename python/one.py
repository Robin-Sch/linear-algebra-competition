import numpy as np
import time

n = 11000 # 5000
A = np.random.randn(n, n)
b = np.random.randn(n)

start = time.time()
x = np.linalg.solve(A, b)
end = time.time()

print(f"{round(end - start, 3)} s")