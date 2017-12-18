package helm

import (
	"reflect"

	"github.com/hashicorp/terraform/terraform"

	"testing"
)

func TestReleaseMigrateState(t *testing.T) {
	cases := map[string]struct {
		Version  int
		State    *terraform.InstanceState
		Expected *terraform.InstanceState
	}{
		"from 0 to 1": {
			Version: 0,
			State: &terraform.InstanceState{
				ID: "foo",
				Attributes: map[string]string{
					"id":                   "foo",
					"chart":                "../../charts/confluent",
					"keyring":              "",
					"metadata.#":           "1",
					"metadata.0.chart":     "confluent",
					"metadata.0.name":      "confluent",
					"metadata.0.namespace": "production",
					"metadata.0.revision":  "24",
					"metadata.0.status":    "DEPLOYED",
					"metadata.0.version":   "0.1.0",
					"name":                 "confluent",
					"namespace":            "production",
					"reuse_values":         "true",
					"timeout":              "300",
					"values":               "kafka:\n  persistence:\n    enabled: true",
				},
			},
			Expected: &terraform.InstanceState{
				ID: "foo",
				Attributes: map[string]string{
					"id":                   "foo",
					"chart":                "../../charts/confluent",
					"keyring":              "",
					"metadata.#":           "1",
					"metadata.0.chart":     "confluent",
					"metadata.0.name":      "confluent",
					"metadata.0.namespace": "production",
					"metadata.0.status":    "DEPLOYED",
					"metadata.0.version":   "0.1.0",
					"name":                 "confluent",
					"namespace":            "production",
					"reuse_values":         "true",
					"timeout":              "300",
					"values":               "kafka:\n  persistence:\n    enabled: true",
				},
			},
		},
	}

	for id, c := range cases {
		state, err := resourceReleaseMigrateState(c.Version, c.State, nil)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(state, c.Expected) {
			t.Errorf("Failed %q: %+v expected\n\ngot: %+v", id, c.Expected, state)
		}
	}
}
