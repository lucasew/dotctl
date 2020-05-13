package main

import (
    "log"
    "os"
    "io"
    "path/filepath"
    "path"

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
        item, err := homedir.Expand(rule)
        if err != nil {
            return err
        }
        copy_thing(item, r.basepath)
    }
    return nil
}

func copy_thing(from, baseto string) error {
    stat, err := os.Stat(from)
    if err != nil {
        return err
    }
    if stat.IsDir() {
        log.Printf("Copying folder %s...", from)
        err = os.MkdirAll(filepath.Join(baseto, from), os.ModePerm)
        if err != nil {
            return err
        }
        isFirst := true
        return filepath.Walk(from, func(path string, info os.FileInfo, err error) error {
            if isFirst {
                isFirst = false
                return nil
            }
            return copy_thing(path, baseto)
        })
    } else {
        destination := filepath.Join(baseto, from)
        log.Printf("Copying file '%s' to '%s'...", from, destination)
        err := os.MkdirAll(path.Dir(destination), os.ModePerm)
        if err != nil {
            return err
        }
        fdest, err := os.Create(destination)
        if err != nil {
            return err
        }
        defer fdest.Close()
        forigin, err := os.Open(from)
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
