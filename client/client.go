// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckUsedPropertyRequest struct {
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s CheckUsedPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyRequest) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyRequest) SetPropertyId(v int64) *CheckUsedPropertyRequest {
	s.PropertyId = &v
	return s
}

type CheckUsedPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UseCount  *int64  `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s CheckUsedPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyResponseBody) SetRequestId(v string) *CheckUsedPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsedPropertyResponseBody) SetUseCount(v int64) *CheckUsedPropertyResponseBody {
	s.UseCount = &v
	return s
}

type CheckUsedPropertyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckUsedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUsedPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyResponse) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyResponse) SetHeaders(v map[string]*string) *CheckUsedPropertyResponse {
	s.Headers = v
	return s
}

func (s *CheckUsedPropertyResponse) SetStatusCode(v int32) *CheckUsedPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsedPropertyResponse) SetBody(v *CheckUsedPropertyResponseBody) *CheckUsedPropertyResponse {
	s.Body = v
	return s
}

type CheckUsedPropertyValueRequest struct {
	// CheckUsedPropertyValue
	PropertyId      *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s CheckUsedPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueRequest) SetPropertyId(v int64) *CheckUsedPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *CheckUsedPropertyValueRequest) SetPropertyValueId(v int64) *CheckUsedPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

type CheckUsedPropertyValueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UseCount  *int64  `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s CheckUsedPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueResponseBody) SetRequestId(v string) *CheckUsedPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsedPropertyValueResponseBody) SetUseCount(v int64) *CheckUsedPropertyValueResponseBody {
	s.UseCount = &v
	return s
}

type CheckUsedPropertyValueResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckUsedPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUsedPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueResponse) SetHeaders(v map[string]*string) *CheckUsedPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *CheckUsedPropertyValueResponse) SetStatusCode(v int32) *CheckUsedPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsedPropertyValueResponse) SetBody(v *CheckUsedPropertyValueResponseBody) *CheckUsedPropertyValueResponse {
	s.Body = v
	return s
}

type CreatePropertyRequest struct {
	PropertyKey    *string   `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	PropertyValues []*string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyRequest) GoString() string {
	return s.String()
}

func (s *CreatePropertyRequest) SetPropertyKey(v string) *CreatePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *CreatePropertyRequest) SetPropertyValues(v []*string) *CreatePropertyRequest {
	s.PropertyValues = v
	return s
}

type CreatePropertyResponseBody struct {
	CreateResult *CreatePropertyResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBody) SetCreateResult(v *CreatePropertyResponseBodyCreateResult) *CreatePropertyResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreatePropertyResponseBody) SetRequestId(v string) *CreatePropertyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePropertyResponseBodyCreateResult struct {
	PropertyId             *int64                                                        `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyKey            *string                                                       `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	SavePropertyValueModel *CreatePropertyResponseBodyCreateResultSavePropertyValueModel `json:"SavePropertyValueModel,omitempty" xml:"SavePropertyValueModel,omitempty" type:"Struct"`
}

func (s CreatePropertyResponseBodyCreateResult) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResult) SetPropertyId(v int64) *CreatePropertyResponseBodyCreateResult {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) SetPropertyKey(v string) *CreatePropertyResponseBodyCreateResult {
	s.PropertyKey = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) SetSavePropertyValueModel(v *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) *CreatePropertyResponseBodyCreateResult {
	s.SavePropertyValueModel = v
	return s
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModel struct {
	FailedPropertyValues []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues `json:"FailedPropertyValues,omitempty" xml:"FailedPropertyValues,omitempty" type:"Repeated"`
	SavePropertyValues   []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues   `json:"SavePropertyValues,omitempty" xml:"SavePropertyValues,omitempty" type:"Repeated"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModel) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModel) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) SetFailedPropertyValues(v []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	s.FailedPropertyValues = v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) SetSavePropertyValues(v []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	s.SavePropertyValues = v
	return s
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PropertyId    *int64  `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetErrorCode(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorCode = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetErrorMessage(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetPropertyId(v int64) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetPropertyValue(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyValue = &v
	return s
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues struct {
	PropertyValue   *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	PropertyValueId *int64  `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) SetPropertyValue(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) SetPropertyValueId(v int64) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValueId = &v
	return s
}

type CreatePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponse) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponse) SetHeaders(v map[string]*string) *CreatePropertyResponse {
	s.Headers = v
	return s
}

func (s *CreatePropertyResponse) SetStatusCode(v int32) *CreatePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePropertyResponse) SetBody(v *CreatePropertyResponseBody) *CreatePropertyResponse {
	s.Body = v
	return s
}

type CreateUsersRequest struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The remarks of the end user.
	Users []*CreateUsersRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersRequest) SetPassword(v string) *CreateUsersRequest {
	s.Password = &v
	return s
}

func (s *CreateUsersRequest) SetUsers(v []*CreateUsersRequestUsers) *CreateUsersRequest {
	s.Users = v
	return s
}

type CreateUsersRequestUsers struct {
	// The name of the end user.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Details of the convenience user that failed to be created.
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error message returned.
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The error code returned if the request failed.
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The mobile number of the end user.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The email address of the end user.
	Phone  *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateUsersRequestUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersRequestUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersRequestUsers) SetEmail(v string) *CreateUsersRequestUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersRequestUsers) SetEndUserId(v string) *CreateUsersRequestUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersRequestUsers) SetOrgId(v string) *CreateUsersRequestUsers {
	s.OrgId = &v
	return s
}

func (s *CreateUsersRequestUsers) SetOwnerType(v string) *CreateUsersRequestUsers {
	s.OwnerType = &v
	return s
}

func (s *CreateUsersRequestUsers) SetPassword(v string) *CreateUsersRequestUsers {
	s.Password = &v
	return s
}

func (s *CreateUsersRequestUsers) SetPhone(v string) *CreateUsersRequestUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersRequestUsers) SetRemark(v string) *CreateUsersRequestUsers {
	s.Remark = &v
	return s
}

type CreateUsersResponseBody struct {
	CreateResult *CreateUsersResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBody) SetCreateResult(v *CreateUsersResponseBodyCreateResult) *CreateUsersResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateUsersResponseBody) SetRequestId(v string) *CreateUsersResponseBody {
	s.RequestId = &v
	return s
}

