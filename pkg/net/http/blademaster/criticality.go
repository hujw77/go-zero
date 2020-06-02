package blademaster

import (
	criticalityPkg "git.atmatrix.org/k12/zero/pkg/net/criticality"
	"git.atmatrix.org/k12/zero/pkg/net/metadata"

	"github.com/pkg/errors"
)

// Criticality is
func Criticality(pathCriticality criticalityPkg.Criticality) HandlerFunc {
	if !criticalityPkg.Exist(pathCriticality) {
		panic(errors.Errorf("This criticality is not exist: %s", pathCriticality))
	}
	return func(ctx *Context) {
		md, ok := metadata.FromContext(ctx)
		if ok {
			md[metadata.Criticality] = string(pathCriticality)
		}
	}
}