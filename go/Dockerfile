FROM golang:1.22

WORKDIR /app

RUN go mod init linear_algebra_competition && go get gonum.org/v1/gonum/...@latest

COPY . .

RUN go build -o one one.go && \
	go build -o two two.go && \
	go build -o three three.go

CMD /app/one && echo "\n" && /app/two && echo "\n" && /app/three && echo "\n"