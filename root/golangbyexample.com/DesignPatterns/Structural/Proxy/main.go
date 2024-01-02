package main

import "fmt"

func main() {
	nginxProxyServer := newNginxProxyServer()

	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpcode, body := nginxProxyServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s \nHttpCode: %d\nBody: %s\n", appStatusURL, httpcode, body)

	httpcode, body = nginxProxyServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s \nHttpCode: %d\nBody: %s\n", appStatusURL, httpcode, body)

	httpcode, body = nginxProxyServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s \nHttpCode: %d\nBody: %s\n", appStatusURL, httpcode, body)

	httpcode, body = nginxProxyServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s \nHttpCode: %d\nBody: %s\n", createUserURL, httpcode, body)

	httpcode, body = nginxProxyServer.handleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s \nHttpCode: %d\nBody: %s\n", createUserURL, httpcode, body)
}
