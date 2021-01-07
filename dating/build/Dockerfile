# Собираем в гошке
FROM golang:1.15.2 as build

ENV BIN_FILE /opt/dating/dating-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}
RUN rm cmd/wire.go

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar"
LABEL MAINTAINERS="sterligov.denis94@yandex.ru"

ENV BIN_FILE "/opt/calendar/calendar-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/calendar/config.yml
COPY ./configs/config.yml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}