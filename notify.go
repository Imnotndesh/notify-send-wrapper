package notify_send_wrapper

import (
	"os/exec"
	"strconv"
)

// Notification represents a notification object.
type Notification struct {
	Summary    string
	Body       string
	Urgency    UrgencyLevel
	ExpireTime int
	Icon       string
	AppName    string
	Category   NotificationCategory
	Hints      map[string]string
}

// UrgencyLevel represents the urgency level of a notification to be used in the call.
type UrgencyLevel string

// Urgency levels.
const (
	UrgencyLow      UrgencyLevel = "low"
	UrgencyNormal   UrgencyLevel = "normal"
	UrgencyCritical UrgencyLevel = "critical"
)

// NotificationCategory represents a notification category.
type NotificationCategory string

// Predefined notification categories.
const (
	CategoryDevice    NotificationCategory = "device"
	CategoryEmail     NotificationCategory = "email"
	CategoryIM        NotificationCategory = "im"
	CategoryNetwork   NotificationCategory = "network"
	CategoryPresence  NotificationCategory = "presence"
	CategorySystem    NotificationCategory = "system"
	CategoryTransport NotificationCategory = "transport"
)

// NotificationHints represent a list of hints
type NotificationHints string

const (
	HintActionIcons           NotificationHints = "action-icons"
	HintCategory              NotificationHints = "category"
	HintDesktopEntry          NotificationHints = "desktop-entry"
	HintImagePath             NotificationHints = "image-path"
	HintResident              NotificationHints = "resident"
	HintSoundName             NotificationHints = "sound-name"
	HintSoundFile             NotificationHints = "sound-file"
	HintSuppressSound         NotificationHints = "suppress-sound"
	HintTransient             NotificationHints = "transient"
	HintXCanonicalSynchronous NotificationHints = "x-canonical-private-synchronous"
)

// NewNotification creates a new Notification object with default values.
func NewNotification() *Notification {
	return &Notification{
		Urgency: UrgencyNormal,
	}
}

// SetSummary sets the summary (title) of the notification.
func (n *Notification) SetSummary(summary string) *Notification {
	n.Summary = summary
	return n
}

// SetBody sets the body (message) of the notification.
func (n *Notification) SetBody(body string) *Notification {
	n.Body = body
	return n
}

// SetUrgency sets the urgency level of the notification.
func (n *Notification) SetUrgency(urgency UrgencyLevel) *Notification {
	n.Urgency = urgency
	return n
}

// SetExpireTime sets the expire time (in milliseconds) of the notification.
func (n *Notification) SetExpireTime(expireTime int) *Notification {
	n.ExpireTime = expireTime
	return n
}

// SetIcon sets the icon of the notification.
func (n *Notification) SetIcon(icon string) *Notification {
	n.Icon = icon
	return n
}

// SetAppName sets the application name of the notification.
func (n *Notification) SetAppName(appName string) *Notification {
	n.AppName = appName
	return n
}

// SetCategory sets the category of the notification.
func (n *Notification) SetCategory(category NotificationCategory) *Notification {
	n.Category = category
	return n
}

// AddHint sets a hint.
func (n *Notification) AddHint(key NotificationHints, value string) *Notification {
	if n.Hints == nil {
		n.Hints = make(map[string]string)
	}
	n.Hints[string(key)] = value
	return n
}

// AddCustomHint sets a custom hint.
func (n *Notification) AddCustomHint(key, value string) *Notification {
	if n.Hints == nil {
		n.Hints = make(map[string]string)
	}
	n.Hints[key] = value
	return n
}

// Send sends the notification using notify-send.
func (n *Notification) Send() error {
	args := []string{n.Summary, n.Body}

	if n.Urgency != "" {
		args = append(args, "--urgency="+string(n.Urgency))
	}

	if n.ExpireTime > 0 {
		args = append(args, "--expire-time="+strconv.Itoa(n.ExpireTime))
	}

	if n.Icon != "" {
		args = append(args, "--icon="+n.Icon)
	}

	if n.AppName != "" {
		args = append(args, "--app-name="+n.AppName)
	}

	if n.Category != "" {
		args = append(args, "--category="+string(n.Category))
	}

	for key, value := range n.Hints {
		args = append(args, "--hint="+key+"="+value)
	}

	cmd := exec.Command("notify-send", args...)
	return cmd.Run()
}
