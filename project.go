package project

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type project struct {
	Name      string
	Priority  string
	Category  string
	CreatedAt time.Time
}

type ProjectList []project

func (l *ProjectList) Add(name string, priority string, category string) {
	p := project{
		Name:      name,
		Priority:  priority,
		Category:  category,
		CreatedAt: time.Now(),
	}

	*l = append(*l, p)
}

func (l *ProjectList) Save(fileName string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, json, 0644)
}

func (l *ProjectList) Get(fileName string) error {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
