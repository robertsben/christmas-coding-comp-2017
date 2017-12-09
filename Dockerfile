FROM gcc
COPY ./src /usr/src
WORKDIR /usr/src
RUN gcc -o main -Ofast main.c
CMD ["./main"]