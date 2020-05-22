package utils

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type HTMLParser struct{}

func (hp *HTMLParser) ParseHTML(htmlString string) (string, error) {
	// Get description li string array
	stringArr, err := hp.getDescLIs(htmlString)
	if err != nil {
		return "", err
	}

	// Return compiled Ul string
	result, err := hp.buildUL(stringArr)
	if err != nil {
		return "", err
	}

	return html.UnescapeString(result), err
}

func (hp *HTMLParser) getDescLIs(htmlString string) ([]string, error) {
	var result []string

	// Parse into html node
	reader := strings.NewReader(htmlString)
	doc, err := html.Parse(reader)
	if err != nil {
		return result, fmt.Errorf("fail to parse %v", htmlString)
	}

	textNodes, err := hp.getTextNode(doc)
	if err != nil {
		return result, err
	}

	for _, textNode := range textNodes {
		result = append(result, textNode.Data)
	}
	return result, err
}

func (hp HTMLParser) buildUL(descArr []string) (string, error) {
	stleAttr := html.Attribute{
		Namespace: "",
		Key:       "style",
		Val:       "font-family:Arial;font-size:14px;",
	}

	resultNode := html.Node{
		Type:      3,
		DataAtom:  183,
		Data:      "ul",
		Namespace: "",
		Attr: []html.Attribute{
			stleAttr,
		},
	}

	for _, desc := range descArr {
		li := hp.newLI(desc)
		resultNode.AppendChild(&li)
	}

	var resultBuffer bytes.Buffer
	err := html.Render(&resultBuffer, &resultNode)
	if err != nil {
		return "", err
	}

	return resultBuffer.String(), err
}

func (hp HTMLParser) newLI(stringContent string) html.Node {
	textNode := html.Node{
		Type:      1,
		DataAtom:  0,
		Data:      stringContent,
		Namespace: "",
		Attr:      []html.Attribute{},
	}
	result := html.Node{
		Type:      3,
		DataAtom:  68,
		Data:      "li",
		Namespace: "",
		Attr:      []html.Attribute{},
	}
	result.AppendChild(&textNode)
	return result
}

func (hp *HTMLParser) getTextNode(doc *html.Node) ([]html.Node, error) {
	var result []html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode && len(strings.ReplaceAll(strings.TrimSpace(node.Data), "\n", "")) > 0 {
			result = append(result, *node)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if result != nil {
		return result, nil
	}
	return nil, fmt.Errorf("missing text tags in the node tree; %v", doc)
}