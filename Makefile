TARGET = toydb
BIN = ./bin
CMD_BLD = go build
FLAGS = -o ${TARGET}
FILES = cmd/toydb/main.go

all: build
clean:
	rm ${BIN}/${TARGET}
build:
	if [ ! -d "bin" ]; then \
		mkdir bin; \
	fi
	${CMD_BLD} ${FLAGS} ${FILES}; mv ${TARGET} ${BIN}
