package main

import (
	"encoding/json"
	"fmt"
	. "github.com/jmervine/GoT"
	"io/ioutil"
	"testing"
)

func rawJson() (raw []byte) {
	raw, err := ioutil.ReadFile("_support/push.json")
	if err != nil {
		panic(err)
	}
	return
}

func unmarshalJson() (push PushEvent) {
	if err := json.Unmarshal(rawJson(), &push); err != nil {
		panic(err)
	}
	return
}

func TestPushEvent_Branch(T *testing.T) {
	push := unmarshalJson()

	Go(T).AssertEqual(push.Branch(), "gh-pages")
}

func TestPushEvent_ByOwner(T *testing.T) {
	push := unmarshalJson()

	Go(T).Assert(push.ByOwner())
}

func TestPushEvent_ByAdmin(T *testing.T) {
	push := unmarshalJson()

	Go(T).Assert(push.ByAdmin())
}

func TestPushEvent_ToMaster(T *testing.T) {
	push := unmarshalJson()

	Go(T).Refute(push.ToMaster())
}

func Example() {
	data := rawJson()

	var push PushEvent
	json.Unmarshal(data, &push)

	branch := push.Branch()
	sender := push.Sender.Login

	fmt.Printf("[INFO] %s was updated by %s", branch, sender)
	if push.ByOwner() {
		fmt.Printf(" -- the owner")
	}

	if push.ByAdmin() {
		fmt.Printf(" and admin")
	}

	if push.ToMaster() {
		fmt.Printf(" <-- master")
	}

	fmt.Println(".")

	// Output:
	// [INFO] gh-pages was updated by jmervine -- the owner and admin.
}
