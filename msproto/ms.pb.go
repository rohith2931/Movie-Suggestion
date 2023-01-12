// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: msproto/ms.proto

package movieSuggestion

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName    string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{0}
}

func (x *NewUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *NewUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NewUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *NewUser) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type EmptyMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMovie) Reset() {
	*x = EmptyMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMovie) ProtoMessage() {}

func (x *EmptyMovie) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMovie.ProtoReflect.Descriptor instead.
func (*EmptyMovie) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{2}
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Director    string `protobuf:"bytes,3,opt,name=director,proto3" json:"director,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Rating      uint64 `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Language    string `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	Category    string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	ReleaseDate string `protobuf:"bytes,8,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{3}
}

func (x *Movie) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Movie) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Movie) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Movie) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

type NewMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Director    string `protobuf:"bytes,2,opt,name=director,proto3" json:"director,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Rating      uint64 `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Language    string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	ReleaseDate string `protobuf:"bytes,7,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
}

func (x *NewMovie) Reset() {
	*x = NewMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMovie) ProtoMessage() {}

func (x *NewMovie) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMovie.ProtoReflect.Descriptor instead.
func (*NewMovie) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{4}
}

func (x *NewMovie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewMovie) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *NewMovie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewMovie) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *NewMovie) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *NewMovie) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NewMovie) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

type MovieCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *MovieCategory) Reset() {
	*x = MovieCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieCategory) ProtoMessage() {}

func (x *MovieCategory) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieCategory.ProtoReflect.Descriptor instead.
func (*MovieCategory) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{5}
}

func (x *MovieCategory) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Movies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *Movies) Reset() {
	*x = Movies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movies) ProtoMessage() {}

func (x *Movies) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movies.ProtoReflect.Descriptor instead.
func (*Movies) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{6}
}

func (x *Movies) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

type AddMovieByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId uint64 `protobuf:"varint,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
}

func (x *AddMovieByUser) Reset() {
	*x = AddMovieByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMovieByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMovieByUser) ProtoMessage() {}

func (x *AddMovieByUser) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMovieByUser.ProtoReflect.Descriptor instead.
func (*AddMovieByUser) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{7}
}

func (x *AddMovieByUser) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddMovieByUser) GetMovieId() uint64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

type NewReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating      uint64 `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
	MovieId     uint64 `protobuf:"varint,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	UserId      uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *NewReview) Reset() {
	*x = NewReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewReview) ProtoMessage() {}

func (x *NewReview) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewReview.ProtoReflect.Descriptor instead.
func (*NewReview) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{8}
}

func (x *NewReview) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *NewReview) GetMovieId() uint64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *NewReview) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NewReview) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rating      uint64 `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	MovieId     uint64 `protobuf:"varint,3,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	UserId      uint64 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msproto_ms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_msproto_ms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_msproto_ms_proto_rawDescGZIP(), []int{9}
}

func (x *Review) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Review) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetMovieId() uint64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *Review) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Review) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_msproto_ms_proto protoreflect.FileDescriptor

var file_msproto_ms_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x07,
	0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0x30, 0x0a, 0x06, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x09,
	0x4e, 0x65, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x32, 0xc5, 0x04, 0x0a, 0x0a, 0x4d, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e,
	0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x0d, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x12, 0x13, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x0e, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x16, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x0e, 0x2e, 0x6d, 0x73, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x2e, 0x6d, 0x73, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x0e, 0x2e, 0x6d, 0x73, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x6d,
	0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x0f, 0x2e, 0x6d,
	0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x6d, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x0e, 0x2e, 0x6d, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x12, 0x2e,
	0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x1a, 0x0f, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x0f, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x0f, 0x2e, 0x6d, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msproto_ms_proto_rawDescOnce sync.Once
	file_msproto_ms_proto_rawDescData = file_msproto_ms_proto_rawDesc
)

