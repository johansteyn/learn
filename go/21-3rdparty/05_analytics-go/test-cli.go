package main

import (
	"os"

	"github.com/segmentio/analytics-go"
	"github.com/segmentio/ksuid"
)

func main() {
	flags := make([]string, 0, 2)
	flags = append(flags, "limit", "page")

	id := ksuid.New()
	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	properties := analytics.NewProperties()
	properties.Set("installationId", "a1b2c3d4e5f6g7h8i9j0k1l2m3n")
	properties.Set("orgId", "1234567890abcdefghijklmn")
	properties.Set("projectId", "0987654321zyxwvutsrqponm")
	properties.Set("authentication", "OAuth")
	properties.Set("default", "true")
	//properties.Set("result", "SUCCESS")
	properties.Set("result", "ERROR")
  properties.Set("error", `Error: invalid argument "x" for "--limit" flag: strconv.ParseInt: parsing "x": invalid syntax`)
	properties.Set("duration", 123)
	properties.Set("flags", flags)
	properties.Set("alias", "ls")
	properties.Set("version", "1.23.0")
	properties.Set("os", "linux/amd64")
	properties.Set("installer", "homebrew")
	properties.Set("terminal", "teletype")
	client.Enqueue(analytics.Track{
		Event: "atlas-projects-list",
		AnonymousId: id.String(),
		UserId: "123abc456def789ghi987jkl654mno321pqr123stu456vwx",
		Properties: properties,
	})
	client.Close()
}

