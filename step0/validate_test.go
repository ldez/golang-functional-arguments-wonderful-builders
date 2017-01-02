package step0

import "testing"


func TestValidateLinuxIsSupported(t *testing.T) {
	valid := Validate(&Node{
		Name: "carthage",
		Hostname: "carthage.sbr.pm",
		Platform: &Platform{
			Architecture: "x86_64",
			OS:           "linux",
		},
	})
	if !valid {
		t.Fatal("linux should be supported, it was not")
	}
}

func TestValidateDarwinIsNotSupported(t *testing.T) {
	valid := Validate(&Node{
		Name: "babylon",
		Hostname: "babylon.sbr.pm",
		Platform: &Platform{
			Architecture: "x86_64",
			OS:           "darwin",
		},
	})
	if valid {
		t.Fatal("darwin should not be supported, it was")
	}
}

// --

type Platform struct {
	Architecture, OS string
}

type Node struct {
	Name, Hostname string
	Platform *Platform
}

func Validate(node *Node) bool {
	return node.Platform.OS == "linux"
}
