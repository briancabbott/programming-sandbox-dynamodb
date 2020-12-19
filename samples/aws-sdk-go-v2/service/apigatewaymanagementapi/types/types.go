// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Identity struct {

	// The source IP address of the TCP connection making the request to API Gateway.
	//
	// This member is required.
	SourceIp *string

	// The User Agent of the API caller.
	//
	// This member is required.
	UserAgent *string
}
