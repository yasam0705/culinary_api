package pkg

import "time"

func TimeToStrRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}
