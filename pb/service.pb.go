// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcf,
	0x05, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x17, 0x69, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x1c, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x44, 0x65, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x12, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x1c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0f, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0x6f, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x13, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x12, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),                    // 0: CreateUserRequest
	(*InsertEmailConfirmTokenRequest)(nil),       // 1: InsertEmailConfirmTokenRequest
	(*CheckAndDelEmailConfirmTokenRequest)(nil),  // 2: CheckAndDelEmailConfirmTokenRequest
	(*GenerateTokenRequest)(nil),                 // 3: GenerateTokenRequest
	(*ParseTokenRequest)(nil),                    // 4: ParseTokenRequest
	(*GetDetailsRequest)(nil),                    // 5: GetDetailsRequest
	(*UpdateRequest)(nil),                        // 6: UpdateRequest
	(*UpdateEmailVerificationTokenRequest)(nil),  // 7: UpdateEmailVerificationTokenRequest
	(*RecoverPasswordRequest)(nil),               // 8: RecoverPasswordRequest
	(*ActionsRequest)(nil),                       // 9: ActionsRequest
	(*MailConfirmRequest)(nil),                   // 10: MailConfirmRequest
	(*ResetRequest)(nil),                         // 11: ResetRequest
	(*CreateUserResponse)(nil),                   // 12: CreateUserResponse
	(*InsertEmailConfirmTokenResponse)(nil),      // 13: InsertEmailConfirmTokenResponse
	(*CheckAndDelEmailConfirmTokenResponse)(nil), // 14: CheckAndDelEmailConfirmTokenResponse
	(*GenerateTokenResponse)(nil),                // 15: GenerateTokenResponse
	(*ParseTokenResponse)(nil),                   // 16: ParseTokenResponse
	(*User)(nil),                                 // 17: User
	(*UpdateResponse)(nil),                       // 18: UpdateResponse
	(*UpdateEmailVerificationTokenResponse)(nil), // 19: UpdateEmailVerificationTokenResponse
	(*RecoverPasswordResponse)(nil),              // 20: RecoverPasswordResponse
	(*ActionsResponse)(nil),                      // 21: ActionsResponse
	(*MailConfirmResponse)(nil),                  // 22: MailConfirmResponse
	(*ResetResponse)(nil),                        // 23: ResetResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: AuthService.createUser:input_type -> CreateUserRequest
	1,  // 1: AuthService.insertEmailConfirmToken:input_type -> InsertEmailConfirmTokenRequest
	2,  // 2: AuthService.checkAndDelEmailConfirmToken:input_type -> CheckAndDelEmailConfirmTokenRequest
	3,  // 3: AuthService.generateToken:input_type -> GenerateTokenRequest
	4,  // 4: AuthService.parseToken:input_type -> ParseTokenRequest
	5,  // 5: AuthService.getDetails:input_type -> GetDetailsRequest
	6,  // 6: AuthService.update:input_type -> UpdateRequest
	7,  // 7: AuthService.updateEmailVerificationToken:input_type -> UpdateEmailVerificationTokenRequest
	8,  // 8: AuthService.recoverPassword:input_type -> RecoverPasswordRequest
	9,  // 9: AuthService.actions:input_type -> ActionsRequest
	10, // 10: MailService.confirm:input_type -> MailConfirmRequest
	11, // 11: MailService.reset:input_type -> ResetRequest
	12, // 12: AuthService.createUser:output_type -> CreateUserResponse
	13, // 13: AuthService.insertEmailConfirmToken:output_type -> InsertEmailConfirmTokenResponse
	14, // 14: AuthService.checkAndDelEmailConfirmToken:output_type -> CheckAndDelEmailConfirmTokenResponse
	15, // 15: AuthService.generateToken:output_type -> GenerateTokenResponse
	16, // 16: AuthService.parseToken:output_type -> ParseTokenResponse
	17, // 17: AuthService.getDetails:output_type -> User
	18, // 18: AuthService.update:output_type -> UpdateResponse
	19, // 19: AuthService.updateEmailVerificationToken:output_type -> UpdateEmailVerificationTokenResponse
	20, // 20: AuthService.recoverPassword:output_type -> RecoverPasswordResponse
	21, // 21: AuthService.actions:output_type -> ActionsResponse
	22, // 22: MailService.confirm:output_type -> MailConfirmResponse
	23, // 23: MailService.reset:output_type -> ResetResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
