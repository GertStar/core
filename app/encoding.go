package app

import (
	enccodec "github.com/tharsis/ethermint/encoding/codec"

	"github.com/Gert-Chain/core/app/params"
)

// MakeEncodingConfig creates an EncodingConfig and registers the app's types on it.
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
