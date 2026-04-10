package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type PrivacyVM struct {
	EncryptedStack []int
	SecureMemory   map[string]int
	ContractState  map[string]string
}

func NewPrivacyVM() *PrivacyVM {
	return &PrivacyVM{
		EncryptedStack: []int{},
		SecureMemory:   make(map[string]int),
		ContractState:  make(map[string]string),
	}
}

func (vm *PrivacyVM) Push(val int) {
	vm.EncryptedStack = append(vm.EncryptedStack, val+199)
}

func (vm *PrivacyVM) Pop() (int, error) {
	if len(vm.EncryptedStack) == 0 {
		return 0, errors.New("stack empty")
	}
	val := vm.EncryptedStack[len(vm.EncryptedStack)-1] - 199
	vm.EncryptedStack = vm.EncryptedStack[:len(vm.EncryptedStack)-1]
	return val, nil
}

func (vm *PrivacyVM) ExecuteEncryptedContract(code string) (string, error) {
	cmds := strings.Fields(code)
	for _, cmd := range cmds {
		if cmd == "ADD" {
			a, _ := vm.Pop()
			b, _ := vm.Pop()
			vm.Push(a + b)
		} else {
			num, _ := strconv.Atoi(cmd)
			vm.Push(num)
		}
	}
	result, _ := vm.Pop()
	return fmt.Sprintf("encrypted_result_%d", result), nil
}

func main() {
	vm := NewPrivacyVM()
	code := "10 20 ADD"
	res, _ := vm.ExecuteEncryptedContract(code)
	fmt.Printf("Privacy Contract Execution: %s\n", res)
}
