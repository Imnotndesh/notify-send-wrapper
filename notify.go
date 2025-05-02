package notify_send_wrapper

import (
	"os/exec"
	"strconv"
)

// Notification represents a notification object.
type Notification struct {
	ReplaceID     int
	SoundName     string
	SoundFile     string
	SuppressSound bool
	Resident      bool
	Transient     bool
	ImagePath     string
	ActionIcons   bool
	DesktopEntry  string
	Summary       string
	Body          string
	Urgency       UrgencyLevel
	ExpireTime    int
	Icon          string
	AppName       string
	Category      NotificationCategory
	Hints         map[string]string
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

// SetIcon sets the icon of the notification (use absolute path).
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
func (n *Notification) SetReplaceID(id int) *Notification {
	n.ReplaceID = id
	return n
}

// SetSoundName sets the sound name of the notification.
func (n *Notification) SetSoundName(name string) *Notification {
	n.SoundName = name
	return n
}

// SetSoundFile sets the sound file of the notification.
func (n *Notification) SetSoundFile(file string) *Notification {
	n.SoundFile = file
	return n
}

// SetSuppressSound sets if sound is suppressed
func (n *Notification) SetSuppressSound(suppress bool) *Notification {
	n.SuppressSound = suppress
	return n
}

// SetResident sets if the notification is resident
func (n *Notification) SetResident(resident bool) *Notification {
	n.Resident = resident
	return n
}

// SetTransient sets if the notification is transient
func (n *Notification) SetTransient(transient bool) *Notification {
	n.Transient = transient
	return n
}

// SetImagePath sets the image path of the notification.
func (n *Notification) SetImagePath(path string) *Notification {
	n.ImagePath = path
	return n
}

// SetActionIcons sets if action icons are enabled.
func (n *Notification) SetActionIcons(actionIcons bool) *Notification {
	n.ActionIcons = actionIcons
	return n
}

// SetDesktopEntry sets the desktop entry of the notification.
func (n *Notification) SetDesktopEntry(entry string) *Notification {
	n.DesktopEntry = entry
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

	if n.ReplaceID != 0 {
		args = append(args, "--replace-id="+strconv.Itoa(n.ReplaceID))
	}

	if n.SoundName != "" {
		args = append(args, "--sound="+n.SoundName)
	} else if n.SoundFile != "" {
		args = append(args, "--sound-file="+n.SoundFile)
	}

	if n.SuppressSound {
		args = append(args, "--suppress-sound")
	}

	if n.Resident {
		args = append(args, "--resident")
	}

	if n.Transient {
		args = append(args, "--transient")
	}

	if n.ImagePath != "" {
		args = append(args, "--image-path="+n.ImagePath)
	}

	if n.ActionIcons {
		args = append(args, "--action-icons")
	}

	if n.DesktopEntry != "" {
		args = append(args, "--desktop-entry="+n.DesktopEntry)
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
