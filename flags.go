package flags

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
)

type Key [2]string
type FlagSet struct {
	sync.Mutex
	flags map[string]interface{}
	help  map[Key]string
}

var (
	flagset = NewFlagSet()
)

func SetFlag(variable interface{}, flags Key, helpstr string) {
	flagset.SetFlag(variable, flags, helpstr)
}

func Parse(args []string) ([]string, error) {
	return flagset.Parse(args)
}

func Help() [][2]string {
	return flagset.Help()
}

func NewFlagSet() *FlagSet {
	return &FlagSet{
		flags: map[string]interface{}{},
		help:  map[Key]string{},
	}
}

func (f *FlagSet) SetFlag(variable interface{}, flags Key, helpstr string) {
	f.Lock()
	defer f.Unlock()
	for _, v := range flags {
		f.flags[v] = variable
	}

	if len(flags[1]) != 0 && len(flags[1]) < len(flags[0]) {
		flags = Key{flags[1], flags[0]}
	}

	f.help[flags] = helpstr
}

func (f *FlagSet) Parse(args []string) ([]string, error) {
	f.Lock()
	defer f.Unlock()
	remaining := []string{}
	numArgs := len(args)
	for i, v := range args {
		if v == "--" {
			numArgs = i
			break
		}
	}
	for i := 0; i < len(args[:numArgs]); i++ {
		variable, ok := f.flags[args[i]]
		if !ok {
			remaining = append(remaining, args[i])
			continue
		}

		switch v := variable.(type) {
		case *bool:
			*v = true
		case *int:
			if !(i+1 < len(args)) {
				return remaining, errors.New("Malformed flag, require input")
			}

			tmp, err := strconv.ParseInt(args[i+1], 10, 0)
			if err != nil {
				return nil, err
			}

			*v = int(tmp)
			i++

		case *float64:
			if !(i+1 < len(args)) {
				return remaining, errors.New("Malformed flag, require input")
			}
			tmp, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				return nil, err
			}
			*v = tmp
			i++

		case *string:
			if !(i+1 < len(args)) {
				return remaining, errors.New("Malformed flag, require input")
			}
			*v = args[i+1]
			i++

		default:
			return remaining, errors.New(fmt.Sprintf("Cannot parse type of variable: %v", v))
		}
	}
	if numArgs != len(args) {
		remaining = args[numArgs+1:]
	}
	return remaining, nil
}

func (f *FlagSet) Help() [][2]string {
	f.Lock()
	defer f.Unlock()
	result := [][2]string{}
	doubles := []Key{}
	singles := []Key{}
	for k := range f.help {
		if k[1] != "" {
			doubles = append(doubles, k)
		} else {
			singles = append(singles, k)
		}
	}

	sort.Slice(doubles, func(i, j int) bool {
		return doubles[i][0] < doubles[j][0]
	})
	sort.Slice(singles, func(i, j int) bool {
		return singles[i][0] < singles[j][0]
	})
	sorted_keys := append(doubles, singles...)
	for _, k := range sorted_keys {
		helpstr := f.help[k]
		keys := ""
		for _, elm := range k {
			if elm != "" {
				keys += elm + ", "
			}
		}
		result = append(result, [2]string{keys[:len(keys)-2], helpstr})
	}
	return result
}
