// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package util

import (
	"fmt"
)

// ErrCode Error code type. This type is used by union-jd-go SDK for error codes returned externally
type ErrCode int32

const (
	// BaseIndexErrCode base index error code
	BaseIndexErrCode = 1000
	// ErrCodeCount     = ErrCodeMax%BaseIndexErrCode + 1
)

const (
	// ErrCodeSuccess no error occurred
	ErrCodeSuccess ErrCode = 0
	// ErrCodeUnknown unknown error
	ErrCodeUnknown ErrCode = BaseIndexErrCode
	// ErrCodeAPIInvalidArgument the api parameter is invalid procedure
	ErrCodeAPIInvalidArgument ErrCode = BaseIndexErrCode + 1
	// ErrCodeAPIInvalidConfig an invalid error code was configured procedure
	ErrCodeAPIInvalidConfig ErrCode = BaseIndexErrCode + 2
	// ErrCodePluginError plugin error error code
	ErrCodePluginError ErrCode = BaseIndexErrCode + 3
	// ErrCodeAPITimeoutError api timeout error error code
	ErrCodeAPITimeoutError ErrCode = BaseIndexErrCode + 4
	// ErrCodeInvalidStateError After the SDK has destroyed, continue to tune API error codes
	ErrCodeInvalidStateError ErrCode = BaseIndexErrCode + 5
	// ErrCodeServerUserError When the server is connected, the server returns error 400
	ErrCodeServerUserError ErrCode = BaseIndexErrCode + 6
	// ErrCodeNetworkError An unknown network exception occurred when connecting to the server. Procedure
	ErrCodeNetworkError ErrCode = BaseIndexErrCode + 7
	// ErrCodeCircuitBreakerError service circuit breaker error
	ErrCodeCircuitBreakerError ErrCode = BaseIndexErrCode + 8
	// ErrCodeInstanceInfoError The instance information is incorrect, for example, the service weight information is empty
	ErrCodeInstanceInfoError ErrCode = BaseIndexErrCode + 9
	// ErrCodeAPIInstanceNotFound description failed to obtain the service instance
	ErrCodeAPIInstanceNotFound ErrCode = BaseIndexErrCode + 10
	// ErrCodeInvalidRule the routing rule is invalid
	ErrCodeInvalidRule ErrCode = BaseIndexErrCode + 11
	// ErrCodeRouteRuleNotMatch failed to match the routing rule procedure
	ErrCodeRouteRuleNotMatch ErrCode = BaseIndexErrCode + 12
	// ErrCodeInvalidResponse the message returned by the server is invalid
	ErrCodeInvalidResponse ErrCode = BaseIndexErrCode + 13
	// ErrCodeInternalError internal algorithm and system errors
	ErrCodeInternalError ErrCode = BaseIndexErrCode + 14
	// ErrCodeServiceNotFound service does not exist
	ErrCodeServiceNotFound ErrCode = BaseIndexErrCode + 15
	// ErrCodeServerException the server returns error 500
	ErrCodeServerException ErrCode = BaseIndexErrCode + 16
	// ErrCodeLocationNotFound description failed to obtain region information
	ErrCodeLocationNotFound ErrCode = BaseIndexErrCode + 17
	// ErrCodeLocationMismatch The nearest route failed because there was no instance on the nearest level. Procedure
	ErrCodeLocationMismatch ErrCode = BaseIndexErrCode + 18
	// ErrCodeDstMetaMismatch Description Failed to filter the metadata of the target rule
	ErrCodeDstMetaMismatch ErrCode = BaseIndexErrCode + 19
	// ErrCodeMeshConfigNotFound grid rule of target type not found
	ErrCodeMeshConfigNotFound ErrCode = BaseIndexErrCode + 20
	// ErrCodeConsumerInitCalleeError Description Failed to initialize the called service
	ErrCodeConsumerInitCalleeError ErrCode = BaseIndexErrCode + 21
	// ErrCodeCount Number of interface error codes. Increase this number by 1 for each error code added
	ErrCodeCount = 23
)

const (
	// BaseServerErrCode base server error code
	BaseServerErrCode = 2000
	// ErrCodeConnectError connection error
	ErrCodeConnectError ErrCode = BaseServerErrCode + 1
	// ErrCodeServerError server error
	ErrCodeServerError ErrCode = BaseServerErrCode + 2
	// ErrorCodeRPCError rpc error
	ErrorCodeRPCError ErrCode = BaseServerErrCode + 3
	// ErrorCodeRPCTimeoutErr rpc timeout
	ErrorCodeRPCTimeoutErr ErrCode = BaseServerErrCode + 4
	// ErrCodeInvalidServerResponse invalid server response
	ErrCodeInvalidServerResponse ErrCode = BaseServerErrCode + 5
	// ErrCodeInvalidRequest invalid request
	ErrCodeInvalidRequest ErrCode = BaseServerErrCode + 6
	// ErrCodeUnauthorized unauthorized
	ErrCodeUnauthorized ErrCode = BaseServerErrCode + 7
	// ErrCodeRequestLimit request limit
	ErrCodeRequestLimit ErrCode = BaseServerErrCode + 8
	// ErrCodeCmdbNotFound cmdb not found
	ErrCodeCmdbNotFound ErrCode = BaseServerErrCode + 9
	// ErrCodeUnknownServerError unknown server error
	ErrCodeUnknownServerError ErrCode = BaseServerErrCode + 100
)

