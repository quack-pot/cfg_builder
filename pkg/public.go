package cfg_builder

import "github.com/quack-pot/cfg_builder/internal"

// *=================================================
// *
// * Error Wrappers
// * - Keeps variables readonly
// *
// *=================================================

func ErrInvalidDst() error { return internal.ErrInvalidDst }

// *=================================================
// *
// * Type Wrappers
// *
// *=================================================

type IConfig = internal.IConfig
type IConfigBuilder = internal.IConfigBuilder
type IConfigProvider = internal.IConfigProvider

// *=================================================
// *
// * Constructors
// *
// *=================================================

func NewConfigBuilder() IConfigBuilder {
	return internal.NewConfigBuilder()
}
