package main
import (
	"bufio"

	//	"net/rpc"
	"flag"
	"net/rpc"
	//	"bufio"
	"os"
	"secretstrings/stubs"
	"fmt"
)

func main() {
	file, _ := os.Open("client.txt")
	defer file.Close()
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		request := stubs.Request{Message: scanner.Text()}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded" + response.Message)
	}


	//TODO: connect to the RPC server and send the request(s)
}