const (
	// BaseSDKInternalErrCode base sdk internal error code
	BaseSDKInternalErrCode = 3000
	// ErrCodeDiskError disk error
	ErrCodeDiskError = BaseSDKInternalErrCode + 1
)

// Retryable retry error code
func (e ErrCode) Retryable() bool {
	return e == ErrCodeNetworkError || e == ErrCodeServerException
}

// SDKError the total wrapper type of the SDK error
type SDKError interface {
	// ErrorCode get error code
	ErrorCode() ErrCode
	// Error getting error messages
	Error() string
	// ServerCode gets the error code returned by the server
	ServerCode() uint32
	// ServerInfo gets the error message returned by the server
	ServerInfo() string
}

// sdkError SDK error type implementation
type sdkError struct {
	errCode    ErrCode
	errDetail  string
	cause      error
	serverCode uint32
	serverInfo string
}

// ErrorCode get error code
func (s *sdkError) ErrorCode() ErrCode {
	return s.errCode
}

// Error getting error messages
func (s *sdkError) Error() string {
	errCodeStr := ErrCodeToString(s.ErrorCode())
	if nil != s.cause {
		return fmt.Sprintf(
			"UNION-JD-%v(%s): %s, cause: %s", s.ErrorCode(), errCodeStr, s.errDetail, s.cause.Error())
	}
	return fmt.Sprintf("UNION-JD-%v(%s): %s", s.ErrorCode(), errCodeStr, s.errDetail)
}

// ServerCode server return code
func (s *sdkError) ServerCode() uint32 {
	return s.serverCode
}

// ServerInfo the server returns information
func (s *sdkError) ServerInfo() string {
	return s.serverInfo
}

// String output string information
func (s *sdkError) String() string {
	return s.Error()
}

// NewSDKError SDK error related class builder
func NewSDKError(errCode ErrCode, cause error, msg string, args ...interface{}) SDKError {
	var errDetail = fmt.Sprintf(msg, args...)
	return &sdkError{
		errCode:   errCode,
		errDetail: errDetail,
		cause:     cause}
}

// NewSDKErrorWithServerInfo sdk error related class builder
func NewSDKErrorWithServerInfo(errCode ErrCode, cause error, serverCode uint32, serverInfo string, msg string, args ...interface{}) SDKError {
	return &sdkError{
		errCode:    errCode,
		errDetail:  fmt.Sprintf(msg, args...),
		cause:      cause,
		serverCode: serverCode,
		serverInfo: serverInfo,
	}
}

const (
	// RetCodeDivFactor the return code takes the base of the module
	RetCodeDivFactor = 1000
	// SuccessRetCode successful error code after modulo
	SuccessRetCode = 200
	// ServerExceptionRetCode an internal error occurred in the server after the module was taken. Procedure
	ServerExceptionRetCode = 500
)

// IsSuccessResultCode error code to determine success
func IsSuccessResultCode(retCode uint32) bool {
	return retCode/RetCodeDivFactor == SuccessRetCode
}

// IsServerException determine whether an internal server error occurs
func IsServerException(retCode uint32) bool {
	return retCode/RetCodeDivFactor == ServerExceptionRetCode
}

// NewServerSDKError construct server related errors
func NewServerSDKError(serverCode uint32, serverInfo string, cause error, msg string, args ...interface{}) SDKError {
	return &sdkError{
		errCode:    serverCodeToErrCode(serverCode),
		errDetail:  fmt.Sprintf(msg, args...),
		cause:      cause,
		serverCode: serverCode,
		serverInfo: serverInfo,
	}
}

// serverCodeToErrCode The server error code is converted to an SDK error code
func serverCodeToErrCode(retCode uint32) ErrCode {
	errCode := ErrCodeServerUserError
	if IsServerException(retCode) {
		errCode = ErrCodeServerException
	}
	return errCode
}

