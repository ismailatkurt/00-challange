# GO Repo base repo
FROM golang:1.20.4-alpine as builder

# Add Maintainer Info
LABEL maintainer="<>"

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go install
COPY go.mod ./

# Download all the dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o recipe-app .

# GO Repo base repo
FROM alpine:latest

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/recipe-app .

COPY inputs.json .

RUN wget https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz
RUN tar -xzf hf_test_calculation_fixtures.tar.gz

# Run Executable
CMD ["./recipe-app"]