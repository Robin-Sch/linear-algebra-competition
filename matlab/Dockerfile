FROM mathworks/matlab:r2024a

WORKDIR /app

COPY . .

CMD matlab -batch one && matlab -batch two && matlab -batch three

# Make sure to run with -it since it asks for MathWorks account credentials