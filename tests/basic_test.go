package main

import (
	"fmt"
	"testing"

	"github.com/Mattemagikern/flags"
)

func TestBasic(t *testing.T) {
	t.Run("String", parseString)
	t.Run("Bool", parseBool)
	t.Run("Int", parseInt)
	t.Run("Float", parseFloat)
	t.Run("-- args", inputArgs)
}

func parseBool(t *testing.T) {
	b := false
	args := []string{"-b", "-str", "hello", "stuff"}
	f := flags.NewFlagSet()
	f.SetFlag(&b, [2]string{"-b"}, "Help")
	_, err := f.Parse(args)
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("Bool not true")
	}

}

func parseString(t *testing.T) {
	str := ""
	str2 := "tmp"
	f := flags.NewFlagSet()
	f.SetFlag(&str, [2]string{"--str", "-s"}, "Help")
	args := []string{"--str", "hello", "stuff"}
	_, err := f.Parse(args)

	if err != nil {
		t.Error(err)
	}

	if str != "hello" {
		t.Error("str != hello")
	}

	str = str2

	_, err = f.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if str != "hello" {
		t.Error("str != hello")
	}

}

func parseInt(t *testing.T) {
	i := 0
	f := flags.NewFlagSet()
	f.SetFlag(&i, [2]string{"-num"}, "Help")
	args := []string{"-num", "32", "stuff"}
	_, err := f.Parse(args)
	if err != nil {
		t.Error(err)
	}
	if i != 32 {
		t.Error("int not parsed")
	}
}

func parseFloat(t *testing.T) {
	i := 0.0
	f := flags.NewFlagSet()
	f.SetFlag(&i, [2]string{"-num"}, "Help")
	args := []string{"-num", "3.14159", "stuff"}
	_, err := f.Parse(args)
	if err != nil {
		t.Error(err)
	}
	if i != 3.14159 {
		t.Error("float not parsed correctly")
	}
}

func TestHelp(t *testing.T) {
	t.Run("help", testHelp)
}

func testHelp(t *testing.T) {
	help := false
	str := ""
	args := []string{"--str", "bonjour", "--help", "woop", "stuff"}
	f := flags.NewFlagSet()
	f.SetFlag(&help, [2]string{"-h", "--help"}, "Help")
	f.SetFlag(&str, [2]string{"--str", "-s"}, "Help")
	f.SetFlag(&help, [2]string{"-h", "--help"}, "Help")
	f.SetFlag(&str, [2]string{"--tri"}, "hello, should be at the bottom")
	f.SetFlag(&str, [2]string{"--ssh"}, "ssh help string")
	remaining, err := f.Parse(args)
	if err != nil {
		t.Error(err)
	}
	if len(remaining) != 2 {
		t.Error("remaining not as expected")
	}

	if !help {
		t.Fatal("Help not found")
	}
	res := f.Help()
	if "--tri" != res[len(res)-1][0] {
		fmt.Print(res)
		t.Fatal("Not sorted")
	}
}

func inputArgs(t *testing.T) {
	b := false
	str := ""
	args := []string{"hello", "--", "stuff"}
	f := flags.NewFlagSet()
	f.SetFlag(&b, [2]string{"-b"}, "Help")
	f.SetFlag(&str, [2]string{"-str", "-s"}, "Help")
	remaining, err := f.Parse(args)
	if err != nil {
		t.Error(err)
	}
	if len(remaining) != 1 {
		t.Fatal("No remaining parsed")
	}
	if remaining[0] != "stuff" {
		t.Error("remaining not ok")
	}
}
