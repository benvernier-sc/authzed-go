//go:build mage

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/magefile/mage/mg"
)

type Gen mg.Namespace

// All runs all generators in parallel
func (g Gen) All() error {
	mg.Deps(g.Proto)
	return nil
}

const (
	ProtoPath     = "proto/authzed/api"
	BufRepository = "buf.build/authzed/api"
	BufTag        = "d12b0c6cb910ec1b21bb4abc5ee43f53efb9896c"
)

// Proto runs proto codegen
func (Gen) Proto() error {
	bufRef := BufRepository + ":" + BufTag
	fmt.Println("generating", bufRef)
	runDirV("magefiles", "go", "run", "github.com/bufbuild/buf/cmd/buf", "generate", bufRef)
	return generateVersionFiles()
}

func generateVersionFiles() error {
	tmpl, err := template.ParseFiles("magefiles/version.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse version template: %w", err)
	}

	entries, err := os.ReadDir(ProtoPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() && !strings.HasSuffix(entry.Name(), "_test") {
			var b bytes.Buffer
			tmpl.Execute(&b, map[string]string{
				"package": entry.Name(),
				"bufRepo": BufRepository,
				"bufTag":  BufTag,
			})

			versionPath := filepath.Join(ProtoPath, entry.Name(), "zz_generated.version.go")
			if err := os.WriteFile(versionPath, b.Bytes(), 0o644); err != nil {
				return err
			}
		}
	}
	return nil
}
