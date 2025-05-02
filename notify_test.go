package notify_send_wrapper

import (
	"testing"
)

// --- Test Variables ---
const (
	testSummary         = "Test Notification Title"
	testBody            = "This is the test notification body."
	testIcon            = "/home/brian/Desktop/test_stuff/check.svg"
	testAppName         = "Test Application"
	testCategory        = "test.category"
	testHintKey         = "test-hint-key"
	testHintValue       = "test-hint-value"
	testCustomHintKey   = "custom-test-key"
	testCustomHintValue = "custom-test-value"
	testSoundName       = "ting"
	testSoundFile       = "/home/brian/Desktop/test_stuff/ting.ogg"
	testImagePath       = "/home/brian/Desktop/test_stuff/check.svg"
)

const (
	testExpireTime = 3000
	testReplaceID  = 123
)

var (
	testUrgency = UrgencyCritical
)

func TestSimpleUsability(t *testing.T) {
	err := NewNotification().SetSummary(testSummary).SetBody(testBody).SetIcon(testIcon).Send()
	if err != nil {
		t.Errorf("Error running simple test: %v", err)
	}
}
func TestNotification_Setters(t *testing.T) {
	n := NewNotification()

	n.SetSummary(testSummary).
		SetBody(testBody).
		SetUrgency(testUrgency).
		SetExpireTime(testExpireTime).
		SetIcon(testIcon).
		SetAppName(testAppName).
		SetCategory(testCategory).
		AddHint(HintXCanonicalSynchronous, "true").
		AddCustomHint(testCustomHintKey, testCustomHintValue).
		SetReplaceID(testReplaceID).
		SetSoundName(testSoundName).
		SetSoundFile(testSoundFile).
		SetSuppressSound(true).
		SetResident(true).
		SetTransient(true).
		SetImagePath(testImagePath).
		SetActionIcons(true)

	if n.Summary != testSummary {
		t.Errorf("SetSummary failed: got %q, want %q", n.Summary, testSummary)
	}
	if n.Body != testBody {
		t.Errorf("SetBody failed: got %q, want %q", n.Body, testBody)
	}
	if n.Urgency != testUrgency {
		t.Errorf("SetUrgency failed: got %q, want %q", n.Urgency, testUrgency)
	}
	if n.ExpireTime != testExpireTime {
		t.Errorf("SetExpireTime failed: got %d, want %d", n.ExpireTime, testExpireTime)
	}
	if n.Icon != testIcon {
		t.Errorf("SetIcon failed: got %q, want %q", n.Icon, testIcon)
	}
	if n.AppName != testAppName {
		t.Errorf("SetAppName failed: got %q, want %q", n.AppName, testAppName)
	}
	if n.Category != testCategory {
		t.Errorf("SetCategory failed: got %q, want %q", n.Category, NotificationCategory(testCategory))
	}
	if n.Hints["x-canonical-private-synchronous"] != "true" {
		t.Errorf("AddHint failed: got %q, want %q", n.Hints["x-canonical-private-synchronous"], "true")
	}
	if n.Hints[testCustomHintKey] != testCustomHintValue {
		t.Errorf("AddCustomHint failed: got %q, want %q", n.Hints[testCustomHintKey], testCustomHintValue)
	}
	if n.ReplaceID != testReplaceID {
		t.Errorf("SetReplaceID failed: got %d, want %d", n.ReplaceID, testReplaceID)
	}
	if n.SoundName != testSoundName {
		t.Errorf("SetSoundName failed: got %q, want %q", n.SoundName, testSoundName)
	}
	if n.SoundFile != testSoundFile {
		t.Errorf("SetSoundFile failed: got %q, want %q", n.SoundFile, testSoundFile)
	}
}
