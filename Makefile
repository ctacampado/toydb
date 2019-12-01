TARGET = toydb
BIN = ./bin
CMD_BLD = go build
FLAGS = -o ${TARGET}
FILES = cmd/toykv/main.go

all: build
clean:
	rm ${BIN}/${TARGET}
build:
	${CMD_BLD} ${FLAGS} ${FILES}; mv ${TARGET} ${BIN}