type CreateUsersResponseBodyCreateResult struct {
	CreatedUsers []*CreateUsersResponseBodyCreateResultCreatedUsers `json:"CreatedUsers,omitempty" xml:"CreatedUsers,omitempty" type:"Repeated"`
	FailedUsers  []*CreateUsersResponseBodyCreateResultFailedUsers  `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
}

func (s CreateUsersResponseBodyCreateResult) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResult) SetCreatedUsers(v []*CreateUsersResponseBodyCreateResultCreatedUsers) *CreateUsersResponseBodyCreateResult {
	s.CreatedUsers = v
	return s
}

func (s *CreateUsersResponseBodyCreateResult) SetFailedUsers(v []*CreateUsersResponseBodyCreateResultFailedUsers) *CreateUsersResponseBodyCreateResult {
	s.FailedUsers = v
	return s
}

type CreateUsersResponseBodyCreateResultCreatedUsers struct {
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateUsersResponseBodyCreateResultCreatedUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResultCreatedUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetEmail(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetEndUserId(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetPhone(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetRemark(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Remark = &v
	return s
}

type CreateUsersResponseBodyCreateResultFailedUsers struct {
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s CreateUsersResponseBodyCreateResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResultFailedUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetEmail(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetEndUserId(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetErrorCode(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetErrorMessage(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetPhone(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.Phone = &v
	return s
}

type CreateUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponse) GoString() string {
	return s.String()
}

func (s *CreateUsersResponse) SetHeaders(v map[string]*string) *CreateUsersResponse {
	s.Headers = v
	return s
}

func (s *CreateUsersResponse) SetStatusCode(v int32) *CreateUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUsersResponse) SetBody(v *CreateUsersResponseBody) *CreateUsersResponse {
	s.Body = v
	return s
}

type DeleteUserPropertyValueRequest struct {
	// DeleteUserPropertyValue
	PropertyId      *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// Dissociates a user property from a user.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueRequest) SetPropertyId(v int64) *DeleteUserPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) SetPropertyValueId(v int64) *DeleteUserPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) SetUserId(v int64) *DeleteUserPropertyValueRequest {
	s.UserId = &v
	return s
}

type DeleteUserPropertyValueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueResponseBody) SetRequestId(v string) *DeleteUserPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserPropertyValueResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueResponse) SetHeaders(v map[string]*string) *DeleteUserPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPropertyValueResponse) SetStatusCode(v int32) *DeleteUserPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPropertyValueResponse) SetBody(v *DeleteUserPropertyValueResponseBody) *DeleteUserPropertyValueResponse {
	s.Body = v
	return s
}

type DescribeMfaDevicesRequest struct {
	// This parameter is unavailable.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The list of username of convenience users.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time when the virtual MFA device was enabled. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The time when a locked virtual MFA device is automatically unlocked. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SerialNumbers []*string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty" type:"Repeated"`
}

func (s DescribeMfaDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesRequest) SetEndUserIds(v []*string) *DescribeMfaDevicesRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeMfaDevicesRequest) SetMaxResults(v int64) *DescribeMfaDevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetNextToken(v string) *DescribeMfaDevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetSerialNumbers(v []*string) *DescribeMfaDevicesRequest {
	s.SerialNumbers = v
	return s
}

type DescribeMfaDevicesResponseBody struct {
	// The serial number of the virtual MFA device, which is a unique identifier.
	MfaDevices []*DescribeMfaDevicesResponseBodyMfaDevices `json:"MfaDevices,omitempty" xml:"MfaDevices,omitempty" type:"Repeated"`
	// The operation that you want to perform. Set the value to DescribeMfaDevices.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is unavailable.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMfaDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponseBody) SetMfaDevices(v []*DescribeMfaDevicesResponseBodyMfaDevices) *DescribeMfaDevicesResponseBody {
	s.MfaDevices = v
	return s
}

func (s *DescribeMfaDevicesResponseBody) SetNextToken(v string) *DescribeMfaDevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMfaDevicesResponseBody) SetRequestId(v string) *DescribeMfaDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMfaDevicesResponseBodyMfaDevices struct {
	// The ID of the request.
	ConsecutiveFails *int32  `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	DeviceType       *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The username of the convenience user that uses the virtual MFA device.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The types of the virtual MFA device. Set the value to TOTP_VIRTUAL, which indicates that the virtual MFA devices follow the Time-based One-time Password (TOTP) algorithm.
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The serial numbers of the virtual MFA devices.
	GmtEnabled *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
	// The serial number of the virtual MFA device, which is a unique identifier.
	GmtUnlock *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 500.
	//
	// Default value: 100.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Queries the information about virtual MFA devices that are bound to convenience users.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMfaDevicesResponseBodyMfaDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesResponseBodyMfaDevices) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetConsecutiveFails(v int32) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetDeviceType(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.DeviceType = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetEmail(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Email = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetEndUserId(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetGmtEnabled(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetGmtUnlock(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetId(v int64) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Id = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetSerialNumber(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetStatus(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Status = &v
	return s
}

type DescribeMfaDevicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMfaDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMfaDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponse) SetHeaders(v map[string]*string) *DescribeMfaDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMfaDevicesResponse) SetStatusCode(v int32) *DescribeMfaDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMfaDevicesResponse) SetBody(v *DescribeMfaDevicesResponseBody) *DescribeMfaDevicesResponse {
	s.Body = v
	return s
}

type DescribeUsersRequest struct {
	EndUserIds        []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	Filter            *string   `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults        *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrgId             *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s DescribeUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersRequest) SetEndUserIds(v []*string) *DescribeUsersRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeUsersRequest) SetExcludeEndUserIds(v []*string) *DescribeUsersRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *DescribeUsersRequest) SetFilter(v string) *DescribeUsersRequest {
	s.Filter = &v
	return s
}

func (s *DescribeUsersRequest) SetMaxResults(v int64) *DescribeUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersRequest) SetNextToken(v string) *DescribeUsersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersRequest) SetOrgId(v string) *DescribeUsersRequest {
	s.OrgId = &v
	return s
}

type DescribeUsersResponseBody struct {
	NextToken *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*DescribeUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBody) SetNextToken(v string) *DescribeUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersResponseBody) SetRequestId(v string) *DescribeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersResponseBody) SetUsers(v []*DescribeUsersResponseBodyUsers) *DescribeUsersResponseBody {
	s.Users = v
	return s
}

