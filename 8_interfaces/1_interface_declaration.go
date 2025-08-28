package main

import "fmt"

type Notifier interface {
    Send(message string) StatusResponse
}

type PushNotification struct {
    DeviceId string
}

func (n *PushNotification) Send(message string) StatusResponse  {

    // BUSINESS LOGIC ...
    fmt.Println("Sending push notification with content", message, "for device with id", n.DeviceId)

    return StatusResponse{Success: true}
}

type EmailNotification struct {
    Email string
}

func (e *EmailNotification) Send(message string) StatusResponse {

    // BUSINESS LOGIC ...
    fmt.Println("Sending email message with content", message, "for address", e.Email)

    return StatusResponse{Success: true}
}

type StatusResponse struct {
    Success bool
}

func main() {

    // 1 - RUN WITH SPECIFIED STRUCT
    // CHECK IF INTERFACE IMPLEMENTED (INTERFACE METHOD SIGNATURE MATCH WITH STRUCT METHOD)
    var _ Notifier = (*PushNotification)(nil)

    pushNotification := PushNotification{DeviceId: "123456"}
    pushStatus := pushNotification.Send("...")
    fmt.Println(pushStatus.Success)

    var _ Notifier = (*EmailNotification)(nil)

    emailNotification := EmailNotification{Email: "test@mail.com"}
    emailStatus := emailNotification.Send("...")
    fmt.Println(emailStatus.Success)

    // 2 - RUN WITH SPECIFIED INTERFACE
    // n will run with each type fulfilling the contract
    // allows to write polymorphic flexible code
    var n Notifier

    n = &PushNotification{DeviceId: "123456"}
    fmt.Println(n.Send("..."))
}
