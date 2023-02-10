# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /learnapigo

EXPOSE 8080

CMD [ "/learnapigo" ]

#https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#labelling-container-images
# docker build -t learnapigo .
# docker images
# echo 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 | docker login ghcr.io -u fndenisovna --password-stdin
# Login Succeeded
# docker push ghcr.io/fndenisovna/learnapigo:latest