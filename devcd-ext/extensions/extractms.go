package extensions

import "devcd_ext/services"

type ExtractMs interface {
	Extract()
}

func NewExtractMs() ExtractMs {
	// add your custom extract script to override default extract
	return services.DefaultExtractms{}
}
