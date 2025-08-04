package utils

import (
	"fmt"
	"testing"
)

func TestGetStructSchema(t *testing.T) {

	type Test struct {
		Abc   []string `json:"abc" binding:"required"`
		Name  string   `json:"name" binding:"required"`
		Pass  string   `json:"pass" binding:"required"`
		Email string   `json:"email" binding:"required"`
	}

	type KubernetesConfiguration struct {
		ApiServerEndpoint string                 `json:"api-server-endpoint" binding:"required"`
		Token             string                 `json:"token" binding:"required"`
		Test              map[string]interface{} `json:"test" binding:"required"`
		Test2             Test                   `json:"test2" binding:"required"`
		Test3             []Test                 `json:"test3" binding:"required"`
	}

	schema, err := GetSchema(KubernetesConfiguration{})
	fmt.Println(schema)
	if err != nil {
		t.Fatalf("Failed to get schema: %v", err)
	}
}

func TestChangeInto(t *testing.T) {

	a := map[string]interface{}{
		"abc":   []string{"1", "2", "3"},
		"name":  "John",
		"pass":  "123456",
		"email": "john@example.com",
	}

	type Test struct {
		Abc   []string `json:"abc" binding:"required"`
		Name  string   `json:"name" binding:"required"`
		Pass  string   `json:"pass" binding:"required"`
		Email string   `json:"email" binding:"required"`
	}

	var b Test
	err := ChangeInto(a, &b)
	if err != nil {
		t.Fatalf("Failed to change into: %v", err)
	}
	fmt.Println(b)
}
