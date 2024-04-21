package service

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func parse(_html string) *html.Node {
	doc, err := html.Parse(strings.NewReader(_html))
	if err != nil {
		panic(err)
	}
	return doc
}

func traverse(n *html.Node, tag string, key string, val string) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// fmt.Printf("DataAtom[%v] Data[%v] Attr[%v] Tag[%v] Type[%v]\n", c.DataAtom, c.Data, c.Attr, c.Parent.Data, c.Type)
		if c.Type == html.ElementNode && c.Data == tag && ((key != "" && contains(c.Attr, key, val)) || key == "") {
			return c
		}

		if res := traverse(c, tag, key, val); res != nil {
			return res
		}
	}
	return nil
}

func text(n *html.Node) (string, bool) {
	if n.Type == html.TextNode {
		return strings.TrimSpace(html.UnescapeString(n.Data)), true
	}
	n = n.FirstChild
	if n != nil && n.Type == html.TextNode {
		return strings.TrimSpace(html.UnescapeString(n.Data)), true
	}
	return "", false
}
func attr(s []html.Attribute, key string) (string, bool) {
	for _, a := range s {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

func contains(s []html.Attribute, key string, val string) bool {
	for _, a := range s {
		// fmt.Printf("ATTR [%v|%v] [%v|%v]\n", a.Key, key, a.Val, val)
		if a.Key == key {
			if val == "" || a.Val == val {
				return true
			} else {
				matched, _ := regexp.Match(`\b`+strings.ReplaceAll(val, "-", "_")+`\b`, s2b(strings.ReplaceAll(a.Val, "-", "_")))
				return matched
			}
		}
	}
	return false
}
