# Linear Algebra Competition

In this repo we have compared the speed of a few programming languages while solving linear algebra problems

### Languages
- C++
- Go
- Julia
- Matlab
- Node.js
- Python
- Rust

### Tasks
- Solving a linear system of a (random) 5000^2 matrix
- Computing eigenvectors of a (random) 1000^2 matrix
- Computing singular vectors of a (random) 1000^2 matrix

### Results

Each test is run 5x and the mean value is shown below

| Language | Linear System (5000^2) [s] | EigenVector (1000^2) [s] | Singular vectors (1000^2) [s] |
| --- | --- | --- | --- |
| C++ | 14.31 | 2.18 | 107.58 |
| Go | 17.93 | 7.18 | 13.36 |
| Julia | 3.52 | 1.92 | 0.77 |
| Matlab | 2.86 vs 2.47 | 1.12 | 0.50 |
| Node.js | 10900 | x | x |
| Python | 3.15 | 1.68 | 0.80 |
| Rust | 61.29 | 2.43 | 12.32 |

Note: for Matlab we used both "linsolve" and "A / b" respectively, which explains why there are 2 entries
Note2: Node.js task 1 ran overnight (once), task 2 failed and 3 is not possible with mathjs

### Observations:
All languages except Julia, Matlab and Python have significant longer running times for all tasks. Therefore, we only consider those languages and run the test again but with slightly bigger matrices.

- Solving a linear system of a (random) 7500^2 matrix
- Computing eigenvectors of a (random) 2000^2 matrix
- Computing singular vectors of a (random) 2000^2 matrix

| Julia | 10.21 | 10.93 | 7.33 |
| Matlab | 8.58 vs 7.82 | 5.76 | 3.434 |
| Python | 24.09 | 10.00 | 6.46 |

Note: for Matlab we used both "linsolve" and "A / b" respectively, which explains why there are 2 entries

As you can see, Matlab is considerable faster than both Julia and Python for all 3 tasks.