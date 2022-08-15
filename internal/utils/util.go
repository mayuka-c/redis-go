package utils

import (
	"encoding/json"
	"fmt"
)

// DeepCopy copies fields from in to out using json tags out should be address
func DeepCopy(in interface{}, out interface{}) {
	temp, _ := json.Marshal(in)
	fmt.Println(temp)
	json.Unmarshal(temp, out)
}
