package tree

import (
	"gopkg.in/yaml.v3"
	"os"
)

func SaveInYaml(node *Node) error {
	yamlData, err := yaml.Marshal(&node)

	if err != nil {
		return err
	}

	err = os.WriteFile("monFichier.yaml", yamlData, 0644)

	if err != nil {
		return err
	}

	return nil
}

func LoadFromYaml() (Node, error) {
	data, err := os.ReadFile("monFichier.yaml")

	if err != nil {
		return NewEmptyNode(), err
	}

	var node Node
	err = yaml.Unmarshal(data, &node)

	if err != nil {
		return NewEmptyNode(), err
	}

	return node, nil
}