func file_msproto_ms_proto_rawDescGZIP() []byte {
	file_msproto_ms_proto_rawDescOnce.Do(func() {
		file_msproto_ms_proto_rawDescData = protoimpl.X.CompressGZIP(file_msproto_ms_proto_rawDescData)
	})
	return file_msproto_ms_proto_rawDescData
}

var file_msproto_ms_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_msproto_ms_proto_goTypes = []interface{}{
	(*NewUser)(nil),        // 0: msproto.NewUser
	(*User)(nil),           // 1: msproto.User
	(*EmptyMovie)(nil),     // 2: msproto.EmptyMovie
	(*Movie)(nil),          // 3: msproto.Movie
	(*NewMovie)(nil),       // 4: msproto.NewMovie
	(*MovieCategory)(nil),  // 5: msproto.MovieCategory
	(*Movies)(nil),         // 6: msproto.Movies
	(*AddMovieByUser)(nil), // 7: msproto.AddMovieByUser
	(*NewReview)(nil),      // 8: msproto.NewReview
	(*Review)(nil),         // 9: msproto.Review
}
var file_msproto_ms_proto_depIdxs = []int32{
	3,  // 0: msproto.Movies.movies:type_name -> msproto.Movie
	0,  // 1: msproto.MsDatabase.CreateUser:input_type -> msproto.NewUser
	2,  // 2: msproto.MsDatabase.GetAllMovies:input_type -> msproto.EmptyMovie
	5,  // 3: msproto.MsDatabase.GetMovieByCategory:input_type -> msproto.MovieCategory
	4,  // 4: msproto.MsDatabase.AddMovie:input_type -> msproto.NewMovie
	3,  // 5: msproto.MsDatabase.DeleteMovie:input_type -> msproto.Movie
	7,  // 6: msproto.MsDatabase.AddMovieToWatchlist:input_type -> msproto.AddMovieByUser
	2,  // 7: msproto.MsDatabase.GetAllWatchlistMovies:input_type -> msproto.EmptyMovie
	3,  // 8: msproto.MsDatabase.DeleteMovieFromWatchlist:input_type -> msproto.Movie
	8,  // 9: msproto.MsDatabase.CreateReview:input_type -> msproto.NewReview
	9,  // 10: msproto.MsDatabase.UpdateReview:input_type -> msproto.Review
	1,  // 11: msproto.MsDatabase.CreateUser:output_type -> msproto.User
	3,  // 12: msproto.MsDatabase.GetAllMovies:output_type -> msproto.Movie
	6,  // 13: msproto.MsDatabase.GetMovieByCategory:output_type -> msproto.Movies
	3,  // 14: msproto.MsDatabase.AddMovie:output_type -> msproto.Movie
	3,  // 15: msproto.MsDatabase.DeleteMovie:output_type -> msproto.Movie
	3,  // 16: msproto.MsDatabase.AddMovieToWatchlist:output_type -> msproto.Movie
	6,  // 17: msproto.MsDatabase.GetAllWatchlistMovies:output_type -> msproto.Movies
	3,  // 18: msproto.MsDatabase.DeleteMovieFromWatchlist:output_type -> msproto.Movie
	9,  // 19: msproto.MsDatabase.CreateReview:output_type -> msproto.Review
	9,  // 20: msproto.MsDatabase.UpdateReview:output_type -> msproto.Review
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_msproto_ms_proto_init() }
func file_msproto_ms_proto_init() {
	if File_msproto_ms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msproto_ms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieCategory); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movies); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMovieByUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewReview); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msproto_ms_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msproto_ms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msproto_ms_proto_goTypes,
		DependencyIndexes: file_msproto_ms_proto_depIdxs,
		MessageInfos:      file_msproto_ms_proto_msgTypes,
	}.Build()
	File_msproto_ms_proto = out.File
	file_msproto_ms_proto_rawDesc = nil
	file_msproto_ms_proto_goTypes = nil
	file_msproto_ms_proto_depIdxs = nil
}
