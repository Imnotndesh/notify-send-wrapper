# Notify-Send Wrapper
A wrapper for the [notify-send](https://man.archlinux.org/man/notify-send.1.en) that allows programmatic structuring of notification messages from Golang applications.
## Prerequisites
- Notify-send is installed in your distro (usually pre-installed)
## Usage
### In project
Imported by appending the line below in your project file:
```go
  import "github.com/Imnotndesh/notify-send-wrapper"
```
Can also be fetched directly from terminal using:
```shell
  go get "github.com/Imnotndesh/notify-send-wrapper"
```
## Usage Examples

Here are a few examples of how to use the `notify-send-wrapper` package in your Go applications.

### Basic Notification

Send a simple notification with a summary and body:

```go
func main() {
	notification := notify_send_wrapper.NewNotification().
		SetSummary("Hello, User!").
		SetBody("This is a basic notification from my Go app.")

	err := notification.Send()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Notification sent successfully!")
}
```
### Notification with Urgency
Send a notification with a specific urgency level:
```go
func main() {
	notification := notify_send_wrapper.NewNotification().
		SetSummary("Important Update!").
		SetBody("Please update your system to the latest version.").
		SetUrgency(notify_send_wrapper.UrgencyCritical)

	err := notification.Send()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Critical notification sent!")
}
```
### Notification with Icon and Expiration
Send a notification with a custom icon and set an expiration time:
```go
func main() {
	notification := notify_send_wrapper.NewNotification().
		SetSummary("New Message").
		SetBody("You have one new message in your inbox.").
		SetIcon("/path/to/your/icon.png"). // Use an absolute path to your icon file
		SetExpireTime(int(5 * time.Second / time.Millisecond)) // Expire after 5 seconds

	err := notification.Send()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Notification with icon and expiration sent!")
}
```
### Notification with Hints
You can add hints to customize the notification behavior:
```go
func main() {
	notification := notify_send_wrapper.NewNotification().
		SetSummary("Low Battery").
		SetBody("Your battery is running low. Please connect to a power source.").
        AddHint(notify_send_wrapper.HintResident, "true") // Make the notification persistent

	err := notification.Send()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Notification with resident hint sent!")
}
```

