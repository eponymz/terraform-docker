package test

import (
	"fmt"
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
)

var validJSON = fmt.Sprint(`{"resource_changes":[{"address":"module.address[0]","change":{"actions":["delete"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[1]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[2]","change":{"actions":["delete"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[3]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[4]","change":{"actions":["delete"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[5]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[6]","change":{"actions":["delete"],"before":{},"after":{},"after_unknown":{}}},{"address":"module.address[7]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}}]}`)

var invalidJSON = fmt.Sprint(`[{"address":"module.address[0]","change":{"actions":["delete"]"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[1]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[2]","change":{"actions":["delete"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[3]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[4]","change":{"actions":["delete"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[5]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}},{"address":"module.address[6]","change":{"actions":["delete"],"before":{},"after":{},"after_unknown":{}}},{"address":"module.address[7]","change":{"actions":["no-op"],"before":{},"after":null,"after_unknown":{}}}]`)

func TestValidateJSON(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		got := util.ValidateJSON(validJSON)
		wants := true
		if got != wants {
			t.Fatalf("ValidateJSON should return %t with valid JSON, got: %t", wants, got)
		}
	})
	t.Run("Invalid", func(t *testing.T) {
		got := util.ValidateJSON(invalidJSON)
		wants := false
		if got != wants {
			t.Fatalf("ValidateJSON should return %t with invalid JSON, got: %t", wants, got)
		}
	})
}

func TestGetJSON(t *testing.T) {
	t.Run("Filtered Array", func (t *testing.T) {
		got := util.GetJSON(validJSON, `resource_changes.#(change.actions.#(=="delete"))#.address`).String()
		wants := `["module.address[0]","module.address[2]","module.address[4]","module.address[6]"]`
		if got != wants {
			t.Fatalf("GetJSON should return the value at the specified path. Got: %s Wants: %s", got, wants)
		}
	})
	t.Run("Negative Filtered Array", func (t *testing.T) {
		got := util.GetJSON(validJSON, `resource_changes.#(change.actions.#(!="delete"))#.address`).String()
		wants := `["module.address[1]","module.address[3]","module.address[5]","module.address[7]"]`
		if got != wants {
			t.Fatalf("GetJSON should return the value at the specified path. Got: %s Wants: %s", got, wants)
		}
	})
	t.Run("Specific Value", func (t *testing.T) {
		got := util.GetJSON(validJSON, `resource_changes.0.address`).String()
		wants := "module.address[0]"
		if got != wants {
			t.Fatalf("GetJSON should return the value at the specified path. Got: %s Wants: %s", got, wants)
		}
	})
}
