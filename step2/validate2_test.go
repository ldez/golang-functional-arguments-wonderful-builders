package step2

import "testing"

func TestValidateLinuxIsSupported(t *testing.T) {
	valid := Validate(ANode().Platform(APlatform().OS("linux").Build()).Build())
	if !valid {
		t.Fatal("linux should be supported, it was not")
	}
}

func TestValidateDarwinIsNotSupported(t *testing.T) {
	ANode().Platform(APlatform().Build())
	valid := Validate(ANode().Platform(APlatform().OS("darwin").Build()).Build())
	if valid {
		t.Fatal("darwin should not be supported, it was")
	}
}

// --

func ANode() *NodeBuilder {
	return &NodeBuilder{
		node: &Node{
			Name: "node",
			// Other defaults
		},
	}
}

type NodeBuilder struct {
	node *Node
}

func (b *NodeBuilder) Build() *Node {
	return b.node
}

func (b *NodeBuilder) Hostname(hostname string) *NodeBuilder {
	b.node.Hostname = hostname
	return b
}

func (b *NodeBuilder) Name(name string) *NodeBuilder {
	b.node.Name = name
	return b
}

func (b *NodeBuilder) Platform(platform *Platform) *NodeBuilder {
	b.node.Platform = platform
	return b
}

// --

func APlatform() *PlatformBuilder {
	return &PlatformBuilder{
		platform: &Platform{
			Architecture: "x64_86",
			OS:           "linux",
		},
	}
}

type PlatformBuilder struct {
	platform *Platform
}

func (b *PlatformBuilder) Build() *Platform {
	return b.platform
}

func (b *PlatformBuilder) OS(os string) *PlatformBuilder {
	b.platform.OS = os
	return b
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
