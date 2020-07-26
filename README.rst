#######
 Flags
#######

Overview
----------
This is a flag parsing library.
The reason to use this instead of the Golang flags library that this library
allows you to re-parse, add, remove flags during execution. You can parse flags
from any array of strings which makes it ideal to use when writing cli programs.


Example
----------

.. code-block:: golang

   package main

   import (
       "fmt"

       "github.com/Mattemagikern/flags"
   )

   var (
       /* Default values */
       version bool    = false
       help    bool    = false
       i       int     = 0
       fl      float64 = 0.0
       str     string  = ""
   )

   func main() {
       args := []string{"-v", "-h", "-i", "42", "-fl", "5.4", "--str", "hello world!"}
       f := flags.NewFlagSet()
       f.SetFlag(&version, flags.Key{"-v", "--version"}, "Prints version")
       f.SetFlag(&help, flags.Key{"-h", "--help"}, "Prints this message")
       f.SetFlag(&i, flags.Key{"-i", "--integer"}, "parses an integer")
       f.SetFlag(&fl, flags.Key{"-fl", "--float"}, "parses a float")
       f.SetFlag(&str, flags.Key{"-s", "--str"}, "parses a string")
       args, err := f.Parse(args)
       if err != nil {
           fmt.Println(err)
           return
       }

       fmt.Println(version, help, i, fl, str)

       if help {
           fmt.Print(f.Help())
       }
   }
