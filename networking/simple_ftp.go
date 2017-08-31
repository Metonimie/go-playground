package main

import (
	"net"
	"log"
	"io"
	"os"
	"io/ioutil"
	"bufio"
	"strings"
	"errors"
	"bytes"
	"strconv"
)

const PATH = "/Users/denis/GoglandProjects/golangBook/GoRoutines/"

// InputError will be raised when the input is not right.
type InputError struct {
	// The operation that caused the error.
	Op  string
	// The error that occurred during the operation.
	Err error
}

func (e *InputError) Error() string { return "Error: " + e.Op + ": " + e.Err.Error() }

var (
	InvalidCommand   = errors.New("Invalid command.")
	TooManyArguments = errors.New("Too many arguments.")
	TooFewArguments  = errors.New("Too few arguments.")
)

// SendFile sends the file to the client and returns true if it succeeds and false otherwise.
func SendFile(c net.Conn, path string) (int64, error) {
	var fileName string

	// Make sure the user can't request any files on the system.
	lastForwardSlash := strings.LastIndex(path, "/")
	if lastForwardSlash != -1 {
		// Eliminate the last forward slash i.e ../../asdas will become asdas
		fileName = path[lastForwardSlash + 1:]
	} else {
		fileName = path
	}

	file, err := os.Open(PATH + fileName)
	if err != nil {
		// Open file failed.
		log.Println(err)
		return 0, err
	}
	defer file.Close() // Closing the fd when the function has exited.

	n, err := io.Copy(c, file)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// Executed when user sends . or .. for directory and prev directory.
	if n == 0 {
		log.Println("0 bits written for:", path)
		return 0, nil
	}

	return n, nil
}

// CheckArgumentsLength returns an error if length is not equal to expected.
func CheckArgumentsLength(length int, expected int) error {
	if length > expected {
		return TooManyArguments
	} else if length < expected {
		return TooFewArguments
	}
	return nil
}

// ListFiles list the files from path and sends them to the connection
func ListFiles(c net.Conn) error {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		return err
	}

	buffer := bytes.NewBufferString("Directory Mode Size LastModified Name\n")
	for _, f := range files {
		buffer.WriteString(strconv.FormatBool(f.IsDir()) + " " + string(f.Mode().String()) + " " +
			strconv.FormatInt(f.Size(), 10) + " " + f.ModTime().String() + " " + string(f.Name()) + " " + "\n")
	}

	_, err = c.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func processInput(c net.Conn, text string) error {
	commands := strings.Fields(text)
	commandsLen := len(commands)

	// Possibly empty input, just go on.
	if commandsLen == 0 {
		return nil
	}

	switch commands[0] {
	case "get":
		// Check arguments
		err := CheckArgumentsLength(commandsLen, 2)
		if err != nil {
			return &InputError{commands[0], err}
		}

		// Get the file
		_, err = SendFile(c, commands[1])
		if err != nil {
			return &InputError{"SendFile", err}
		}
	case "ls":
		// Check arguments
		err := CheckArgumentsLength(commandsLen, 1)
		if err != nil {
			return &InputError{commands[0], err}
		}

		err = ListFiles(c)
		if err != nil {
			return &InputError{commands[0], err}
		}
	case "clear":
		// Check arguments
		err := CheckArgumentsLength(commandsLen, 1)
		if err != nil {
			return &InputError{commands[0], err}
		}

		// Ansi clear: 1b 5b 48 1b 5b 4a
		// clear | hexdump -C
		var b []byte = []byte{0x1b, 0x5b, 0x48, 0x1b, 0x5b, 0x4a}
		c.Write(b)
	case "exit":
		err := CheckArgumentsLength(commandsLen, 1)
		if err != nil {
			return &InputError{commands[0], err}
		}

		c.Close()
	default:
		return &InputError{commands[0], InvalidCommand}
	}

	return nil
}

func handleConnection(c net.Conn) {
	defer c.Close()
	io.WriteString(c, "Hello and welcome to simple ftp\n")

	log.Println(c.RemoteAddr(), "has connected.")

	// Process input
	input := bufio.NewScanner(c)
	for input.Scan() {
		log.Println(c.RemoteAddr(), ":", input.Text())

		err := processInput(c, input.Text())
		if err != nil {
			log.Println(err)
			io.WriteString(c, err.Error()+"\n")
		}
	}

	// Client has left.
	log.Println(c.RemoteAddr(), "has disconnected.")
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConnection(conn)
	}
}
