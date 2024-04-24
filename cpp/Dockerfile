FROM gcc:13

WORKDIR /app

RUN apt update && \
	apt install git && \
	git clone https://gitlab.com/libeigen/eigen.git

COPY . .

RUN g++ -I eigen one.cpp -O -o one && \
	g++ -I eigen two.cpp -O -o two && \
	g++ -I eigen three.cpp -O -o three

CMD /app/one && echo "\n" && /app/two && echo "\n" && /app/three && echo "\n"