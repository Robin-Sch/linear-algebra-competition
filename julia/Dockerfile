FROM julia:1.10

WORKDIR /app

RUN julia -e 'import Pkg; Pkg.add("BenchmarkTools")'

COPY . .

CMD julia /app/one.jl && julia /app/two.jl && julia /app/three.jl