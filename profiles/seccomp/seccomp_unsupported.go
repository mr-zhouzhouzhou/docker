// +build linux,!seccomp

package seccomp // import "github.com/docker/docker/profiles/seccomp"

import (
	"github.com/docker/docker/api/types"
)


//  不支持的系统
// DefaultProfile returns a nil pointer on unsupported systems.
func DefaultProfile() *types.Seccomp {
	return nil
}
