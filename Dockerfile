FROM golang:1.20.6
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o server
ENTRYPOINT [ "/bin/bash", "-c" ]
CMD [ "./server" ]
