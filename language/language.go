package language

import (
	"errors"
	"fmt"
	"strings"
)

type Operation struct {
	Opcode    string
	Arguments []string
}

func (o Operation) String() string {
	if len(o.Arguments) == 0 {
		return o.Opcode
	}
	return o.Opcode + " " + strings.Join(o.Arguments, " ")
}

type CodeBlock struct {
	Operations []Operation
	Outgoing   Jump
	Incoming   []*CodeBlock
}

func (cb *CodeBlock) Merge(other *CodeBlock) error {
	if cb.Outgoing != nil {
		return errors.New("cannot merge into a CodeBlock that has jumps")
	}
	if len(other.Incoming) != 0 {
		return errors.New("cannot merge a CodeBlock that has an incoming node")
	}
	cb.Operations = append(cb.Operations, other.Operations...)
	cb.Outgoing = other.Outgoing
	if other.Outgoing != nil {
		other.Outgoing.ReplaceIncoming(other, cb)
	}
	return nil
}

type Jump interface {
	GetNodes() []*CodeBlock
	ReplaceIncoming(oldIncoming, newIncoming *CodeBlock)
}

type UnconditionalJump struct {
	Next *CodeBlock
}

func (j *UnconditionalJump) GetNodes() []*CodeBlock {
	return []*CodeBlock{j.Next}
}

func (j *UnconditionalJump) ReplaceIncoming(oldIncoming, newIncoming *CodeBlock) {
	if j.Next == oldIncoming {
		j.Next = newIncoming
	}
}

type ConditionalJump struct {
	TrueBranch  *CodeBlock
	FalseBranch *CodeBlock
}

func (j *ConditionalJump) GetNodes() []*CodeBlock {
	return []*CodeBlock{j.TrueBranch, j.FalseBranch}
}

func (j *ConditionalJump) ReplaceIncoming(oldIncoming, newIncoming *CodeBlock) {
	if j.TrueBranch == oldIncoming {
		j.TrueBranch = newIncoming
	}
	if j.FalseBranch == oldIncoming {
		j.FalseBranch = newIncoming
	}
}

type ControlFlowGraph struct {
	Start *CodeBlock
	End   *CodeBlock
}

func MakeControlFlowGraph(operations []Operation) ControlFlowGraph {
	singleBlock := CodeBlock{Operations: operations}
	return ControlFlowGraph{
		Start: &singleBlock,
		End:   &singleBlock,
	}
}

func (g *ControlFlowGraph) Append(h ControlFlowGraph) error {
	// combine g.End and h.Start
	err := g.End.Merge(h.Start)
	if err != nil {
		return err
	}

	if h.Start != h.End {
		// If h is not a single-node graph, update g.End to point to the new
		// end node for the graph
		g.End = h.End
	}
	return nil
}

func (g *ControlFlowGraph) AppendConditional(trueBranch, falseBranch ControlFlowGraph) error {
	if g.End.Outgoing != nil {
		return errors.New("end node must be terminal")
	}
	if trueBranch.End.Outgoing != nil {
		return errors.New("true branch end node must be terminal")
	}
	if falseBranch.End.Outgoing != nil {
		return errors.New("false branch end node must be terminal")
	}

	// newEnd will postdominate trueBranch and falseBranch to become the new graph end node
	newEnd := CodeBlock{
		Incoming: []*CodeBlock{
			trueBranch.End,
			falseBranch.End,
		},
	}
	trueBranch.End.Outgoing = &UnconditionalJump{
		Next: &newEnd,
	}
	falseBranch.End.Outgoing = &UnconditionalJump{
		Next: &newEnd,
	}

	g.End.Outgoing = &ConditionalJump{
		TrueBranch:  trueBranch.Start,
		FalseBranch: falseBranch.Start,
	}
	trueBranch.Start.Incoming = append(trueBranch.Start.Incoming, g.End)
	falseBranch.Start.Incoming = append(falseBranch.Start.Incoming, g.End)

	g.End = &newEnd
	return nil
}

func (g *ControlFlowGraph) Sort() []*CodeBlock {
	var order []*CodeBlock

	stack := []*CodeBlock{g.Start}
	visited := make(map[*CodeBlock]bool)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := visited[node]; ok {
			continue
		}

		if node.Outgoing != nil {
			stack = append(stack, node.Outgoing.GetNodes()...)
		}

		order = append(order, node)
		// presence in visited map is all that matters
		visited[node] = false
	}

	endIndex := -1
	for i, block := range order {
		if block == g.End {
			endIndex = i
			break
		}
	}

	if endIndex == -1 {
		panic("end block not present")
	}

	if endIndex != len(order)-1 {
		// Make sure end block is last so that the program doesn't run off into another block
		copy(order[endIndex:], order[endIndex+1:])
		order[len(order)-1] = g.End
	}

	return order
}

func blockIndex(blocks []*CodeBlock, target *CodeBlock) int {
	for i, block := range blocks {
		if block == target {
			return i
		}
	}
	panic("block not fond")
}

func Flatten(blocks []*CodeBlock) []Operation {
	var codeblocks [][]Operation
	references := make(map[*CodeBlock]int)

	for i, block := range blocks {
		code := block.Operations[:]

		if block.Outgoing == nil {
			// TODO: also check if this block has an op that would make outgoing dead code, i.e. return, retsub, or err
			codeblocks = append(codeblocks, code)
			continue
		}

		switch outgoing := block.Outgoing.(type) {
		case *UnconditionalJump:
			next := outgoing.Next
			nextIndex := blockIndex(blocks, next)
			if nextIndex != i+1 {
				references[next]++
				code = append(code, Operation{Opcode: "b", Arguments: []string{fmt.Sprintf("l%d", nextIndex)}})
			}
		case *ConditionalJump:
			trueBranch := outgoing.TrueBranch
			falseBranch := outgoing.FalseBranch

			trueBranchIndex := blockIndex(blocks, trueBranch)
			falseBranchIndex := blockIndex(blocks, falseBranch)

			if falseBranchIndex == i+1 {
				references[trueBranch]++
				code = append(code, Operation{Opcode: "bnz", Arguments: []string{fmt.Sprintf("l%d", trueBranchIndex)}})
				codeblocks = append(codeblocks, code)
				continue
			}

			if trueBranchIndex == i+1 {
				references[falseBranch]++
				code = append(code, Operation{Opcode: "bz", Arguments: []string{fmt.Sprintf("l%d", falseBranchIndex)}})
				codeblocks = append(codeblocks, code)
				continue
			}

			references[trueBranch]++
			code = append(code, Operation{Opcode: "bnz", Arguments: []string{fmt.Sprintf("l%d", trueBranchIndex)}})

			references[falseBranch]++
			code = append(code, Operation{Opcode: "b", Arguments: []string{fmt.Sprintf("l%d", falseBranchIndex)}})
		default:
			panic("unknown jump type")
		}

		codeblocks = append(codeblocks, code)
	}

	var ops []Operation

	for i, code := range codeblocks {
		if references[blocks[i]] != 0 {
			ops = append(ops, Operation{Opcode: fmt.Sprintf("l%d:", i)})
		}
		ops = append(ops, code...)
	}

	return ops
}

func Serialize(ops []Operation) string {
	opsAsStrings := make([]string, len(ops))
	for i, op := range ops {
		opsAsStrings[i] = op.String()
	}
	return strings.Join(opsAsStrings, "\n")
}
