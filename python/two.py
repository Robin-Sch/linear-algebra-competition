import numpy as np
import time

n = 2000 # 1000
A = np.random.randn(n, n)

start = time.time()
x = np.linalg.eig(A)
end = time.time()

print(f"{round(end - start, 3)} s")