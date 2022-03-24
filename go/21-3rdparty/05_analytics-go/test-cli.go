package main

import (
	"os"

	"github.com/segmentio/analytics-go"
	"github.com/segmentio/ksuid"
)

func main() {
	id := ksuid.New()
	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	properties := analytics.NewProperties()
	// Leave out the command property, as it may contain PII...
	//properties.Set("command", "mongocli iam projects ls --limit 10")
	//properties.Set("command", "mongocli iam projects ls --limit 10 --page 1")
	properties.Set("installationId", "a1b2c3d4e5f6g7h8i9j0k1l2m3n")
	properties.Set("orgId", "1234567890abcdefghijklmn")
	properties.Set("projectId", "0987654321zyxwvutsrqponm")
	properties.Set("default", "true")
	properties.Set("result", "success")
//	properties.Set("result", "error")
//  properties.Set("error", `Error: invalid argument "x" for "--limit" flag: strconv.ParseInt: parsing "x": invalid syntax`)
	properties.Set("duration", "PT0.123S")
	properties.Set("flags", "--limit,--page")
	properties.Set("alias", "ls")
	properties.Set("version", "mongocli/v1.23.0-2-g89c1fa27")
	properties.Set("os", "linux/amd64")
	properties.Set("installer", "homebrew")
	properties.Set("terminal", "teletype")
	properties.Set("authentication", "OAuth")
	client.Enqueue(analytics.Track{
		Event: "mongocli-iam-projects-list",
		AnonymousId: id.String(),
		Properties: properties,
	})
	client.Close()
}