type DescribeUsersResponseBodyUsers struct {
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndUserId       *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsTenantManager *bool   `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	OrgId           *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OwnerType       *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	WyId            *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribeUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsers) SetEmail(v string) *DescribeUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetEndUserId(v string) *DescribeUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetId(v int64) *DescribeUsersResponseBodyUsers {
	s.Id = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetIsTenantManager(v bool) *DescribeUsersResponseBodyUsers {
	s.IsTenantManager = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOrgId(v string) *DescribeUsersResponseBodyUsers {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOwnerType(v string) *DescribeUsersResponseBodyUsers {
	s.OwnerType = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetPhone(v string) *DescribeUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetRemark(v string) *DescribeUsersResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetStatus(v int64) *DescribeUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetWyId(v string) *DescribeUsersResponseBodyUsers {
	s.WyId = &v
	return s
}

type DescribeUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponse) SetHeaders(v map[string]*string) *DescribeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersResponse) SetStatusCode(v int32) *DescribeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersResponse) SetBody(v *DescribeUsersResponseBody) *DescribeUsersResponse {
	s.Body = v
	return s
}

type FilterUsersRequest struct {
	// The IDs of excluded users.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that you enter for a fuzzy search. You can enter a string to match the username or email address.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return information about cloud desktops that are assigned to the convenience user.
	IncludeDesktopCount *bool `json:"IncludeDesktopCount,omitempty" xml:"IncludeDesktopCount,omitempty"`
	// Specifies whether to return the number of desktop groups that are assigned to the user.
	IncludeDesktopGroupCount *bool `json:"IncludeDesktopGroupCount,omitempty" xml:"IncludeDesktopGroupCount,omitempty"`
	// The number of entries to return on each page. If you set this parameter to a value greater than 100, the system resets the value to 100.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. You do not need to configure this parameter if you call this operation for the first time. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to perform the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The parameter that is supported to sort query results.
	OrderParam *FilterUsersRequestOrderParam `json:"OrderParam,omitempty" xml:"OrderParam,omitempty" type:"Struct"`
	// The ID of the organization.
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The type of the account ownership.
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// Details of the user property that you want to perform fuzzy search.
	PropertyFilterParam []*FilterUsersRequestPropertyFilterParam `json:"PropertyFilterParam,omitempty" xml:"PropertyFilterParam,omitempty" type:"Repeated"`
	// Details of the properties and property values.
	PropertyKeyValueFilterParam []*FilterUsersRequestPropertyKeyValueFilterParam `json:"PropertyKeyValueFilterParam,omitempty" xml:"PropertyKeyValueFilterParam,omitempty" type:"Repeated"`
}

func (s FilterUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequest) GoString() string {
	return s.String()
}

func (s *FilterUsersRequest) SetExcludeEndUserIds(v []*string) *FilterUsersRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *FilterUsersRequest) SetFilter(v string) *FilterUsersRequest {
	s.Filter = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeDesktopCount(v bool) *FilterUsersRequest {
	s.IncludeDesktopCount = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeDesktopGroupCount(v bool) *FilterUsersRequest {
	s.IncludeDesktopGroupCount = &v
	return s
}

func (s *FilterUsersRequest) SetMaxResults(v int64) *FilterUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *FilterUsersRequest) SetNextToken(v string) *FilterUsersRequest {
	s.NextToken = &v
	return s
}

func (s *FilterUsersRequest) SetOrderParam(v *FilterUsersRequestOrderParam) *FilterUsersRequest {
	s.OrderParam = v
	return s
}

func (s *FilterUsersRequest) SetOrgId(v string) *FilterUsersRequest {
	s.OrgId = &v
	return s
}

func (s *FilterUsersRequest) SetOwnerType(v string) *FilterUsersRequest {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersRequest) SetPropertyFilterParam(v []*FilterUsersRequestPropertyFilterParam) *FilterUsersRequest {
	s.PropertyFilterParam = v
	return s
}

func (s *FilterUsersRequest) SetPropertyKeyValueFilterParam(v []*FilterUsersRequestPropertyKeyValueFilterParam) *FilterUsersRequest {
	s.PropertyKeyValueFilterParam = v
	return s
}

type FilterUsersRequestOrderParam struct {
	// The method that you want to use to sort query results.
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// Specifies whether to sort query results in ascending or descending order.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s FilterUsersRequestOrderParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequestOrderParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestOrderParam) SetOrderField(v string) *FilterUsersRequestOrderParam {
	s.OrderField = &v
	return s
}

func (s *FilterUsersRequestOrderParam) SetOrderType(v string) *FilterUsersRequestOrderParam {
	s.OrderType = &v
	return s
}

type FilterUsersRequestPropertyFilterParam struct {
	// The ID of the property.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The IDs of the property values.
	PropertyValueIds *string `json:"PropertyValueIds,omitempty" xml:"PropertyValueIds,omitempty"`
}

func (s FilterUsersRequestPropertyFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequestPropertyFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestPropertyFilterParam) SetPropertyId(v int64) *FilterUsersRequestPropertyFilterParam {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersRequestPropertyFilterParam) SetPropertyValueIds(v string) *FilterUsersRequestPropertyFilterParam {
	s.PropertyValueIds = &v
	return s
}

type FilterUsersRequestPropertyKeyValueFilterParam struct {
	// The name of the property.
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The values of the property.
	PropertyValues *string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty"`
}

func (s FilterUsersRequestPropertyKeyValueFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequestPropertyKeyValueFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) SetPropertyKey(v string) *FilterUsersRequestPropertyKeyValueFilterParam {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) SetPropertyValues(v string) *FilterUsersRequestPropertyKeyValueFilterParam {
	s.PropertyValues = &v
	return s
}

type FilterUsersShrinkRequest struct {
	// The IDs of excluded users.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that you enter for a fuzzy search. You can enter a string to match the username or email address.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return information about cloud desktops that are assigned to the convenience user.
	IncludeDesktopCount *bool `json:"IncludeDesktopCount,omitempty" xml:"IncludeDesktopCount,omitempty"`
	// Specifies whether to return the number of desktop groups that are assigned to the user.
	IncludeDesktopGroupCount *bool `json:"IncludeDesktopGroupCount,omitempty" xml:"IncludeDesktopGroupCount,omitempty"`
	// The number of entries to return on each page. If you set this parameter to a value greater than 100, the system resets the value to 100.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. You do not need to configure this parameter if you call this operation for the first time. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to perform the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The parameter that is supported to sort query results.
	OrderParamShrink *string `json:"OrderParam,omitempty" xml:"OrderParam,omitempty"`
	// The ID of the organization.
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The type of the account ownership.
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// Details of the user property that you want to perform fuzzy search.
	PropertyFilterParam []*FilterUsersShrinkRequestPropertyFilterParam `json:"PropertyFilterParam,omitempty" xml:"PropertyFilterParam,omitempty" type:"Repeated"`
	// Details of the properties and property values.
	PropertyKeyValueFilterParam []*FilterUsersShrinkRequestPropertyKeyValueFilterParam `json:"PropertyKeyValueFilterParam,omitempty" xml:"PropertyKeyValueFilterParam,omitempty" type:"Repeated"`
}

func (s FilterUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequest) SetExcludeEndUserIds(v []*string) *FilterUsersShrinkRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *FilterUsersShrinkRequest) SetFilter(v string) *FilterUsersShrinkRequest {
	s.Filter = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeDesktopCount(v bool) *FilterUsersShrinkRequest {
	s.IncludeDesktopCount = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeDesktopGroupCount(v bool) *FilterUsersShrinkRequest {
	s.IncludeDesktopGroupCount = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetMaxResults(v int64) *FilterUsersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetNextToken(v string) *FilterUsersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOrderParamShrink(v string) *FilterUsersShrinkRequest {
	s.OrderParamShrink = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOrgId(v string) *FilterUsersShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOwnerType(v string) *FilterUsersShrinkRequest {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetPropertyFilterParam(v []*FilterUsersShrinkRequestPropertyFilterParam) *FilterUsersShrinkRequest {
	s.PropertyFilterParam = v
	return s
}

func (s *FilterUsersShrinkRequest) SetPropertyKeyValueFilterParam(v []*FilterUsersShrinkRequestPropertyKeyValueFilterParam) *FilterUsersShrinkRequest {
	s.PropertyKeyValueFilterParam = v
	return s
}

type FilterUsersShrinkRequestPropertyFilterParam struct {
	// The ID of the property.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The IDs of the property values.
	PropertyValueIds *string `json:"PropertyValueIds,omitempty" xml:"PropertyValueIds,omitempty"`
}

func (s FilterUsersShrinkRequestPropertyFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersShrinkRequestPropertyFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) SetPropertyId(v int64) *FilterUsersShrinkRequestPropertyFilterParam {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) SetPropertyValueIds(v string) *FilterUsersShrinkRequestPropertyFilterParam {
	s.PropertyValueIds = &v
	return s
}

type FilterUsersShrinkRequestPropertyKeyValueFilterParam struct {
	// The name of the property.
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The values of the property.
	PropertyValues *string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty"`
}

func (s FilterUsersShrinkRequestPropertyKeyValueFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersShrinkRequestPropertyKeyValueFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) SetPropertyKey(v string) *FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) SetPropertyValues(v string) *FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	s.PropertyValues = &v
	return s
}

type FilterUsersResponseBody struct {
	// The token that is used to query the next page. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to perform the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the convenience users.
	Users []*FilterUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s FilterUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBody) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBody) SetNextToken(v string) *FilterUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *FilterUsersResponseBody) SetRequestId(v string) *FilterUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *FilterUsersResponseBody) SetUsers(v []*FilterUsersResponseBodyUsers) *FilterUsersResponseBody {
	s.Users = v
	return s
}

