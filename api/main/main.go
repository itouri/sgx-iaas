package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"k8s.io/apimachinery/pkg/util/yaml"

	"k8s.io/client-go/kubernetes/scheme"
)

const y = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: baz
  namespace: bat
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: foo
  namespace: bar
`

func main() {
	b := bufio.NewReader(strings.NewReader(y))
	r := yaml.NewYAMLReader(b)

	for {
		doc, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		d := scheme.Codecs.UniversalDeserializer()
		obj, _, err := d.Decode(doc, nil, nil)
		if err != nil {
			log.Fatalf("could not decode yaml: %s\n%s", y, err)
		}
		fmt.Println(obj)
	}
}
