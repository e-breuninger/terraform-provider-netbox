package netbox

import (
	"encoding/json"
	"io"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

const defaultTagWeight = 1000

type tagAPIFields struct {
	ObjectTypes []string `json:"object_types"`
	Weight      int64    `json:"weight"`
}

type tagAPIModel struct {
	models.Tag
	tagAPIFields
}

type tagRequestWriter struct {
	runtime.ClientRequest
	fields tagAPIFields
}

func (w tagRequestWriter) SetBodyParam(payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	values := make(map[string]any)
	if err := json.Unmarshal(body, &values); err != nil {
		return err
	}
	values["object_types"] = w.fields.ObjectTypes
	values["weight"] = w.fields.Weight
	return w.ClientRequest.SetBodyParam(values)
}

type tagRequestParams struct {
	inner  runtime.ClientRequestWriter
	fields tagAPIFields
}

func (p tagRequestParams) WriteToRequest(
	request runtime.ClientRequest,
	registry strfmt.Registry,
) error {
	writer := tagRequestWriter{
		ClientRequest: request,
		fields:        p.fields,
	}
	return p.inner.WriteToRequest(writer, registry)
}

func serializeTagAPIFields(fields tagAPIFields) extras.ClientOption {
	return func(operation *runtime.ClientOperation) {
		operation.Params = tagRequestParams{
			inner:  operation.Params,
			fields: fields,
		}
	}
}

type tagReadReader struct {
	inner  runtime.ClientResponseReader
	fields *tagAPIFields
}

func (r tagReadReader) ReadResponse(
	response runtime.ClientResponse,
	consumer runtime.Consumer,
) (any, error) {
	if response.Code() != 200 {
		return r.inner.ReadResponse(response, consumer)
	}

	payload := new(tagAPIModel)
	if err := consumer.Consume(response.Body(), payload); err != nil && err != io.EOF {
		return nil, err
	}

	*r.fields = payload.tagAPIFields
	result := extras.NewExtrasTagsReadOK()
	result.Payload = &payload.Tag
	return result, nil
}

func captureTagReadAPIFields(fields *tagAPIFields) extras.ClientOption {
	return func(operation *runtime.ClientOperation) {
		operation.Reader = tagReadReader{
			inner:  operation.Reader,
			fields: fields,
		}
	}
}

type tagListPayload struct {
	Count    *int64         `json:"count"`
	Next     *strfmt.URI    `json:"next,omitempty"`
	Previous *strfmt.URI    `json:"previous,omitempty"`
	Results  []*tagAPIModel `json:"results"`
}

type tagListReader struct {
	inner  runtime.ClientResponseReader
	fields map[int64]tagAPIFields
}

func (r tagListReader) ReadResponse(
	response runtime.ClientResponse,
	consumer runtime.Consumer,
) (any, error) {
	if response.Code() != 200 {
		return r.inner.ReadResponse(response, consumer)
	}

	payload := new(tagListPayload)
	if err := consumer.Consume(response.Body(), payload); err != nil && err != io.EOF {
		return nil, err
	}

	tags := make([]*models.Tag, 0, len(payload.Results))
	for _, tag := range payload.Results {
		if tag == nil {
			continue
		}
		tags = append(tags, &tag.Tag)
		r.fields[tag.ID] = tag.tagAPIFields
	}

	result := extras.NewExtrasTagsListOK()
	result.Payload = &extras.ExtrasTagsListOKBody{
		Count:    payload.Count,
		Next:     payload.Next,
		Previous: payload.Previous,
		Results:  tags,
	}
	return result, nil
}

func captureTagListAPIFields(
	fields map[int64]tagAPIFields,
) extras.ClientOption {
	return func(operation *runtime.ClientOperation) {
		operation.Reader = tagListReader{
			inner:  operation.Reader,
			fields: fields,
		}
	}
}
