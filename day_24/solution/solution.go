package solution

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

func (wires Wire) build_binary(operations []Operation) map[string]bool {
	remaining_ops := operations

	wrong_wire := map[string]bool{}
	for len(remaining_ops) > 0 {
		new_ops := []Operation{}

		for _, operation := range remaining_ops {
			_, exists1 := wires[operation.wire_1]
			_, exists2 := wires[operation.wire_2]

			if exists1 && exists2 {
				if operation.wire_result[0] == 'z' && operation.op != "XOR" && operation.wire_result != "z45" {
					wrong_wire[operation.wire_result] = true
				}

				if operation.wire_result[0] != 'z' && operation.wire_1[0] != 'x' && operation.wire_1[0] != 'y' && operation.op == "XOR" {
					wrong_wire[operation.wire_result] = true
				}

				if operation.op == "AND" && operation.wire_1 != "x00" && operation.wire_2 != "x00" {
					for _, subop := range operations {
						if (operation.wire_result == subop.wire_1 || operation.wire_result == subop.wire_2) && subop.op == "XOR" {
							wrong_wire[operation.wire_result] = true
						}
					}
				}

				if operation.op == "XOR" {
					for _, subop := range operations {
						if (operation.wire_result == subop.wire_1 || operation.wire_result == subop.wire_2) && subop.op == "OR" {
							wrong_wire[operation.wire_result] = true
						}
					}
				}

				switch operation.op {
				case "AND":
					wires[operation.wire_result] = wires[operation.wire_1] & wires[operation.wire_2]
				case "OR":
					wires[operation.wire_result] = wires[operation.wire_1] | wires[operation.wire_2]
				case "XOR":
					wires[operation.wire_result] = wires[operation.wire_1] ^ wires[operation.wire_2]
				}
			} else {
				new_ops = append(new_ops, operation)
			}
		}
		remaining_ops = new_ops
	}

	return wrong_wire
}

func (wires Wire) binary_to_int() int64 {
	binary_str := []string{}
	for wire := range wires {
		if wire[0] == 'z' {
			binary_str = append(binary_str, wire)
		}
	}

	sort.Slice(binary_str, func(i, j int) bool {
		return binary_str[i] > binary_str[j]
	})

	binary := ""
	for _, b := range binary_str {
		binary += strconv.Itoa(wires[b])
	}

	number, _ := strconv.ParseInt(binary, 2, 64)

	return number
}

func Part_One(wire Wire, operations []Operation) int64 {
	defer Track(time.Now(), "Part 1")

	wire.build_binary(operations)
	res := wire.binary_to_int()

	return res
}

func Part_Two(wire Wire, operations []Operation) string {
	defer Track(time.Now(), "Part 2")

	wrong_wire := wire.build_binary(operations)

	answer := []string{}
	for wire := range wrong_wire {
		answer = append(answer, wire)
	}

	sort.Strings(answer)

	return strings.Join(answer, ",")
}
