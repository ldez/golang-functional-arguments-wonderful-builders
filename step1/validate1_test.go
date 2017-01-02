package step1

import "testing"


func TestValidateLinuxIsSupported(t *testing.T) {
	valid := Validate(ANode().Platform(&Platform{
		Architecture: "x86_64",
		OS:           "linux",
	}).Build())
	if !valid {
		t.Fatal("linux should be supported, it was not")
	}
}

func TestValidateDarwinIsNotSupported(t *testing.T) {
	valid := Validate(ANode().Platform(&Platform{
		Architecture: "x86_64",
		OS:           "darwin",
	}).Build())
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
