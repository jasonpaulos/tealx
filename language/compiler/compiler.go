package compiler

import (
	"fmt"
	"io"
	"strconv"

	"github.com/jasonpaulos/tealx/language"
	"github.com/jasonpaulos/tealx/language/element"
)

type opWriter struct {
	io.StringWriter
}

func (w opWriter) WriteOp(op language.Operation) error {
	_, err := w.StringWriter.WriteString(op.String() + "\n")
	return err
}

func subroutineNameToLabel(name string) string {
	return "sub_" + name
}

func Compile(program element.Program, stringWriter io.StringWriter) error {
	w := opWriter{StringWriter: stringWriter}
	err := w.WriteOp(language.Operation{Opcode: "#pragma", Arguments: []string{"version", strconv.FormatUint(program.Version, 10)}})
	if err != nil {
		return err
	}

	mainCFG := program.Main.Codegen(element.CodegenContext{})
	subroutineCFGs := make([]language.ControlFlowGraph, len(program.Subroutines))
	for i, subroutine := range program.Subroutines {
		subroutineCFGs[i] = subroutine.Codegen(element.CodegenContext{})
	}

	mainOperations := flatten(mainCFG.Sort(), "main")
	for _, op := range mainOperations {
		err = w.WriteOp(op)
		if err != nil {
			return err
		}
	}

	for i, cfg := range subroutineCFGs {
		label := subroutineNameToLabel(program.Subroutines[i].Name)
		subroutineOps := flatten(cfg.Sort(), label)

		err = w.WriteOp(language.Operation{Opcode: label + ":"})
		if err != nil {
			return err
		}

		for _, op := range subroutineOps {
			err = w.WriteOp(op)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func blockIndex(blocks []*language.CodeBlock, target *language.CodeBlock) int {
	for i, block := range blocks {
		if block == target {
			return i
		}
	}
	panic("block not fond")
}

func blockIndexToLabel(index int, prefix string) string {
	return fmt.Sprintf("%s_l%d", prefix, index)
}

func flatten(blocks []*language.CodeBlock, labelPrefix string) []language.Operation {
	var codeblocks [][]language.Operation
	references := make(map[*language.CodeBlock]int)

	for i, block := range blocks {
		code := block.Operations[:]

		if block.Outgoing == nil {
			// TODO: also check if this block has an op that would make outgoing dead code, i.e. return, retsub, or err
			codeblocks = append(codeblocks, code)
			continue
		}

		switch outgoing := block.Outgoing.(type) {
		case *language.UnconditionalJump:
			next := outgoing.Next
			nextIndex := blockIndex(blocks, next)
			if nextIndex != i+1 {
				references[next]++
				code = append(code, language.Operation{Opcode: "b", Arguments: []string{blockIndexToLabel(nextIndex, labelPrefix)}})
			}
		case *language.ConditionalJump:
			trueBranch := outgoing.TrueBranch
			falseBranch := outgoing.FalseBranch

			trueBranchIndex := blockIndex(blocks, trueBranch)
			falseBranchIndex := blockIndex(blocks, falseBranch)

			if falseBranchIndex == i+1 {
				references[trueBranch]++
				code = append(code, language.Operation{Opcode: "bnz", Arguments: []string{blockIndexToLabel(trueBranchIndex, labelPrefix)}})
				codeblocks = append(codeblocks, code)
				continue
			}

			if trueBranchIndex == i+1 {
				references[falseBranch]++
				code = append(code, language.Operation{Opcode: "bz", Arguments: []string{blockIndexToLabel(falseBranchIndex, labelPrefix)}})
				codeblocks = append(codeblocks, code)
				continue
			}

			references[trueBranch]++
			code = append(code, language.Operation{Opcode: "bnz", Arguments: []string{blockIndexToLabel(trueBranchIndex, labelPrefix)}})

			references[falseBranch]++
			code = append(code, language.Operation{Opcode: "b", Arguments: []string{blockIndexToLabel(falseBranchIndex, labelPrefix)}})
		case *language.MatchJump:
			targetLabels := make([]string, len(outgoing.Targets))
			for i, target := range outgoing.Targets {
				index := blockIndex(blocks, target)
				targetLabels[i] = blockIndexToLabel(index, labelPrefix)
				references[target]++
			}
			code = append(code, language.Operation{Opcode: "match", Arguments: targetLabels})

			defaultPathIndex := blockIndex(blocks, outgoing.DefaultPath)
			if defaultPathIndex != i+1 {
				references[outgoing.DefaultPath]++
				code = append(code, language.Operation{Opcode: "b", Arguments: []string{blockIndexToLabel(defaultPathIndex, labelPrefix)}})
			}
		default:
			panic("unknown jump type")
		}

		codeblocks = append(codeblocks, code)
	}

	var ops []language.Operation

	for i, code := range codeblocks {
		if references[blocks[i]] != 0 {
			ops = append(ops, language.Operation{Opcode: blockIndexToLabel(i, labelPrefix) + ":"})
		}
		ops = append(ops, code...)
	}

	return ops
}
