# Build stage  
FROM golang:1.24-alpine AS builder  
WORKDIR /app  

COPY . .  
RUN go mod init metatheif
RUN go get github.com/PuerkitoBio/goquery
RUN go mod tidy

RUN go mod download  

# Run tests
RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/metatheif ./cmd/metatheif  

# Final stage
FROM alpine:latest  
COPY --from=builder /bin/metatheif /bin/metatheif  
ENTRYPOINT ["/bin/metatheif"]  