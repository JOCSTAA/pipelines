// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new run service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for run service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ArchiveRun archives a run in an experiment given by run ID and experiment ID
*/
func (a *Client) ArchiveRun(params *ArchiveRunParams, authInfo runtime.ClientAuthInfoWriter) (*ArchiveRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ArchiveRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs/{run_id}:archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ArchiveRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ArchiveRunOK), nil

}

/*
CreateRun creates a new run in an experiment specified by experiment ID if experiment ID is not specified the run is created in the default experiment
*/
func (a *Client) CreateRun(params *CreateRunParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRunOK), nil

}

/*
DeleteRun deletes a run in an experiment given by run ID and experiment ID
*/
func (a *Client) DeleteRun(params *DeleteRunParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRun",
		Method:             "DELETE",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRunOK), nil

}

/*
GetRun finds a specific run by ID
*/
func (a *Client) GetRun(params *GetRunParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRun",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunOK), nil

}

/*
ListRuns finds all runs in an experiment given by experiment ID if experiment id is not specified finds all runs across all experiments
*/
func (a *Client) ListRuns(params *ListRunsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRuns",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRunsOK), nil

}

/*
ReadArtifact finds artifact data in a run
*/
func (a *Client) ReadArtifact(params *ReadArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*ReadArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadArtifact",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs/{run_id}/nodes/{node_id}/artifacts/{artifact_name}:read",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadArtifactOK), nil

}

/*
TerminateRun terminates an active run
*/
func (a *Client) TerminateRun(params *TerminateRunParams, authInfo runtime.ClientAuthInfoWriter) (*TerminateRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TerminateRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs/{run_id}:terminate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TerminateRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TerminateRunOK), nil

}

/*
UnarchiveRun restores an archived run in an experiment given by run ID and experiment ID
*/
func (a *Client) UnarchiveRun(params *UnarchiveRunParams, authInfo runtime.ClientAuthInfoWriter) (*UnarchiveRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnarchiveRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UnarchiveRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/experiments/{experiment_id}/runs/{run_id}:unarchive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnarchiveRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UnarchiveRunOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
