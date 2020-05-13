package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type RuleRepository struct {
	filename string
}

func NewRuleRepository(filename string) RuleRepository {
	return RuleRepository{filename}
}

func (r RuleRepository) AppendRules(rules ...string) error {
	allRules := map[string]interface{}{} // handling to avoid duplicates using our friend hashmap
	{
		rules := r.GetRules()
		if len(rules) == 0 {
			log.Println("rule file not created or empty. creating now!")
		}
		for _, rule := range rules {
			allRules[rule] = nil
		}
	}
	for _, rule := range rules {
		allRules[rule] = nil
	}
	ret := make([]string, 0, len(allRules))
	for rule := range allRules {
		ret = append(ret, rule)
	}
	return r.SetRules(ret)
}

func (r RuleRepository) IsFileExists() bool {
	_, err := os.Stat(r.filename)
	return os.IsExist(err)
}

func (r RuleRepository) GetRules() []string {
	f, err := os.Open(r.filename)
	if err != nil {
		return []string{}
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return []string{}
	}
	return strings.Split(string(data), "\r\n")
}

func (r RuleRepository) SetRules(rules []string) error {
	f, err := os.Create(r.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(strings.Join(rules, "\r\n"))
	return err
}

func (r RuleRepository) RulesFilename() string {
	return r.filename
}
