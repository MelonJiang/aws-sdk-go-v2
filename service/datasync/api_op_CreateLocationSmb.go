// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines a file system on a Server Message Block (SMB) server that can be read
// from or written to.
func (c *Client) CreateLocationSmb(ctx context.Context, params *CreateLocationSmbInput, optFns ...func(*Options)) (*CreateLocationSmbOutput, error) {
	if params == nil {
		params = &CreateLocationSmbInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationSmb", params, optFns, c.addOperationCreateLocationSmbMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationSmbOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationSmbRequest
type CreateLocationSmbInput struct {

	// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block
	// (SMB) location.
	//
	// This member is required.
	AgentArns []string

	// The password of the user who can mount the share, has the permissions to access
	// files and folders in the SMB share.
	//
	// This member is required.
	Password *string

	// The name of the SMB server. This value is the IP address or Domain Name Service
	// (DNS) name of the SMB server. An agent that is installed on-premises uses this
	// hostname to mount the SMB server in a network. This name must either be
	// DNS-compliant or must be an IP version 4 (IPv4) address.
	//
	// This member is required.
	ServerHostname *string

	// The subdirectory in the SMB file system that is used to read data from the SMB
	// source location or write data to the SMB destination. The SMB path should be a
	// path that's exported by the SMB server, or a subdirectory of that path. The path
	// should be such that it can be mounted by other SMB clients in your network.
	// Subdirectory must be specified with forward slashes. For example,
	// /path/to/folder. To transfer all the data in the folder you specified, DataSync
	// needs to have permissions to mount the SMB share, as well as to access all the
	// data in that share. To ensure this, either ensure that the user/password
	// specified belongs to the user who can mount the share, and who has the
	// appropriate permissions for all of the files and directories that you want
	// DataSync to access, or use credentials of a member of the Backup Operators group
	// to mount the share. Doing either enables the agent to access the data. For the
	// agent to access directories, you must additionally enable all execute access.
	//
	// This member is required.
	Subdirectory *string

	// The user who can mount the share, has the permissions to access files and
	// folders in the SMB share. For information about choosing a user name that
	// ensures sufficient permissions to files, folders, and metadata, see the User
	// setting for SMB locations.
	//
	// This member is required.
	User *string

	// The name of the Windows domain that the SMB server belongs to.
	Domain *string

	// The mount options used by DataSync to access the SMB server.
	MountOptions *types.SmbMountOptions

	// The key-value pair that represents the tag that you want to add to the location.
	// The value can be an empty string. We recommend using tags to name your
	// resources.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationSmbResponse
type CreateLocationSmbOutput struct {

	// The Amazon Resource Name (ARN) of the source SMB file system location that is
	// created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationSmbMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateLocationSmbValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationSmb(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateLocationSmb(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationSmb",
	}
}
