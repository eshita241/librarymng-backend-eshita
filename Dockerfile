FROM golang:latest as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -o librarymng-backend

FROM golang:latest

WORKDIR /root/

# copy the binary
COPY --from=builder /app/librarymng-backend . 
#  copy the env variable file
COPY --from=builder /app/.env .

# or instead of lines 12-19 directly run env vairable via command: docker run -e PORT=3000 -p 3000:3000 librarymng-backend
# Lines 12-13 copy contents of env into the binary command: librarymng-backend
EXPOSE 8080

CMD ["./librarymng-backend"]