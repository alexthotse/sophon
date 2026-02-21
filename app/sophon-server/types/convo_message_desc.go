package types

import (
	"sophon-server/db"

	shared "sophon-shared"
)

func HasPendingBuilds(planDescs []*db.ConvoMessageDescription) bool {
	apiDescs := make([]*shared.ConvoMessageDescription, len(planDescs))
	for i, desc := range planDescs {
		apiDescs[i] = desc.ToApi()
	}

	return shared.HasPendingBuilds(apiDescs)
}
