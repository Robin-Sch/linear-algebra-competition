import numpy as np
import time

m = 2000 # 1000
n = 2000 # 1000
A = np.random.randn(m, n)

start = time.time()
S = np.linalg.svd(A, full_matrices=False)
end = time.time()

print(f"{round(end - start, 3)} s")