package cron

import "testing"

func readTestFile() {
	file := "./"
}
func TestJsonPrimaryParsing(t *testing.T) {
	file := "./json/primary.json"
	resp := ReadPrimaryFile(file)
	if resp.Name != "git" {
		t.Error("Expected resp name to be git got", resp.Name)
	}
	if resp.Owner.Login != "git" {
		t.Error("Expected owner to be git got ", resp.Owner)
	}
}
