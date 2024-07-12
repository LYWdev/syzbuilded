// Copyright 2024 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

// gen generates instruction tables (ifuzz_types/insns.go) from ARM64 JSON.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/syzkaller/pkg/ifuzz/arm64"
	"github.com/google/syzkaller/pkg/serializer"
	"github.com/google/syzkaller/pkg/tool"
)

func main() {
	if len(os.Args) != 2 {
		tool.Failf("usage: gen arm64.json")
	}
	jsonStr, err := os.ReadFile(os.Args[1])
	if err != nil {
		tool.Failf("failed to open input file: %v", err)
	}
	insns := JSONToInsns(jsonStr)

	fmt.Printf(`// Code generated by pkg/ifuzz/gen. DO NOT EDIT.

// go:build !codeanalysis

package generated

import (
	. "github.com/google/syzkaller/pkg/ifuzz/arm64"
)

func init() {
	Register(insns_arm64)
}

var insns_arm64 = 
`)
	serializer.Write(os.Stdout, insns)

	fmt.Fprintf(os.Stderr, "handled %v\n", len(insns))
}

type insnDesc struct {
	Name   string
	Bits   string
	Arch   string
	Syntax string
	Code   string
	Alias  string
}

func isPrivateInsn(insn arm64.Insn) bool {
	switch insn.Name {
	case "AT", "DC", "IC", "SYS", "SYSL", "TLBI":
		return true
	}
	return false
}

func JSONToInsns(jsonStr []byte) []*arm64.Insn {
	var insnDescriptions []insnDesc
	err := json.Unmarshal(jsonStr, &insnDescriptions)
	if err != nil {
		return nil
	}
	ret := []*arm64.Insn{}
	for _, desc := range insnDescriptions {
		mask := uint32(0)
		opcode := uint32(0)
		curBit := uint(31)
		fields := []arm64.InsnField{}
		pieces := strings.Split(desc.Bits, "|")
		for _, piece := range pieces {
			size := uint(1)
			pair := strings.Split(piece, ":")
			var pattern = piece
			if len(pair) == 2 {
				size64, err := strconv.ParseUint(pair[1], 10, 0)
				if err != nil {
					return nil
				}
				size = uint(size64)
				pattern = pair[0]
			}
			updateOpcode := true
			opPart := uint32(0)
			maskPart := uint32(0)
			if pattern[0:1] != "(" {
				number, err := strconv.ParseUint(pattern, 2, 32)
				if err != nil {
					// This is a named region.
					field := arm64.InsnField{
						Name:   pattern,
						Start:  curBit,
						Length: size,
					}
					fields = append(fields, field)
					updateOpcode = false
				} else {
					// This is a binary mask.
					opPart = uint32(number)
					maskPart = ((1 << size) - 1)
				}
			}
			opcode <<= size
			mask <<= size
			curBit -= size
			if updateOpcode {
				opcode |= opPart
				mask |= maskPart
			}
		}
		templ := arm64.Insn{
			Name:       desc.Name,
			OpcodeMask: mask,
			Opcode:     opcode,
			Fields:     fields,
			AsUInt32:   opcode,
		}
		templ.Priv = isPrivateInsn(templ)
		insn := new(arm64.Insn)
		*insn = templ
		ret = append(ret, insn)
	}
	return ret
}
