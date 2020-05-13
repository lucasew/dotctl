package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func (r DotctlRepo) Bring(extra ...string) error {
	allRules := map[string]interface{}{}
	{
		{
			rules := r.RulesRepository().GetRules()
			for _, rule := range rules {
				allRules[rule] = nil
			}
		}
		for _, rule := range extra {
			allRules[rule] = nil
		}
	}
	for rule := range allRules {
		log.Printf("Processing rule '%s'...", rule)
		err := copy_thing(rule, r.basepath)
		if err != nil {
			return err
		}
	}
	return nil
}

func copy_thing(from, baseto string) error {
	item, err := homedir.Expand(from)
	if err != nil {
		return err
	}
	fromplusbase := strings.Replace(filepath.Join(baseto, from), "~", "HOME", 1)
	stat, err := os.Stat(item)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		log.Printf("Copying folder %s...", item)
		err = os.MkdirAll(fromplusbase, os.ModePerm)
		if err != nil {
			return err
		}
		isFirst := true
		return filepath.Walk(item, func(path string, info os.FileInfo, err error) error {
			if isFirst {
				isFirst = false
				return nil
			}
			return copy_thing(path, baseto)
		})
	} else {
		log.Printf("Copying file '%s' to '%s'...", item, fromplusbase)
		err = os.MkdirAll(filepath.Dir(fromplusbase), os.ModePerm)
		if err != nil {
			return err
		}
		fdest, err := os.Create(fromplusbase)
		if err != nil {
			return err
		}
		defer fdest.Close()
		forigin, err := os.Open(item)
		if err != nil {
			return err
		}
		defer forigin.Close()
		_, err = io.Copy(fdest, forigin)
		if err != nil {
			return err
		}
	}
	return nil
}
