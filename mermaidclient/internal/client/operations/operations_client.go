// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteClusterClusterID deleteCluster Description
*/
func (a *Client) DeleteClusterClusterID(params *DeleteClusterClusterIDParams) (*DeleteClusterClusterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterID",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterClusterIDOK), nil

}

/*
DeleteClusterClusterIDRepairConfig deleteConfig cluster description
*/
func (a *Client) DeleteClusterClusterIDRepairConfig(params *DeleteClusterClusterIDRepairConfigParams) (*DeleteClusterClusterIDRepairConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDRepairConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterIDRepairConfig",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}/repair/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDRepairConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterClusterIDRepairConfigOK), nil

}

/*
DeleteClusterClusterIDRepairConfigConfigTypeExternalID deleteConfig cluster description
*/
func (a *Client) DeleteClusterClusterIDRepairConfigConfigTypeExternalID(params *DeleteClusterClusterIDRepairConfigConfigTypeExternalIDParams) (*DeleteClusterClusterIDRepairConfigConfigTypeExternalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDRepairConfigConfigTypeExternalIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterIDRepairConfigConfigTypeExternalID",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}/repair/config/{config_type}/{external_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDRepairConfigConfigTypeExternalIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterClusterIDRepairConfigConfigTypeExternalIDOK), nil

}

/*
DeleteClusterClusterIDRepairUnitUnitID deleteUnit Description
*/
func (a *Client) DeleteClusterClusterIDRepairUnitUnitID(params *DeleteClusterClusterIDRepairUnitUnitIDParams) (*DeleteClusterClusterIDRepairUnitUnitIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDRepairUnitUnitIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterIDRepairUnitUnitID",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}/repair/unit/{unit_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDRepairUnitUnitIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterClusterIDRepairUnitUnitIDOK), nil

}

/*
DeleteClusterClusterIDTaskTaskTypeTaskID deleteTask Description
*/
func (a *Client) DeleteClusterClusterIDTaskTaskTypeTaskID(params *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) (*DeleteClusterClusterIDTaskTaskTypeTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDTaskTaskTypeTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterIDTaskTaskTypeTaskID",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDTaskTaskTypeTaskIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterClusterIDTaskTaskTypeTaskIDOK), nil

}

/*
GetClusterClusterID loadCluster Description
*/
func (a *Client) GetClusterClusterID(params *GetClusterClusterIDParams) (*GetClusterClusterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDOK), nil

}

/*
GetClusterClusterIDRepairConfig getConfig cluster description
*/
func (a *Client) GetClusterClusterIDRepairConfig(params *GetClusterClusterIDRepairConfigParams) (*GetClusterClusterIDRepairConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDRepairConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDRepairConfig",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/repair/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDRepairConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDRepairConfigOK), nil

}

/*
GetClusterClusterIDRepairConfigConfigTypeExternalID getConfig cluster description
*/
func (a *Client) GetClusterClusterIDRepairConfigConfigTypeExternalID(params *GetClusterClusterIDRepairConfigConfigTypeExternalIDParams) (*GetClusterClusterIDRepairConfigConfigTypeExternalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDRepairConfigConfigTypeExternalIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDRepairConfigConfigTypeExternalID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/repair/config/{config_type}/{external_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDRepairConfigConfigTypeExternalIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDRepairConfigConfigTypeExternalIDOK), nil

}

/*
GetClusterClusterIDRepairUnitUnitID loadUnit Description
*/
func (a *Client) GetClusterClusterIDRepairUnitUnitID(params *GetClusterClusterIDRepairUnitUnitIDParams) (*GetClusterClusterIDRepairUnitUnitIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDRepairUnitUnitIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDRepairUnitUnitID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/repair/unit/{unit_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDRepairUnitUnitIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDRepairUnitUnitIDOK), nil

}

/*
GetClusterClusterIDRepairUnitUnitIDProgress get cluster cluster ID repair unit unit ID progress API
*/
func (a *Client) GetClusterClusterIDRepairUnitUnitIDProgress(params *GetClusterClusterIDRepairUnitUnitIDProgressParams) (*GetClusterClusterIDRepairUnitUnitIDProgressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDRepairUnitUnitIDProgressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDRepairUnitUnitIDProgress",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/repair/unit/{unit_id}/progress",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDRepairUnitUnitIDProgressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDRepairUnitUnitIDProgressOK), nil

}

/*
GetClusterClusterIDRepairUnits listUnits Description
*/
func (a *Client) GetClusterClusterIDRepairUnits(params *GetClusterClusterIDRepairUnitsParams) (*GetClusterClusterIDRepairUnitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDRepairUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDRepairUnits",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/repair/units",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDRepairUnitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDRepairUnitsOK), nil

}

