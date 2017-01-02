package step3

import (
	"testing"
)

func TestValidateLinuxIsSupported(t *testing.T) {
	valid := Validate(ANode(WithPlatform(Linux)))
	if !valid {
		t.Fatal("linux should be supported, it was not")
	}
}

func TestValidateDarwinIsNotSupported(t *testing.T) {
	valid := Validate(ANode(WithPlatform(Darwin)))
	if valid {
		t.Fatal("darwin should not be supported, it was")
	}
}

// Function builders
func WithPlatform(builders ...func(*Platform)) func(n *Node) {
	return func(n *Node) {
		n.Platform = APlatform(builders...)
	}
}

func Linux(p *Platform) {
	p.OS = "linux"
}

func Darwin(p *Platform) {
	p.OS = "darwin"
}

// --

func ANode(nodeBuilders ...func(*Node)) *Node {
	node := &Node{
		Name: "node",
		// Other defaults
	}

	for _, build := range nodeBuilders {
		build(node)
	}

	return node
}

func APlatform(platformBuilders ...func(*Platform)) *Platform {
	platform := &Platform{
		Architecture: "x64_86",
		OS:           "linux",
	}

	for _, build := range platformBuilders {
		build(platform)
	}

	return platform
}

// --

type Platform struct {
	Architecture, OS string
}

type Node struct {
	Name, Hostname string
	Platform       *Platform
}

func Validate(node *Node) bool {
	return node.Platform.OS == "linux"
}

// -- examples

//func toto() {
//	// a default node
//	node1 := ANode()
//
//	// a node with a specific Hostname
//	node2 := ANode(func(n *Node) {
//		n.Hostname = "custom-hostname"
//	})
//	// a node with a specific name and platform
//	node3 := ANode(func(n *Node) {
//		n.Name = "custom-name"
//	}, func(n *Node) {
//		n.Platform = APlatform(func(p *Platform) {
//			p.OS = "darwin"
//		})
//	})
//}
