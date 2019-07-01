package figma

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUnmarshalFile(t *testing.T) {
	jsonFile, err := os.Open("testdata/file.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		t.Errorf("test failed with unexpected reason")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("test failed with unexpected reason")
	}

	f, err := UnmarshalFile(byteValue)

	if f.Name != "App Colors" {
		t.Errorf("test failed")
	}

	d := f.Document

	if d.Children[0].Name != "Page 1" {
		t.Errorf("test failed")
	}
}

func TestUnmarshalNode(t *testing.T) {
	jsonFile, err := os.Open("testdata/filenode.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		t.Errorf("test failed with unexpected reason")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("test failed with unexpected reason")
	}

	f, err := UnmarshalFileNodes(byteValue)

	if f.Name != "App Colors" {
		t.Errorf("test failed")
	}

	if f.Nodes["1:2"].Styles["2:2"].Key != "12345678910" {
		t.Errorf("invalid key value")
	}
}
