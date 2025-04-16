FROM golang:1.24



COPY . /dozen

WORKDIR /dozen

EXPOSE 8000

RUN go mod tidy
CMD ["go", "run", "cmd/app/main.go"] 


