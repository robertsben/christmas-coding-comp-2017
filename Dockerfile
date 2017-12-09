FROM gcc
WORKDIR /usr/src
COPY ./src/main.c .
RUN gcc -o main -Ofast main.c
CMD ["./main"]