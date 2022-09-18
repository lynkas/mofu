FROM golang:1.19-bullseye
WORKDIR /app
RUN apt install git -y
ADD go.* ./
RUN go mod download
ADD ./ ./
RUN go build -o /app/main -buildvcs=false

RUN chmod +x /app/main
CMD [ "/app/main" ]