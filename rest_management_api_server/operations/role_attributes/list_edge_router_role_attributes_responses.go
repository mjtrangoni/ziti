// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package role_attributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ListEdgeRouterRoleAttributesOKCode is the HTTP code returned for type ListEdgeRouterRoleAttributesOK
const ListEdgeRouterRoleAttributesOKCode int = 200

/*ListEdgeRouterRoleAttributesOK A list of role attributes

swagger:response listEdgeRouterRoleAttributesOK
*/
type ListEdgeRouterRoleAttributesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListRoleAttributesEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterRoleAttributesOK creates ListEdgeRouterRoleAttributesOK with default headers values
func NewListEdgeRouterRoleAttributesOK() *ListEdgeRouterRoleAttributesOK {

	return &ListEdgeRouterRoleAttributesOK{}
}

// WithPayload adds the payload to the list edge router role attributes o k response
func (o *ListEdgeRouterRoleAttributesOK) WithPayload(payload *rest_model.ListRoleAttributesEnvelope) *ListEdgeRouterRoleAttributesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router role attributes o k response
func (o *ListEdgeRouterRoleAttributesOK) SetPayload(payload *rest_model.ListRoleAttributesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterRoleAttributesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterRoleAttributesBadRequestCode is the HTTP code returned for type ListEdgeRouterRoleAttributesBadRequest
const ListEdgeRouterRoleAttributesBadRequestCode int = 400

/*ListEdgeRouterRoleAttributesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listEdgeRouterRoleAttributesBadRequest
*/
type ListEdgeRouterRoleAttributesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterRoleAttributesBadRequest creates ListEdgeRouterRoleAttributesBadRequest with default headers values
func NewListEdgeRouterRoleAttributesBadRequest() *ListEdgeRouterRoleAttributesBadRequest {

	return &ListEdgeRouterRoleAttributesBadRequest{}
}

// WithPayload adds the payload to the list edge router role attributes bad request response
func (o *ListEdgeRouterRoleAttributesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterRoleAttributesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router role attributes bad request response
func (o *ListEdgeRouterRoleAttributesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterRoleAttributesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterRoleAttributesUnauthorizedCode is the HTTP code returned for type ListEdgeRouterRoleAttributesUnauthorized
const ListEdgeRouterRoleAttributesUnauthorizedCode int = 401

/*ListEdgeRouterRoleAttributesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listEdgeRouterRoleAttributesUnauthorized
*/
type ListEdgeRouterRoleAttributesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterRoleAttributesUnauthorized creates ListEdgeRouterRoleAttributesUnauthorized with default headers values
func NewListEdgeRouterRoleAttributesUnauthorized() *ListEdgeRouterRoleAttributesUnauthorized {

	return &ListEdgeRouterRoleAttributesUnauthorized{}
}

// WithPayload adds the payload to the list edge router role attributes unauthorized response
func (o *ListEdgeRouterRoleAttributesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterRoleAttributesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router role attributes unauthorized response
func (o *ListEdgeRouterRoleAttributesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterRoleAttributesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