type FilterUsersResponseBodyUsers struct {
	// The number of cloud desktops that are assigned to the user.
	DesktopCount *int64 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of authorized desktop groups that are owned by the user. This value is returned if you set `IncludeDesktopGroupCount` to `true`.
	DesktopGroupCount *int64 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The email address of the user.
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnableAdminAccess *bool   `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// The name of the user.
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The additional information about the user.
	ExternalInfo *FilterUsersResponseBodyUsersExternalInfo `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty" type:"Struct"`
	// The ID of the user.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether the user is a tenant administrator.
	IsTenantManager *bool `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	// The type of the account ownership.
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The mobile number of the user.
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The remarks of the user.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the user.
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Details of the user properties.
	UserSetPropertiesModels []*FilterUsersResponseBodyUsersUserSetPropertiesModels `json:"UserSetPropertiesModels,omitempty" xml:"UserSetPropertiesModels,omitempty" type:"Repeated"`
}

func (s FilterUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsers) SetDesktopCount(v int64) *FilterUsersResponseBodyUsers {
	s.DesktopCount = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetDesktopGroupCount(v int64) *FilterUsersResponseBodyUsers {
	s.DesktopGroupCount = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEmail(v string) *FilterUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEnableAdminAccess(v bool) *FilterUsersResponseBodyUsers {
	s.EnableAdminAccess = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEndUserId(v string) *FilterUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetExternalInfo(v *FilterUsersResponseBodyUsersExternalInfo) *FilterUsersResponseBodyUsers {
	s.ExternalInfo = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetId(v int64) *FilterUsersResponseBodyUsers {
	s.Id = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetIsTenantManager(v bool) *FilterUsersResponseBodyUsers {
	s.IsTenantManager = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetOwnerType(v string) *FilterUsersResponseBodyUsers {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPhone(v string) *FilterUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetRemark(v string) *FilterUsersResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetStatus(v int64) *FilterUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetUserSetPropertiesModels(v []*FilterUsersResponseBodyUsersUserSetPropertiesModels) *FilterUsersResponseBodyUsers {
	s.UserSetPropertiesModels = v
	return s
}

type FilterUsersResponseBodyUsersExternalInfo struct {
	// The name of the external system account to which the user is connected.
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// The student ID or employee ID of the external system account that is connected to the user.
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
}

func (s FilterUsersResponseBodyUsersExternalInfo) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersExternalInfo) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersExternalInfo) SetExternalName(v string) *FilterUsersResponseBodyUsersExternalInfo {
	s.ExternalName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersExternalInfo) SetJobNumber(v string) *FilterUsersResponseBodyUsersExternalInfo {
	s.JobNumber = &v
	return s
}

type FilterUsersResponseBodyUsersUserSetPropertiesModels struct {
	// The ID of the property.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The name of the property.
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The ID of property.
	PropertyType *int32 `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// Details of the property value.
	PropertyValues []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
	// The ID of the user that is bound to the property.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user that is bound to the property.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModels) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModels) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyKey(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyType(v int32) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyType = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyValues(v []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyValues = v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetUserId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.UserId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetUserName(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.UserName = &v
	return s
}

type FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues struct {
	// The property value.
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) SetPropertyValue(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) SetPropertyValueId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValueId = &v
	return s
}

type FilterUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FilterUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FilterUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponse) GoString() string {
	return s.String()
}

func (s *FilterUsersResponse) SetHeaders(v map[string]*string) *FilterUsersResponse {
	s.Headers = v
	return s
}

func (s *FilterUsersResponse) SetStatusCode(v int32) *FilterUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *FilterUsersResponse) SetBody(v *FilterUsersResponseBody) *FilterUsersResponse {
	s.Body = v
	return s
}

type ListPropertyResponseBody struct {
	// The information about the properties.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	Properties []*ListPropertyResponseBodyProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The ID of the property.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBody) SetNextToken(v string) *ListPropertyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPropertyResponseBody) SetProperties(v []*ListPropertyResponseBodyProperties) *ListPropertyResponseBody {
	s.Properties = v
	return s
}

func (s *ListPropertyResponseBody) SetRequestId(v string) *ListPropertyResponseBody {
	s.RequestId = &v
	return s
}

type ListPropertyResponseBodyProperties struct {
	// The operation that you want to perform. Set the value to **ListProperty**.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// Queries all user properties within an Alibaba Cloud account.
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// ListProperty
	PropertyValues []*ListPropertyResponseBodyPropertiesPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s ListPropertyResponseBodyProperties) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponseBodyProperties) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBodyProperties) SetPropertyId(v int64) *ListPropertyResponseBodyProperties {
	s.PropertyId = &v
	return s
}

func (s *ListPropertyResponseBodyProperties) SetPropertyKey(v string) *ListPropertyResponseBodyProperties {
	s.PropertyKey = &v
	return s
}

func (s *ListPropertyResponseBodyProperties) SetPropertyValues(v []*ListPropertyResponseBodyPropertiesPropertyValues) *ListPropertyResponseBodyProperties {
	s.PropertyValues = v
	return s
}

