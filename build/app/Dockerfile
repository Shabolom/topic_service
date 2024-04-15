FROM golang:1.22.2 AS base
WORKDIR /app

RUN mkdir /out

COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM base AS build
ENV APP_NAME="topic_service"
WORKDIR /app
ADD . /app
RUN go build -o /out/${APP_NAME} cmd/main.go


FROM golang:1.22.2 AS release
ENV APP_NAME="topic_service"
WORKDIR /app

COPY --from=build /out/${APP_NAME} /app/
CMD ["/app/topic_service"]
