package app

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/furya-official/fury-blockchain/app/params"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing
func MakeTestEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
