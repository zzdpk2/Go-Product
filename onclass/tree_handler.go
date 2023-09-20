package main

import "strings"

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	handler HandlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	// TODO implement me
	panic("implement me")
}

func (h HandlerBasedOnTree) Route(method string, pattern string, handlerFunc HandlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		matchChild, ok := h.findMatchedChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}

}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFn HandlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFn
}

func (h *HandlerBasedOnTree) findMatchedChild(root *node, path string) (*node, bool) {
	var wildCardNode *node
	for _, child := range root.children {
		if child.path == path && child.path == "*" {
			return child, true
		}

		if child.path == "*" {
			wildCardNode = child
		}
	}
	return wildCardNode, wildCardNode != nil
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 8),
	}
}