type ListPropertyResponseBodyPropertiesPropertyValues struct {
	PropertyValue   *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	PropertyValueId *int64  `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s ListPropertyResponseBodyPropertiesPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponseBodyPropertiesPropertyValues) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) SetPropertyValue(v string) *ListPropertyResponseBodyPropertiesPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) SetPropertyValueId(v int64) *ListPropertyResponseBodyPropertiesPropertyValues {
	s.PropertyValueId = &v
	return s
}

type ListPropertyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponse) GoString() string {
	return s.String()
}

func (s *ListPropertyResponse) SetHeaders(v map[string]*string) *ListPropertyResponse {
	s.Headers = v
	return s
}

func (s *ListPropertyResponse) SetStatusCode(v int32) *ListPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPropertyResponse) SetBody(v *ListPropertyResponseBody) *ListPropertyResponse {
	s.Body = v
	return s
}

type ListPropertyValueRequest struct {
	// Queries property values of a user property.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s ListPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *ListPropertyValueRequest) SetPropertyId(v int64) *ListPropertyValueRequest {
	s.PropertyId = &v
	return s
}

type ListPropertyValueResponseBody struct {
	PropertyValueInfos []*ListPropertyValueResponseBodyPropertyValueInfos `json:"PropertyValueInfos,omitempty" xml:"PropertyValueInfos,omitempty" type:"Repeated"`
	// ListPropertyValue
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponseBody) SetPropertyValueInfos(v []*ListPropertyValueResponseBodyPropertyValueInfos) *ListPropertyValueResponseBody {
	s.PropertyValueInfos = v
	return s
}

func (s *ListPropertyValueResponseBody) SetRequestId(v string) *ListPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

type ListPropertyValueResponseBodyPropertyValueInfos struct {
	PropertyValue   *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	PropertyValueId *int64  `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s ListPropertyValueResponseBodyPropertyValueInfos) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueResponseBodyPropertyValueInfos) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) SetPropertyValue(v string) *ListPropertyValueResponseBodyPropertyValueInfos {
	s.PropertyValue = &v
	return s
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) SetPropertyValueId(v int64) *ListPropertyValueResponseBodyPropertyValueInfos {
	s.PropertyValueId = &v
	return s
}

type ListPropertyValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponse) SetHeaders(v map[string]*string) *ListPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *ListPropertyValueResponse) SetStatusCode(v int32) *ListPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPropertyValueResponse) SetBody(v *ListPropertyValueResponseBody) *ListPropertyValueResponse {
	s.Body = v
	return s
}

type LockMfaDeviceRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s LockMfaDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceRequest) SetSerialNumber(v string) *LockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

type LockMfaDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockMfaDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceResponseBody) SetRequestId(v string) *LockMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

type LockMfaDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LockMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockMfaDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceResponse) SetHeaders(v map[string]*string) *LockMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *LockMfaDeviceResponse) SetStatusCode(v int32) *LockMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *LockMfaDeviceResponse) SetBody(v *LockMfaDeviceResponseBody) *LockMfaDeviceResponse {
	s.Body = v
	return s
}

type LockUsersRequest struct {
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s LockUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s LockUsersRequest) GoString() string {
	return s.String()
}

func (s *LockUsersRequest) SetUsers(v []*string) *LockUsersRequest {
	s.Users = v
	return s
}

type LockUsersResponseBody struct {
	LockUsersResult *LockUsersResponseBodyLockUsersResult `json:"LockUsersResult,omitempty" xml:"LockUsersResult,omitempty" type:"Struct"`
	RequestId       *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBody) SetLockUsersResult(v *LockUsersResponseBodyLockUsersResult) *LockUsersResponseBody {
	s.LockUsersResult = v
	return s
}

func (s *LockUsersResponseBody) SetRequestId(v string) *LockUsersResponseBody {
	s.RequestId = &v
	return s
}

type LockUsersResponseBodyLockUsersResult struct {
	FailedUsers []*LockUsersResponseBodyLockUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	LockedUsers []*string                                          `json:"LockedUsers,omitempty" xml:"LockedUsers,omitempty" type:"Repeated"`
}

func (s LockUsersResponseBodyLockUsersResult) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBodyLockUsersResult) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyLockUsersResult) SetFailedUsers(v []*LockUsersResponseBodyLockUsersResultFailedUsers) *LockUsersResponseBodyLockUsersResult {
	s.FailedUsers = v
	return s
}

func (s *LockUsersResponseBodyLockUsersResult) SetLockedUsers(v []*string) *LockUsersResponseBodyLockUsersResult {
	s.LockedUsers = v
	return s
}

type LockUsersResponseBodyLockUsersResultFailedUsers struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s LockUsersResponseBodyLockUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBodyLockUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetEndUserId(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetErrorCode(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetErrorMessage(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type LockUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponse) GoString() string {
	return s.String()
}

func (s *LockUsersResponse) SetHeaders(v map[string]*string) *LockUsersResponse {
	s.Headers = v
	return s
}

func (s *LockUsersResponse) SetStatusCode(v int32) *LockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *LockUsersResponse) SetBody(v *LockUsersResponseBody) *LockUsersResponse {
	s.Body = v
	return s
}

type ModifyUserRequest struct {
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetEndUserId(v string) *ModifyUserRequest {
	s.EndUserId = &v
	return s
}

func (s *ModifyUserRequest) SetPhone(v string) *ModifyUserRequest {
	s.Phone = &v
	return s
}

type ModifyUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserResponseBody) SetRequestId(v string) *ModifyUserResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserResponse) SetHeaders(v map[string]*string) *ModifyUserResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserResponse) SetStatusCode(v int32) *ModifyUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserResponse) SetBody(v *ModifyUserResponseBody) *ModifyUserResponse {
	s.Body = v
	return s
}

type QuerySyncStatusByAliUidResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QuerySyncStatusByAliUidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySyncStatusByAliUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponseBody) SetCode(v string) *QuerySyncStatusByAliUidResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetData(v *QuerySyncStatusByAliUidResponseBodyData) *QuerySyncStatusByAliUidResponseBody {
	s.Data = v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetHttpStatusCode(v int32) *QuerySyncStatusByAliUidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetMessage(v string) *QuerySyncStatusByAliUidResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetRequestId(v string) *QuerySyncStatusByAliUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetSuccess(v bool) *QuerySyncStatusByAliUidResponseBody {
	s.Success = &v
	return s
}

type QuerySyncStatusByAliUidResponseBodyData struct {
	AliUid            *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CorpId            *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	GmtCreated        *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified       *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LatestBeginTime   *string `json:"LatestBeginTime,omitempty" xml:"LatestBeginTime,omitempty"`
	LatestEndTime     *string `json:"LatestEndTime,omitempty" xml:"LatestEndTime,omitempty"`
	LatestSuccessTime *string `json:"LatestSuccessTime,omitempty" xml:"LatestSuccessTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySyncStatusByAliUidResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetAliUid(v int64) *QuerySyncStatusByAliUidResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetCorpId(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetGmtCreated(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetGmtModified(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetId(v int64) *QuerySyncStatusByAliUidResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestBeginTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestBeginTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestEndTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestEndTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestSuccessTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestSuccessTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetStatus(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.Status = &v
	return s
}

type QuerySyncStatusByAliUidResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySyncStatusByAliUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySyncStatusByAliUidResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponse) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponse) SetHeaders(v map[string]*string) *QuerySyncStatusByAliUidResponse {
	s.Headers = v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) SetStatusCode(v int32) *QuerySyncStatusByAliUidResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) SetBody(v *QuerySyncStatusByAliUidResponseBody) *QuerySyncStatusByAliUidResponse {
	s.Body = v
	return s
}

type RemoveMfaDeviceRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s RemoveMfaDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceRequest) SetSerialNumber(v string) *RemoveMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

type RemoveMfaDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMfaDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceResponseBody) SetRequestId(v string) *RemoveMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

type RemoveMfaDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveMfaDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceResponse) SetHeaders(v map[string]*string) *RemoveMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *RemoveMfaDeviceResponse) SetStatusCode(v int32) *RemoveMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMfaDeviceResponse) SetBody(v *RemoveMfaDeviceResponseBody) *RemoveMfaDeviceResponse {
	s.Body = v
	return s
}

type RemovePropertyRequest struct {
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s RemovePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePropertyRequest) GoString() string {
	return s.String()
}

func (s *RemovePropertyRequest) SetPropertyId(v int64) *RemovePropertyRequest {
	s.PropertyId = &v
	return s
}

type RemovePropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePropertyResponseBody) SetRequestId(v string) *RemovePropertyResponseBody {
	s.RequestId = &v
	return s
}

type RemovePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemovePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemovePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePropertyResponse) GoString() string {
	return s.String()
}

func (s *RemovePropertyResponse) SetHeaders(v map[string]*string) *RemovePropertyResponse {
	s.Headers = v
	return s
}

func (s *RemovePropertyResponse) SetStatusCode(v int32) *RemovePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePropertyResponse) SetBody(v *RemovePropertyResponseBody) *RemovePropertyResponse {
	s.Body = v
	return s
}

type RemoveUsersRequest struct {
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) SetUsers(v []*string) *RemoveUsersRequest {
	s.Users = v
	return s
}

type RemoveUsersResponseBody struct {
	RemoveUsersResult *RemoveUsersResponseBodyRemoveUsersResult `json:"RemoveUsersResult,omitempty" xml:"RemoveUsersResult,omitempty" type:"Struct"`
	RequestId         *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) SetRemoveUsersResult(v *RemoveUsersResponseBodyRemoveUsersResult) *RemoveUsersResponseBody {
	s.RemoveUsersResult = v
	return s
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUsersResponseBodyRemoveUsersResult struct {
	FailedUsers  []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	RemovedUsers []*string                                              `json:"RemovedUsers,omitempty" xml:"RemovedUsers,omitempty" type:"Repeated"`
}

func (s RemoveUsersResponseBodyRemoveUsersResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBodyRemoveUsersResult) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) SetFailedUsers(v []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers) *RemoveUsersResponseBodyRemoveUsersResult {
	s.FailedUsers = v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) SetRemovedUsers(v []*string) *RemoveUsersResponseBodyRemoveUsersResult {
	s.RemovedUsers = v
	return s
}

type RemoveUsersResponseBodyRemoveUsersResultFailedUsers struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s RemoveUsersResponseBodyRemoveUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBodyRemoveUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetEndUserId(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetErrorCode(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetErrorMessage(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type RemoveUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponse) SetHeaders(v map[string]*string) *RemoveUsersResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersResponse) SetStatusCode(v int32) *RemoveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersResponse) SetBody(v *RemoveUsersResponseBody) *RemoveUsersResponse {
	s.Body = v
	return s
}

type ResetUserPasswordRequest struct {
	NotifyType *int32    `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Users      []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ResetUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) SetNotifyType(v int32) *ResetUserPasswordRequest {
	s.NotifyType = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUsers(v []*string) *ResetUserPasswordRequest {
	s.Users = v
	return s
}

type ResetUserPasswordResponseBody struct {
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResetUsersResult *ResetUserPasswordResponseBodyResetUsersResult `json:"ResetUsersResult,omitempty" xml:"ResetUsersResult,omitempty" type:"Struct"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetResetUsersResult(v *ResetUserPasswordResponseBodyResetUsersResult) *ResetUserPasswordResponseBody {
	s.ResetUsersResult = v
	return s
}

type ResetUserPasswordResponseBodyResetUsersResult struct {
	FailedUsers []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	ResetUsers  []*string                                                   `json:"ResetUsers,omitempty" xml:"ResetUsers,omitempty" type:"Repeated"`
}

func (s ResetUserPasswordResponseBodyResetUsersResult) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBodyResetUsersResult) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) SetFailedUsers(v []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers) *ResetUserPasswordResponseBodyResetUsersResult {
	s.FailedUsers = v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) SetResetUsers(v []*string) *ResetUserPasswordResponseBodyResetUsersResult {
	s.ResetUsers = v
	return s
}

type ResetUserPasswordResponseBodyResetUsersResultFailedUsers struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ResetUserPasswordResponseBodyResetUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBodyResetUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetEndUserId(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetErrorCode(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetErrorMessage(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type ResetUserPasswordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponse) SetHeaders(v map[string]*string) *ResetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetUserPasswordResponse) SetStatusCode(v int32) *ResetUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetUserPasswordResponse) SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse {
	s.Body = v
	return s
}

type SetUserPropertyValueRequest struct {
	PropertyId      *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// Associates a user property with a user.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// SetUserPropertyValue
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SetUserPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueRequest) SetPropertyId(v int64) *SetUserPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetPropertyValueId(v int64) *SetUserPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetUserId(v int64) *SetUserPropertyValueRequest {
	s.UserId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetUserName(v string) *SetUserPropertyValueRequest {
	s.UserName = &v
	return s
}

type SetUserPropertyValueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueResponseBody) SetRequestId(v string) *SetUserPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

type SetUserPropertyValueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetUserPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetUserPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueResponse) SetHeaders(v map[string]*string) *SetUserPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *SetUserPropertyValueResponse) SetStatusCode(v int32) *SetUserPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserPropertyValueResponse) SetBody(v *SetUserPropertyValueResponseBody) *SetUserPropertyValueResponse {
	s.Body = v
	return s
}

type SyncAllEduInfoResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncAllEduInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncAllEduInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncAllEduInfoResponseBody) SetCode(v string) *SyncAllEduInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetHttpStatusCode(v int32) *SyncAllEduInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetMessage(v string) *SyncAllEduInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetRequestId(v string) *SyncAllEduInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetSuccess(v bool) *SyncAllEduInfoResponseBody {
	s.Success = &v
	return s
}

type SyncAllEduInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncAllEduInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncAllEduInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncAllEduInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncAllEduInfoResponse) SetHeaders(v map[string]*string) *SyncAllEduInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncAllEduInfoResponse) SetStatusCode(v int32) *SyncAllEduInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncAllEduInfoResponse) SetBody(v *SyncAllEduInfoResponseBody) *SyncAllEduInfoResponse {
	s.Body = v
	return s
}

type UnlockMfaDeviceRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnlockMfaDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceRequest) SetSerialNumber(v string) *UnlockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

type UnlockMfaDeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockMfaDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceResponseBody) SetRequestId(v string) *UnlockMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnlockMfaDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnlockMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockMfaDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceResponse) SetHeaders(v map[string]*string) *UnlockMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockMfaDeviceResponse) SetStatusCode(v int32) *UnlockMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockMfaDeviceResponse) SetBody(v *UnlockMfaDeviceResponseBody) *UnlockMfaDeviceResponse {
	s.Body = v
	return s
}

type UnlockUsersRequest struct {
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s UnlockUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersRequest) GoString() string {
	return s.String()
}

func (s *UnlockUsersRequest) SetUsers(v []*string) *UnlockUsersRequest {
	s.Users = v
	return s
}

type UnlockUsersResponseBody struct {
	RequestId         *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UnlockUsersResult *UnlockUsersResponseBodyUnlockUsersResult `json:"UnlockUsersResult,omitempty" xml:"UnlockUsersResult,omitempty" type:"Struct"`
}

func (s UnlockUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBody) SetRequestId(v string) *UnlockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockUsersResponseBody) SetUnlockUsersResult(v *UnlockUsersResponseBodyUnlockUsersResult) *UnlockUsersResponseBody {
	s.UnlockUsersResult = v
	return s
}

type UnlockUsersResponseBodyUnlockUsersResult struct {
	FailedUsers   []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	UnlockedUsers []*string                                              `json:"UnlockedUsers,omitempty" xml:"UnlockedUsers,omitempty" type:"Repeated"`
}

func (s UnlockUsersResponseBodyUnlockUsersResult) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBodyUnlockUsersResult) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) SetFailedUsers(v []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers) *UnlockUsersResponseBodyUnlockUsersResult {
	s.FailedUsers = v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) SetUnlockedUsers(v []*string) *UnlockUsersResponseBodyUnlockUsersResult {
	s.UnlockedUsers = v
	return s
}

type UnlockUsersResponseBodyUnlockUsersResultFailedUsers struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s UnlockUsersResponseBodyUnlockUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBodyUnlockUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetEndUserId(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetErrorCode(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetErrorMessage(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type UnlockUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnlockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponse) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponse) SetHeaders(v map[string]*string) *UnlockUsersResponse {
	s.Headers = v
	return s
}

func (s *UnlockUsersResponse) SetStatusCode(v int32) *UnlockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockUsersResponse) SetBody(v *UnlockUsersResponseBody) *UnlockUsersResponse {
	s.Body = v
	return s
}

type UpdatePropertyRequest struct {
	// The operation that you want to perform. Set the value to **UpdateProperty**.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property values that failed to be modified.
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The new property value N.
	PropertyValues []*UpdatePropertyRequestPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s UpdatePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePropertyRequest) SetPropertyId(v int64) *UpdatePropertyRequest {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyRequest) SetPropertyKey(v string) *UpdatePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *UpdatePropertyRequest) SetPropertyValues(v []*UpdatePropertyRequestPropertyValues) *UpdatePropertyRequest {
	s.PropertyValues = v
	return s
}

type UpdatePropertyRequestPropertyValues struct {
	// The error code.
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property that you want to modify. You can call the [ListProperty](~~410890~~) operation to query the property ID.
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s UpdatePropertyRequestPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyRequestPropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyRequestPropertyValues) SetPropertyValue(v string) *UpdatePropertyRequestPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyRequestPropertyValues) SetPropertyValueId(v int64) *UpdatePropertyRequestPropertyValues {
	s.PropertyValueId = &v
	return s
}

type UpdatePropertyResponseBody struct {
	// The name of the property.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the property.
	UpdateResult *UpdatePropertyResponseBodyUpdateResult `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty" type:"Struct"`
}

func (s UpdatePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBody) SetRequestId(v string) *UpdatePropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePropertyResponseBody) SetUpdateResult(v *UpdatePropertyResponseBodyUpdateResult) *UpdatePropertyResponseBody {
	s.UpdateResult = v
	return s
}

type UpdatePropertyResponseBodyUpdateResult struct {
	// The ID of the property.
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The ID of property value N that you want to modify. You can call the [ListProperty](~~410890~~) operation to query the property value ID.
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property values that were modified.
	SavePropertyValueModel *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel `json:"SavePropertyValueModel,omitempty" xml:"SavePropertyValueModel,omitempty" type:"Struct"`
}

func (s UpdatePropertyResponseBodyUpdateResult) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResult) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetPropertyId(v int64) *UpdatePropertyResponseBodyUpdateResult {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetPropertyKey(v string) *UpdatePropertyResponseBodyUpdateResult {
	s.PropertyKey = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetSavePropertyValueModel(v *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) *UpdatePropertyResponseBodyUpdateResult {
	s.SavePropertyValueModel = v
	return s
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel struct {
	// UpdateProperty
	FailedPropertyValues []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues `json:"FailedPropertyValues,omitempty" xml:"FailedPropertyValues,omitempty" type:"Repeated"`
	// The ID of the request.
	SavePropertyValues []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues `json:"SavePropertyValues,omitempty" xml:"SavePropertyValues,omitempty" type:"Repeated"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) SetFailedPropertyValues(v []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	s.FailedPropertyValues = v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) SetSavePropertyValues(v []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	s.SavePropertyValues = v
	return s
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PropertyId    *int64  `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetErrorCode(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetErrorMessage(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetPropertyId(v int64) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetPropertyValue(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyValue = &v
	return s
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues struct {
	// Modifies a user property.
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The value of the property.
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) SetPropertyValue(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) SetPropertyValueId(v int64) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValueId = &v
	return s
}

type UpdatePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponse) SetHeaders(v map[string]*string) *UpdatePropertyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePropertyResponse) SetStatusCode(v int32) *UpdatePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePropertyResponse) SetBody(v *UpdatePropertyResponseBody) *UpdatePropertyResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eds-user"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUsedPropertyWithOptions(request *CheckUsedPropertyRequest, runtime *util.RuntimeOptions) (_result *CheckUsedPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		query["PropertyId"] = request.PropertyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUsedProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUsedPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUsedProperty(request *CheckUsedPropertyRequest) (_result *CheckUsedPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUsedPropertyResponse{}
	_body, _err := client.CheckUsedPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The operation that you want to perform. Set the value to **CheckUsedPropertyValue**.
 *
 * @param request CheckUsedPropertyValueRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CheckUsedPropertyValueResponse
 */
func (client *Client) CheckUsedPropertyValueWithOptions(request *CheckUsedPropertyValueRequest, runtime *util.RuntimeOptions) (_result *CheckUsedPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		query["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValueId)) {
		query["PropertyValueId"] = request.PropertyValueId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUsedPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUsedPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The operation that you want to perform. Set the value to **CheckUsedPropertyValue**.
 *
 * @param request CheckUsedPropertyValueRequest
 * @return CheckUsedPropertyValueResponse
 */
func (client *Client) CheckUsedPropertyValue(request *CheckUsedPropertyValueRequest) (_result *CheckUsedPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUsedPropertyValueResponse{}
	_body, _err := client.CheckUsedPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePropertyWithOptions(request *CreatePropertyRequest, runtime *util.RuntimeOptions) (_result *CreatePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyKey)) {
		body["PropertyKey"] = request.PropertyKey
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValues)) {
		body["PropertyValues"] = request.PropertyValues
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProperty(request *CreatePropertyRequest) (_result *CreatePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePropertyResponse{}
	_body, _err := client.CreatePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The mobile number of the end user.
 *
 * @param request CreateUsersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateUsersResponse
 */
func (client *Client) CreateUsersWithOptions(request *CreateUsersRequest, runtime *util.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The mobile number of the end user.
 *
 * @param request CreateUsersRequest
 * @return CreateUsersResponse
 */
func (client *Client) CreateUsers(request *CreateUsersRequest) (_result *CreateUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUsersResponse{}
	_body, _err := client.CreateUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The operation that you want to perform. Set the value to **DeleteUserPropertyValue**.
 *
 * @param request DeleteUserPropertyValueRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteUserPropertyValueResponse
 */
func (client *Client) DeleteUserPropertyValueWithOptions(request *DeleteUserPropertyValueRequest, runtime *util.RuntimeOptions) (_result *DeleteUserPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValueId)) {
		body["PropertyValueId"] = request.PropertyValueId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The operation that you want to perform. Set the value to **DeleteUserPropertyValue**.
 *
 * @param request DeleteUserPropertyValueRequest
 * @return DeleteUserPropertyValueResponse
 */
func (client *Client) DeleteUserPropertyValue(request *DeleteUserPropertyValueRequest) (_result *DeleteUserPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserPropertyValueResponse{}
	_body, _err := client.DeleteUserPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMfaDevicesWithOptions(request *DescribeMfaDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeMfaDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumbers)) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMfaDevices"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMfaDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMfaDevices(request *DescribeMfaDevicesRequest) (_result *DescribeMfaDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMfaDevicesResponse{}
	_body, _err := client.DescribeMfaDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsersWithOptions(request *DescribeUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeEndUserIds)) {
		body["ExcludeEndUserIds"] = request.ExcludeEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsers(request *DescribeUsersRequest) (_result *DescribeUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersResponse{}
	_body, _err := client.DescribeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FilterUsersWithOptions(tmpReq *FilterUsersRequest, runtime *util.RuntimeOptions) (_result *FilterUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FilterUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderParam)) {
		request.OrderParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderParam, tea.String("OrderParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeEndUserIds)) {
		query["ExcludeEndUserIds"] = request.ExcludeEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDesktopCount)) {
		query["IncludeDesktopCount"] = request.IncludeDesktopCount
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDesktopGroupCount)) {
		query["IncludeDesktopGroupCount"] = request.IncludeDesktopGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderParamShrink)) {
		query["OrderParam"] = request.OrderParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerType)) {
		query["OwnerType"] = request.OwnerType
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyFilterParam)) {
		query["PropertyFilterParam"] = request.PropertyFilterParam
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyKeyValueFilterParam)) {
		query["PropertyKeyValueFilterParam"] = request.PropertyKeyValueFilterParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FilterUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FilterUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FilterUsers(request *FilterUsersRequest) (_result *FilterUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FilterUsersResponse{}
	_body, _err := client.FilterUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPropertyWithOptions(runtime *util.RuntimeOptions) (_result *ListPropertyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProperty() (_result *ListPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPropertyResponse{}
	_body, _err := client.ListPropertyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPropertyValueWithOptions(request *ListPropertyValueRequest, runtime *util.RuntimeOptions) (_result *ListPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		query["PropertyId"] = request.PropertyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPropertyValue(request *ListPropertyValueRequest) (_result *ListPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPropertyValueResponse{}
	_body, _err := client.ListPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Locks a virtual MFA device that is bound to a convenience user.
 *
 * @param request LockMfaDeviceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return LockMfaDeviceResponse
 */
func (client *Client) LockMfaDeviceWithOptions(request *LockMfaDeviceRequest, runtime *util.RuntimeOptions) (_result *LockMfaDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockMfaDevice"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockMfaDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Locks a virtual MFA device that is bound to a convenience user.
 *
 * @param request LockMfaDeviceRequest
 * @return LockMfaDeviceResponse
 */
func (client *Client) LockMfaDevice(request *LockMfaDeviceRequest) (_result *LockMfaDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockMfaDeviceResponse{}
	_body, _err := client.LockMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockUsersWithOptions(request *LockUsersRequest, runtime *util.RuntimeOptions) (_result *LockUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LockUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockUsers(request *LockUsersRequest) (_result *LockUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockUsersResponse{}
	_body, _err := client.LockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *util.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUser"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUser(request *ModifyUserRequest) (_result *ModifyUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserResponse{}
	_body, _err := client.ModifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySyncStatusByAliUidWithOptions(runtime *util.RuntimeOptions) (_result *QuerySyncStatusByAliUidResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QuerySyncStatusByAliUid"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySyncStatusByAliUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySyncStatusByAliUid() (_result *QuerySyncStatusByAliUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySyncStatusByAliUidResponse{}
	_body, _err := client.QuerySyncStatusByAliUidWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMfaDeviceWithOptions(request *RemoveMfaDeviceRequest, runtime *util.RuntimeOptions) (_result *RemoveMfaDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveMfaDevice"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveMfaDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMfaDevice(request *RemoveMfaDeviceRequest) (_result *RemoveMfaDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveMfaDeviceResponse{}
	_body, _err := client.RemoveMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePropertyWithOptions(request *RemovePropertyRequest, runtime *util.RuntimeOptions) (_result *RemovePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemovePropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveProperty(request *RemovePropertyRequest) (_result *RemovePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePropertyResponse{}
	_body, _err := client.RemovePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		body["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetUserPassword"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the request.
 *
 * @param request SetUserPropertyValueRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetUserPropertyValueResponse
 */
func (client *Client) SetUserPropertyValueWithOptions(request *SetUserPropertyValueRequest, runtime *util.RuntimeOptions) (_result *SetUserPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValueId)) {
		body["PropertyValueId"] = request.PropertyValueId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the request.
 *
 * @param request SetUserPropertyValueRequest
 * @return SetUserPropertyValueResponse
 */
func (client *Client) SetUserPropertyValue(request *SetUserPropertyValueRequest) (_result *SetUserPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserPropertyValueResponse{}
	_body, _err := client.SetUserPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncAllEduInfoWithOptions(runtime *util.RuntimeOptions) (_result *SyncAllEduInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("SyncAllEduInfo"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncAllEduInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncAllEduInfo() (_result *SyncAllEduInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncAllEduInfoResponse{}
	_body, _err := client.SyncAllEduInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockMfaDeviceWithOptions(request *UnlockMfaDeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockMfaDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockMfaDevice"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockMfaDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockMfaDevice(request *UnlockMfaDeviceRequest) (_result *UnlockMfaDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockMfaDeviceResponse{}
	_body, _err := client.UnlockMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockUsersWithOptions(request *UnlockUsersRequest, runtime *util.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockUsers(request *UnlockUsersRequest) (_result *UnlockUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockUsersResponse{}
	_body, _err := client.UnlockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePropertyWithOptions(request *UpdatePropertyRequest, runtime *util.RuntimeOptions) (_result *UpdatePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyKey)) {
		body["PropertyKey"] = request.PropertyKey
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValues)) {
		body["PropertyValues"] = request.PropertyValues
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProperty(request *UpdatePropertyRequest) (_result *UpdatePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePropertyResponse{}
	_body, _err := client.UpdatePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