/*
GetClusterClusterIDTaskTaskTypeTaskID loadTask Description
*/
func (a *Client) GetClusterClusterIDTaskTaskTypeTaskID(params *GetClusterClusterIDTaskTaskTypeTaskIDParams) (*GetClusterClusterIDTaskTaskTypeTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTaskTaskTypeTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTaskTaskTypeTaskID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTaskTaskTypeTaskIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDTaskTaskTypeTaskIDOK), nil

}

/*
GetClusterClusterIDTasks listTasks Description
*/
func (a *Client) GetClusterClusterIDTasks(params *GetClusterClusterIDTasksParams) (*GetClusterClusterIDTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTasks",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterClusterIDTasksOK), nil

}

/*
GetClusters listClusters Description
*/
func (a *Client) GetClusters(params *GetClustersParams) (*GetClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClustersOK), nil

}

/*
GetVersion returns the server version
*/
func (a *Client) GetVersion(params *GetVersionParams) (*GetVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersionOK), nil

}

/*
PostClusterClusterIDRepairUnits createUnit Description
*/
func (a *Client) PostClusterClusterIDRepairUnits(params *PostClusterClusterIDRepairUnitsParams) (*PostClusterClusterIDRepairUnitsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterClusterIDRepairUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClusterClusterIDRepairUnits",
		Method:             "POST",
		PathPattern:        "/cluster/{cluster_id}/repair/units",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClusterClusterIDRepairUnitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClusterClusterIDRepairUnitsCreated), nil

}

/*
PostClusterClusterIDTasks createTask Description
*/
func (a *Client) PostClusterClusterIDTasks(params *PostClusterClusterIDTasksParams) (*PostClusterClusterIDTasksCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterClusterIDTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClusterClusterIDTasks",
		Method:             "POST",
		PathPattern:        "/cluster/{cluster_id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClusterClusterIDTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClusterClusterIDTasksCreated), nil

}

/*
PostClusters createCluster Description
*/
func (a *Client) PostClusters(params *PostClustersParams) (*PostClustersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClusters",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClustersCreated), nil

}

/*
PutClusterClusterID updateCluster Description
*/
func (a *Client) PutClusterClusterID(params *PutClusterClusterIDParams) (*PutClusterClusterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterID",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDOK), nil

}

/*
PutClusterClusterIDRepairConfig updateConfig cluster description
*/
func (a *Client) PutClusterClusterIDRepairConfig(params *PutClusterClusterIDRepairConfigParams) (*PutClusterClusterIDRepairConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDRepairConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDRepairConfig",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/repair/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDRepairConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDRepairConfigOK), nil

}

/*
PutClusterClusterIDRepairConfigConfigTypeExternalID updateConfig cluster description
*/
func (a *Client) PutClusterClusterIDRepairConfigConfigTypeExternalID(params *PutClusterClusterIDRepairConfigConfigTypeExternalIDParams) (*PutClusterClusterIDRepairConfigConfigTypeExternalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDRepairConfigConfigTypeExternalIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDRepairConfigConfigTypeExternalID",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/repair/config/{config_type}/{external_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDRepairConfigConfigTypeExternalIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDRepairConfigConfigTypeExternalIDOK), nil

}

/*
PutClusterClusterIDRepairUnitUnitID updateUnit Description
*/
func (a *Client) PutClusterClusterIDRepairUnitUnitID(params *PutClusterClusterIDRepairUnitUnitIDParams) (*PutClusterClusterIDRepairUnitUnitIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDRepairUnitUnitIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDRepairUnitUnitID",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/repair/unit/{unit_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDRepairUnitUnitIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDRepairUnitUnitIDOK), nil

}

/*
PutClusterClusterIDTaskTaskTypeTaskID updateTask Description
*/
func (a *Client) PutClusterClusterIDTaskTaskTypeTaskID(params *PutClusterClusterIDTaskTaskTypeTaskIDParams) (*PutClusterClusterIDTaskTaskTypeTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTaskTaskTypeTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTaskTaskTypeTaskID",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTaskTaskTypeTaskIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDTaskTaskTypeTaskIDOK), nil

}

/*
PutClusterClusterIDTaskTaskTypeTaskIDStart startTask Description
*/
func (a *Client) PutClusterClusterIDTaskTaskTypeTaskIDStart(params *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) (*PutClusterClusterIDTaskTaskTypeTaskIDStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTaskTaskTypeTaskIDStartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTaskTaskTypeTaskIDStart",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTaskTaskTypeTaskIDStartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDTaskTaskTypeTaskIDStartOK), nil

}

/*
PutClusterClusterIDTaskTaskTypeTaskIDStop stopTask Description
*/
func (a *Client) PutClusterClusterIDTaskTaskTypeTaskIDStop(params *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) (*PutClusterClusterIDTaskTaskTypeTaskIDStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTaskTaskTypeTaskIDStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTaskTaskTypeTaskIDStop",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTaskTaskTypeTaskIDStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClusterClusterIDTaskTaskTypeTaskIDStopOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
