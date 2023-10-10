package testing

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotification(t *testing.T) {
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	var notifiedUser, notifiedMsg string
	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}
	const user = "joe@example.com"
	checkQuota(user)

	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+"want substring %q", notifiedMsg, wantSubstring)
	}
}
