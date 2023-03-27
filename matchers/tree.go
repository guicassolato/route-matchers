package matchers

import (
	"fmt"
	"strings"

	"github.com/guicassolato/route-matchers/utils"
)

func NewTree(options ...StoreOption) Store {
	t := &tree{}
	for _, option := range options {
		option(t)
	}
	return t
}

type tree struct {
	roots   []*node
	Matcher
}

func (t *tree) String() string {
	s := ""
	for _, root := range t.roots {
		s = fmt.Sprintf("%s%s\n", s, root.print(0))
	}
	return s
}

// impl:Store

func (t *tree) Add(value string) {
	newNode := &node{value: value}

	for _, root := range t.roots {
		if newNode.value == root.value {
		  return
		}
		if insertNode := root.find(t.nodeContainsValueFunc(value)); insertNode != nil {
		  if insertNode.value != value {
			  reassign, keep := utils.Partition(insertNode.children, t.valueContainsNodeFunc(value))
				for _, node := range reassign {
					newNode.append(node)
				}
				insertNode.children = keep
				insertNode.append(newNode)
			}
			return
		}
	}

	reassign, keep := utils.Partition(t.roots, t.valueContainsNodeFunc(value))
	for _, node := range reassign {
		newNode.append(node)
	}
	t.roots = append(keep, newNode)
}

func (t *tree) List() []string {
	var list []string
	depth := 0
	for _, root := range t.roots {
		list = append(list, root.list(depth)...)
	}
	return list
}

func (t *tree) Size() int {
	return len(t.List())
}

func (t *tree) SetMatcher(matcher Matcher) {
	t.Matcher = matcher
}

func (t *tree) matcher() Matcher {
	if t.Matcher != nil {
		return t.Matcher
	}
	return DefaultMatcher
}

func (t *tree) nodeContainsValueFunc(value string) func(*node) bool {
	return func(node *node) bool {
		return t.matcher().Contains(node.value, value)
	}
}

func (t *tree) valueContainsNodeFunc(value string) func(*node) bool {
	return func(node *node) bool {
		return t.matcher().Contains(value, node.value)
	}
}

type node struct {
	value    string
	parent   *node
	children []*node
}

func (n *node) String() string {
	return n.value
}

func (n *node) append(child *node) {
	child.parent = n
	n.children = append(n.children, child)
}

func (n *node) list(depth int) []string {
	list := []string{ n.print(depth) }
	for _, child := range n.children {
		list = append(list, child.list(depth+1)...)
	}
	return list
}

func (n *node) find(f func(*node) bool) *node {
	for _, child := range n.children {
		if node := child.find(f); node != nil {
			return node
		}
	}
	if f(n) {
		return n
	}
	return nil
}

func (n *node) print(depth int) string {
  return fmt.Sprintf("%s﹂%s", strings.Repeat(" ", depth*2), n)
}
