/*The Notifier interface defines a single method, Notify(message string), which abstracts the action of sending notifications. This approach focuses solely on the behavior that all notifiers must have: the ability to send a notification message. This abstraction leaves the specifics of how the program delivers the notification to the implementing types. Run the code to see how this implementation handles different types. */

package main

import "fmt"

func main() {
	var notifier Notifier

	notifier = EmailNotifier{}
	notifier.Notify("Hello via Email!")

	notifier = SMSNotifier{}
	notifier.Notify("Hello via SMS!")
}

// Define behavior-focused and abstract interface
type Notifier interface {
	Notify(message string)
}

// Implement the interface in different ways
type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending email with message:", message)
}

type SMSNotifier struct{}

func (s SMSNotifier) Notify(message string) {
	fmt.Println("Sending SMS with message:", message)
}
