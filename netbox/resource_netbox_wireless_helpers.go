package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-viper/mapstructure/v2"
)

type wirelessInterceptWriter struct {
	runtime.ClientRequest
	fields map[string]any
}

func (iw wirelessInterceptWriter) SetBodyParam(p any) error {
	out := make(map[string]any)
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &out,
	})
	if err != nil {
		return err
	}
	if err := dec.Decode(p); err != nil {
		return err
	}
	for fieldName, value := range iw.fields {
		out[fieldName] = value
	}
	return iw.ClientRequest.SetBodyParam(out)
}

type wirelessInterceptParams struct {
	inner  runtime.ClientRequestWriter
	fields map[string]any
}

func (ip wirelessInterceptParams) WriteToRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	writer := wirelessInterceptWriter{ClientRequest: req, fields: ip.fields}
	return ip.inner.WriteToRequest(writer, reg)
}

func hackSerializeWirelessWithValues(fields map[string]any) wireless.ClientOption {
	overrideFields := make(map[string]any, len(fields))
	for fieldName, value := range fields {
		overrideFields[fieldName] = value
	}
	return func(co *runtime.ClientOperation) {
		originalParams := co.Params
		co.Params = wirelessInterceptParams{inner: originalParams, fields: overrideFields}
	}
}

func hackSerializeWirelessAsNull(fields ...string) wireless.ClientOption {
	overrideFields := make(map[string]any, len(fields))
	for _, field := range fields {
		overrideFields[field] = nil
	}
	return hackSerializeWirelessWithValues(overrideFields)
}
