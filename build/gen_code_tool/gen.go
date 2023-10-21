package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	// 引用 flag包，将ablity name作为 flag 传入
	var abilityName string
	flag.StringVar(&abilityName, "ability", "", "ability name")
	flag.Parse()
	if err := verifyAbilityName(abilityName); err != nil {
		panic(err)
	}

	templatePath := "build/gen_code_tool/xxxx.go.tpl"
	outputPath := fmt.Sprintf("pkg/check/%s.go", strings.ToLower(abilityName))
	testFileTplPath := "build/gen_code_tool/xxxx_test.go.tpl"
	testFileOutputPath := fmt.Sprintf("pkg/check/%s_test.go", strings.ToLower(abilityName))
	if err := genFile(abilityName, templatePath, outputPath); err != nil {
		logrus.Fatal(err)
	}

	if err := genFile(abilityName, testFileTplPath, testFileOutputPath); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Now you can refer to \"pkg/check/cpu.go\", \"pkg/check/cpu_test.go\" to complete your item, Get, diff, Check and mockOption")
}

// 为空, 报错返回
// 不是字母数字下划线，报错返回
func verifyAbilityName(name string) error {
	if name == "" {
		return fmt.Errorf("ability name is empty")
	}
	for _, c := range name {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return fmt.Errorf("ability name %s is invalid, must be [0-9] or [a-z] or [_]", name)
		}
	}
	return nil
}

func genFile(abilityName, templatePath, outputPath string) error {
	// 如何outputpath存在，直接退出
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Warnf("Output file %s already exists", outputPath)
		return nil
	}
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		logrus.Errorf("Failed to read template file: %v", err)
		return errors.Trace(err)
	}

	// 首字母大写
	title := strings.Title(abilityName)
	outputBytes := []byte(strings.ReplaceAll(string(templateBytes), "XXXX", title))
	if err := os.WriteFile(outputPath, outputBytes, 0644); err != nil {
		logrus.Errorf("Failed to write output file: %v\n", err)
		return errors.Trace(err)
	}

	logrus.Infof("Generated %s", outputPath)
	return nil
}
