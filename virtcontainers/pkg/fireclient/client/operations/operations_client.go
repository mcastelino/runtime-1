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
GetMachineConfig gets the machine configuration of the VM

Gets the machine configuration of the VM. When called before the PUT operation, it will return the default values for the vCPU count (=1), memory size (=128 MiB). By default Hyperthreading is disabled and there is no CPU Template.
*/
func (a *Client) GetMachineConfig(params *GetMachineConfigParams) (*GetMachineConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachineConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMachineConfig",
		Method:             "GET",
		PathPattern:        "/machine-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMachineConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMachineConfigOK), nil

}

/*
GetMmds gets the m m d s data store
*/
func (a *Client) GetMmds(params *GetMmdsParams) (*GetMmdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMmdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMmds",
		Method:             "GET",
		PathPattern:        "/mmds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMmdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMmdsOK), nil

}

/*
PatchMmds updates the m m d s data store
*/
func (a *Client) PatchMmds(params *PatchMmdsParams) (*PatchMmdsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchMmdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchMmds",
		Method:             "PATCH",
		PathPattern:        "/mmds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchMmdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchMmdsNoContent), nil

}

/*
PutMmds creates a m m d s microvm metadata service data store
*/
func (a *Client) PutMmds(params *PutMmdsParams) (*PutMmdsCreated, *PutMmdsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMmdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutMmds",
		Method:             "PUT",
		PathPattern:        "/mmds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutMmdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutMmdsCreated:
		return value, nil, nil
	case *PutMmdsNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
CreateSyncAction creates a synchronous action
*/
func (a *Client) CreateSyncAction(params *CreateSyncActionParams) (*CreateSyncActionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSyncActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSyncAction",
		Method:             "PUT",
		PathPattern:        "/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSyncActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSyncActionNoContent), nil

}

/*
DescribeInstance returns general information about an instance
*/
func (a *Client) DescribeInstance(params *DescribeInstanceParams) (*DescribeInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "describeInstance",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeInstanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeInstanceOK), nil

}

/*
PatchGuestDriveByID updates the properties of a drive

Updates the properties of the drive with the ID specified by drive_id path parameter. Will fail if update is not possible.
*/
func (a *Client) PatchGuestDriveByID(params *PatchGuestDriveByIDParams) (*PatchGuestDriveByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchGuestDriveByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchGuestDriveByID",
		Method:             "PATCH",
		PathPattern:        "/drives/{drive_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchGuestDriveByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchGuestDriveByIDNoContent), nil

}

/*
PutGuestBootSource creates or updates the boot source

Creates new boot source if one does not already exist, otherwise updates it. Will fail if update is not possible. Note that the only currently supported boot source is LocalImage.
*/
func (a *Client) PutGuestBootSource(params *PutGuestBootSourceParams) (*PutGuestBootSourceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGuestBootSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGuestBootSource",
		Method:             "PUT",
		PathPattern:        "/boot-source",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutGuestBootSourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutGuestBootSourceNoContent), nil

}

/*
PutGuestDriveByID creates or updates a drive

Creates new drive with ID specified by drive_id path parameter. If a drive with the specified ID already exists, updates its state based on new input. Will fail if update is not possible.
*/
func (a *Client) PutGuestDriveByID(params *PutGuestDriveByIDParams) (*PutGuestDriveByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGuestDriveByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGuestDriveByID",
		Method:             "PUT",
		PathPattern:        "/drives/{drive_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutGuestDriveByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutGuestDriveByIDNoContent), nil

}

/*
PutGuestNetworkInterfaceByID creates a network interface

Creates new network interface with ID specified by iface_id path parameter. Updating existing interfaces is currently not allowed.
*/
func (a *Client) PutGuestNetworkInterfaceByID(params *PutGuestNetworkInterfaceByIDParams) (*PutGuestNetworkInterfaceByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGuestNetworkInterfaceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGuestNetworkInterfaceByID",
		Method:             "PUT",
		PathPattern:        "/network-interfaces/{iface_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutGuestNetworkInterfaceByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutGuestNetworkInterfaceByIDNoContent), nil

}

/*
PutGuestVsockByID creates new vsock with ID specified by the id parameter

If the vsock device with the specified ID already exists, its body will be updated based on the new input. May fail if update is not possible.
*/
func (a *Client) PutGuestVsockByID(params *PutGuestVsockByIDParams) (*PutGuestVsockByIDCreated, *PutGuestVsockByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGuestVsockByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGuestVsockByID",
		Method:             "PUT",
		PathPattern:        "/vsocks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutGuestVsockByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutGuestVsockByIDCreated:
		return value, nil, nil
	case *PutGuestVsockByIDNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PutLogger initializes the logger by specifying two named pipes i e for the logs and metrics output
*/
func (a *Client) PutLogger(params *PutLoggerParams) (*PutLoggerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLoggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putLogger",
		Method:             "PUT",
		PathPattern:        "/logger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutLoggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutLoggerNoContent), nil

}

/*
PutMachineConfiguration updates the machine configuration of the VM

Updates the Virtual Machine Configuration with the specified input. Firecracker starts with default values for vCPU count (=1) and memory size (=128 MiB). With Hyperthreading enabled, the vCPU count is restricted to be 1 or an even number, otherwise there are no restrictions regarding the vCPU count. If any of the parameters has an incorrect value, the whole update fails.
*/
func (a *Client) PutMachineConfiguration(params *PutMachineConfigurationParams) (*PutMachineConfigurationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMachineConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putMachineConfiguration",
		Method:             "PUT",
		PathPattern:        "/machine-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutMachineConfigurationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutMachineConfigurationNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