var errCodeString = map[ErrCode]string{
	ErrCodeSuccess:                 "Success",
	ErrCodeUnknown:                 "ErrCodeUnknown",
	ErrCodeAPIInvalidArgument:      "ErrCodeAPIInvalidArgument",
	ErrCodeAPIInvalidConfig:        "ErrCodeAPIInvalidConfig",
	ErrCodePluginError:             "ErrCodePluginError",
	ErrCodeAPITimeoutError:         "ErrCodeAPITimeoutError",
	ErrCodeInvalidStateError:       "ErrCodeInvalidStateError",
	ErrCodeServerUserError:         "ErrCodeServerUserError",
	ErrCodeNetworkError:            "ErrCodeNetworkError",
	ErrCodeCircuitBreakerError:     "ErrCodeCircuitBreakerError",
	ErrCodeInstanceInfoError:       "ErrCodeInstanceInfoError",
	ErrCodeAPIInstanceNotFound:     "ErrCodeAPIInstanceNotFound",
	ErrCodeInvalidRule:             "ErrCodeInvalidRule",
	ErrCodeRouteRuleNotMatch:       "ErrCodeRouteRuleNotMatch",
	ErrCodeInvalidResponse:         "ErrCodeInvalidResponse",
	ErrCodeServiceNotFound:         "ErrCodeServiceNotFound",
	ErrCodeInternalError:           "ErrCodeInternalError",
	ErrCodeServerException:         "ErrCodeServerException",
	ErrCodeLocationNotFound:        "ErrCodeLocationNotFound",
	ErrCodeLocationMismatch:        "ErrCodeLocationMismatch",
	ErrCodeDstMetaMismatch:         "ErrCodeDstMetaMismatch",
	ErrCodeMeshConfigNotFound:      "ErrCodeMeshConfigNotFound",
	ErrCodeConsumerInitCalleeError: "ErrCodeConsumerInitCalleeError",
}

var errCodeArray = []ErrCode{ErrCodeSuccess, ErrCodeUnknown, ErrCodeAPIInvalidArgument,
	ErrCodeAPIInvalidConfig, ErrCodePluginError, ErrCodeAPITimeoutError, ErrCodeInvalidStateError,
	ErrCodeServerUserError, ErrCodeNetworkError, ErrCodeCircuitBreakerError, ErrCodeInstanceInfoError,
	ErrCodeAPIInstanceNotFound, ErrCodeInvalidRule, ErrCodeRouteRuleNotMatch, ErrCodeInvalidResponse,
	ErrCodeInternalError, ErrCodeServiceNotFound, ErrCodeServerException, ErrCodeLocationNotFound,
	ErrCodeLocationMismatch, ErrCodeDstMetaMismatch, ErrCodeMeshConfigNotFound, ErrCodeConsumerInitCalleeError,
}

// ErrCodeFromIndex Returns an error code based on the error code index
func ErrCodeFromIndex(i int) ErrCode {
	return errCodeArray[i]
}

// ErrCodeToString converts an error code to a string
func ErrCodeToString(ec ErrCode) string {
	res, ok := errCodeString[ec]
	if !ok {
		return "ErrCodeUnknown"
	}
	return res
}

// ErrCodeType error code type
type ErrCodeType int

const (
	// UnionSDKError SDK system error
	UnionSDKError ErrCodeType = 0
	// UserError user error
	UserError ErrCodeType = 1
)

// mapping of error codes to error types
var errCodeTypeMap = map[ErrCode]ErrCodeType{
	ErrCodeUnknown:             UnionSDKError,
	ErrCodePluginError:         UnionSDKError,
	ErrCodeAPITimeoutError:     UnionSDKError,
	ErrCodeNetworkError:        UnionSDKError,
	ErrCodeInvalidResponse:     UnionSDKError,
	ErrCodeInternalError:       UnionSDKError,
	ErrCodeServerException:     UnionSDKError,
	ErrCodeCircuitBreakerError: UnionSDKError,
	ErrCodeLocationNotFound:    UnionSDKError,

	ErrCodeAPIInvalidArgument:      UserError,
	ErrCodeAPIInvalidConfig:        UserError,
	ErrCodeInvalidStateError:       UserError,
	ErrCodeServerUserError:         UserError,
	ErrCodeInstanceInfoError:       UserError,
	ErrCodeAPIInstanceNotFound:     UserError,
	ErrCodeInvalidRule:             UserError,
	ErrCodeServiceNotFound:         UserError,
	ErrCodeRouteRuleNotMatch:       UserError,
	ErrCodeLocationMismatch:        UserError,
	ErrCodeDstMetaMismatch:         UserError,
	ErrCodeMeshConfigNotFound:      UserError,
	ErrCodeConsumerInitCalleeError: UserError,
}

// GetErrCodeType gets the error code type
func GetErrCodeType(e ErrCode) ErrCodeType {
	t, ok := errCodeTypeMap[e]
	if ok {
		return t
	}
	return UnionSDKError
}
