FROM golang:1.22
LABEL authors="vamage"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o baconator
EXPOSE 8081
ENTRYPOINT ["./baconator"]