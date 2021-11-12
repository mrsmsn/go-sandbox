package model

type (
	CarContainer struct {
		Cars []Car `yaml:"cars"`
	}

	Car struct {
		Name     string `yaml:"name"`
		JName    string `yaml:"jname"`
		Category string `yaml:"category"`
		Color    string `yaml:"color"`
	}
)
