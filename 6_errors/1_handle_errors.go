package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func waitForServer(url string) error {

    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)

    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)

        if err == nil {
            return nil
        }
        log.Printf("Server is not responding: %v", err)
        time.Sleep(time.Second << uint(tries))
    }

    return fmt.Errorf("Server is not responding: %v", url)
}

func main() {

    if err := waitForServer("http://not-existing-domain.com"); err != nil {
        log.Fatalf("Site is not working: %v", err)
    }
}
