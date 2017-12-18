package helm

import (
	"fmt"

	"github.com/hashicorp/terraform/terraform"
)

func resourceReleaseMigrateState(currentVersion int, state *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch currentVersion {
	case 0:
		return resourceReleaseMigrateState0to1(state, meta)
	default:
		return state, fmt.Errorf("unexpected schema version: %d", currentVersion)
	}
}

func resourceReleaseMigrateState0to1(state *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	if state.Empty() {
		return state, nil
	}

	delete(state.Attributes, "metadata.0.revision")

	return state, nil
}
