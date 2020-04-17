// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccountType int32

const (
	AccountType_RegularUser AccountType = 0
	AccountType_Startup     AccountType = 1
	AccountType_Investor    AccountType = 2
)

var AccountType_name = map[int32]string{
	0: "RegularUser",
	1: "Startup",
	2: "Investor",
}

var AccountType_value = map[string]int32{
	"RegularUser": 0,
	"Startup":     1,
	"Investor":    2,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}

func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

type TeamAccountType int32

const (
	TeamAccountType_StartupTeam  TeamAccountType = 0
	TeamAccountType_InvestorTeam TeamAccountType = 1
)

var TeamAccountType_name = map[int32]string{
	0: "StartupTeam",
	1: "InvestorTeam",
}

var TeamAccountType_value = map[string]int32{
	"StartupTeam":  0,
	"InvestorTeam": 1,
}

func (x TeamAccountType) String() string {
	return proto.EnumName(TeamAccountType_name, int32(x))
}

func (TeamAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

type Address struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Longitude            string               `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             string               `protobuf:"bytes,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	City                 string               `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	State                string               `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Country              string               `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode              string               `protobuf:"bytes,7,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	Street               string               `protobuf:"bytes,8,opt,name=street,proto3" json:"street,omitempty"`
	BuildingNumber       string               `protobuf:"bytes,9,opt,name=building_number,json=buildingNumber,proto3" json:"building_number,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,12,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *Address) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetBuildingNumber() string {
	if m != nil {
		return m.BuildingNumber
	}
	return ""
}

func (m *Address) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Address) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Address) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type Education struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	School               string               `protobuf:"bytes,2,opt,name=school,proto3" json:"school,omitempty"`
	Degree               string               `protobuf:"bytes,3,opt,name=degree,proto3" json:"degree,omitempty"`
	FieldOfStudy         string               `protobuf:"bytes,4,opt,name=field_of_study,json=fieldOfStudy,proto3" json:"field_of_study,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	CurrentlyAttending   bool                 `protobuf:"varint,7,opt,name=currently_attending,json=currentlyAttending,proto3" json:"currently_attending,omitempty"`
	Gpa                  float32              `protobuf:"fixed32,8,opt,name=gpa,proto3" json:"gpa,omitempty"`
	Activities           string               `protobuf:"bytes,9,opt,name=activities,proto3" json:"activities,omitempty"`
	Societies            string               `protobuf:"bytes,10,opt,name=societies,proto3" json:"societies,omitempty"`
	Description          string               `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	MediaId              *Media               `protobuf:"bytes,15,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Education) Reset()         { *m = Education{} }
func (m *Education) String() string { return proto.CompactTextString(m) }
func (*Education) ProtoMessage()    {}
func (*Education) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *Education) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Education.Unmarshal(m, b)
}
func (m *Education) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Education.Marshal(b, m, deterministic)
}
func (m *Education) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Education.Merge(m, src)
}
func (m *Education) XXX_Size() int {
	return xxx_messageInfo_Education.Size(m)
}
func (m *Education) XXX_DiscardUnknown() {
	xxx_messageInfo_Education.DiscardUnknown(m)
}

var xxx_messageInfo_Education proto.InternalMessageInfo

func (m *Education) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Education) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *Education) GetDegree() string {
	if m != nil {
		return m.Degree
	}
	return ""
}

func (m *Education) GetFieldOfStudy() string {
	if m != nil {
		return m.FieldOfStudy
	}
	return ""
}

func (m *Education) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *Education) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *Education) GetCurrentlyAttending() bool {
	if m != nil {
		return m.CurrentlyAttending
	}
	return false
}

func (m *Education) GetGpa() float32 {
	if m != nil {
		return m.Gpa
	}
	return 0
}

func (m *Education) GetActivities() string {
	if m != nil {
		return m.Activities
	}
	return ""
}

func (m *Education) GetSocieties() string {
	if m != nil {
		return m.Societies
	}
	return ""
}

func (m *Education) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Education) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Education) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Education) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Education) GetMediaId() *Media {
	if m != nil {
		return m.MediaId
	}
	return nil
}

type Media struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DocumentLinks        []string             `protobuf:"bytes,2,rep,name=document_links,json=documentLinks,proto3" json:"document_links,omitempty"`
	PhotoLinks           []string             `protobuf:"bytes,3,rep,name=photo_links,json=photoLinks,proto3" json:"photo_links,omitempty"`
	VideoLinks           []string             `protobuf:"bytes,4,rep,name=video_links,json=videoLinks,proto3" json:"video_links,omitempty"`
	PresentationLinks    []string             `protobuf:"bytes,5,rep,name=presentation_links,json=presentationLinks,proto3" json:"presentation_links,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}
func (*Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{2}
}

func (m *Media) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Media.Unmarshal(m, b)
}
func (m *Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Media.Marshal(b, m, deterministic)
}
func (m *Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Media.Merge(m, src)
}
func (m *Media) XXX_Size() int {
	return xxx_messageInfo_Media.Size(m)
}
func (m *Media) XXX_DiscardUnknown() {
	xxx_messageInfo_Media.DiscardUnknown(m)
}

var xxx_messageInfo_Media proto.InternalMessageInfo

func (m *Media) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Media) GetDocumentLinks() []string {
	if m != nil {
		return m.DocumentLinks
	}
	return nil
}

func (m *Media) GetPhotoLinks() []string {
	if m != nil {
		return m.PhotoLinks
	}
	return nil
}

func (m *Media) GetVideoLinks() []string {
	if m != nil {
		return m.VideoLinks
	}
	return nil
}

func (m *Media) GetPresentationLinks() []string {
	if m != nil {
		return m.PresentationLinks
	}
	return nil
}

func (m *Media) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Media) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Media) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Subscriptions struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionName     string               `protobuf:"bytes,2,opt,name=subscription_name,json=subscriptionName,proto3" json:"subscription_name,omitempty"`
	SubscriptionStatus   string               `protobuf:"bytes,3,opt,name=subscription_status,json=subscriptionStatus,proto3" json:"subscription_status,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	AccessType           string               `protobuf:"bytes,6,opt,name=access_type,json=accessType,proto3" json:"access_type,omitempty"`
	IsActive             bool                 `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Subscriptions) Reset()         { *m = Subscriptions{} }
func (m *Subscriptions) String() string { return proto.CompactTextString(m) }
func (*Subscriptions) ProtoMessage()    {}
func (*Subscriptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{3}
}

func (m *Subscriptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscriptions.Unmarshal(m, b)
}
func (m *Subscriptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscriptions.Marshal(b, m, deterministic)
}
func (m *Subscriptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscriptions.Merge(m, src)
}
func (m *Subscriptions) XXX_Size() int {
	return xxx_messageInfo_Subscriptions.Size(m)
}
func (m *Subscriptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscriptions.DiscardUnknown(m)
}

var xxx_messageInfo_Subscriptions proto.InternalMessageInfo

func (m *Subscriptions) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Subscriptions) GetSubscriptionName() string {
	if m != nil {
		return m.SubscriptionName
	}
	return ""
}

func (m *Subscriptions) GetSubscriptionStatus() string {
	if m != nil {
		return m.SubscriptionStatus
	}
	return ""
}

func (m *Subscriptions) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *Subscriptions) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *Subscriptions) GetAccessType() string {
	if m != nil {
		return m.AccessType
	}
	return ""
}

func (m *Subscriptions) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Subscriptions) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Subscriptions) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Subscriptions) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type SocialMedia struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GithubUrl            uint32               `protobuf:"varint,2,opt,name=github_url,json=githubUrl,proto3" json:"github_url,omitempty"`
	WebsiteUrl           string               `protobuf:"bytes,3,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	FacebookUrl          string               `protobuf:"bytes,4,opt,name=facebook_url,json=facebookUrl,proto3" json:"facebook_url,omitempty"`
	TwitterUrl           string               `protobuf:"bytes,5,opt,name=twitter_url,json=twitterUrl,proto3" json:"twitter_url,omitempty"`
	LinkedUrl            string               `protobuf:"bytes,6,opt,name=linked_url,json=linkedUrl,proto3" json:"linked_url,omitempty"`
	YoutubeUrl           string               `protobuf:"bytes,7,opt,name=youtube_url,json=youtubeUrl,proto3" json:"youtube_url,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SocialMedia) Reset()         { *m = SocialMedia{} }
func (m *SocialMedia) String() string { return proto.CompactTextString(m) }
func (*SocialMedia) ProtoMessage()    {}
func (*SocialMedia) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{4}
}

func (m *SocialMedia) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocialMedia.Unmarshal(m, b)
}
func (m *SocialMedia) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocialMedia.Marshal(b, m, deterministic)
}
func (m *SocialMedia) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocialMedia.Merge(m, src)
}
func (m *SocialMedia) XXX_Size() int {
	return xxx_messageInfo_SocialMedia.Size(m)
}
func (m *SocialMedia) XXX_DiscardUnknown() {
	xxx_messageInfo_SocialMedia.DiscardUnknown(m)
}

var xxx_messageInfo_SocialMedia proto.InternalMessageInfo

func (m *SocialMedia) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SocialMedia) GetGithubUrl() uint32 {
	if m != nil {
		return m.GithubUrl
	}
	return 0
}

func (m *SocialMedia) GetWebsiteUrl() string {
	if m != nil {
		return m.WebsiteUrl
	}
	return ""
}

func (m *SocialMedia) GetFacebookUrl() string {
	if m != nil {
		return m.FacebookUrl
	}
	return ""
}

func (m *SocialMedia) GetTwitterUrl() string {
	if m != nil {
		return m.TwitterUrl
	}
	return ""
}

func (m *SocialMedia) GetLinkedUrl() string {
	if m != nil {
		return m.LinkedUrl
	}
	return ""
}

func (m *SocialMedia) GetYoutubeUrl() string {
	if m != nil {
		return m.YoutubeUrl
	}
	return ""
}

func (m *SocialMedia) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SocialMedia) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *SocialMedia) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Details struct {
	IPOStatus            string   `protobuf:"bytes,1,opt,name=IPOStatus,proto3" json:"IPOStatus,omitempty"`
	CompanyType          string   `protobuf:"bytes,2,opt,name=CompanyType,proto3" json:"CompanyType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{5}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetIPOStatus() string {
	if m != nil {
		return m.IPOStatus
	}
	return ""
}

func (m *Details) GetCompanyType() string {
	if m != nil {
		return m.CompanyType
	}
	return ""
}

type Experience struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName          string               `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Title                string               `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	EmploymentType       string               `protobuf:"bytes,4,opt,name=employment_type,json=employmentType,proto3" json:"employment_type,omitempty"`
	Location             string               `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	IsCurrentJob         bool                 `protobuf:"varint,8,opt,name=is_current_job,json=isCurrentJob,proto3" json:"is_current_job,omitempty"`
	Headline             string               `protobuf:"bytes,9,opt,name=headline,proto3" json:"headline,omitempty"`
	Description          string               `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	MediaId              *Media               `protobuf:"bytes,11,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Experience) Reset()         { *m = Experience{} }
func (m *Experience) String() string { return proto.CompactTextString(m) }
func (*Experience) ProtoMessage()    {}
func (*Experience) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{6}
}

func (m *Experience) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experience.Unmarshal(m, b)
}
func (m *Experience) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experience.Marshal(b, m, deterministic)
}
func (m *Experience) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experience.Merge(m, src)
}
func (m *Experience) XXX_Size() int {
	return xxx_messageInfo_Experience.Size(m)
}
func (m *Experience) XXX_DiscardUnknown() {
	xxx_messageInfo_Experience.DiscardUnknown(m)
}

var xxx_messageInfo_Experience proto.InternalMessageInfo

func (m *Experience) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Experience) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Experience) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Experience) GetEmploymentType() string {
	if m != nil {
		return m.EmploymentType
	}
	return ""
}

func (m *Experience) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Experience) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *Experience) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *Experience) GetIsCurrentJob() bool {
	if m != nil {
		return m.IsCurrentJob
	}
	return false
}

func (m *Experience) GetHeadline() string {
	if m != nil {
		return m.Headline
	}
	return ""
}

func (m *Experience) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Experience) GetMediaId() *Media {
	if m != nil {
		return m.MediaId
	}
	return nil
}

func (m *Experience) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Experience) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Experience) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type Investment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName          string   `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Industry             string   `protobuf:"bytes,3,opt,name=industry,proto3" json:"industry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Investment) Reset()         { *m = Investment{} }
func (m *Investment) String() string { return proto.CompactTextString(m) }
func (*Investment) ProtoMessage()    {}
func (*Investment) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{7}
}

func (m *Investment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Investment.Unmarshal(m, b)
}
func (m *Investment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Investment.Marshal(b, m, deterministic)
}
func (m *Investment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Investment.Merge(m, src)
}
func (m *Investment) XXX_Size() int {
	return xxx_messageInfo_Investment.Size(m)
}
func (m *Investment) XXX_DiscardUnknown() {
	xxx_messageInfo_Investment.DiscardUnknown(m)
}

var xxx_messageInfo_Investment proto.InternalMessageInfo

func (m *Investment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Investment) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Investment) GetIndustry() string {
	if m != nil {
		return m.Industry
	}
	return ""
}

func init() {
	proto.RegisterEnum("AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("TeamAccountType", TeamAccountType_name, TeamAccountType_value)
	proto.RegisterType((*Address)(nil), "Address")
	proto.RegisterType((*Education)(nil), "Education")
	proto.RegisterType((*Media)(nil), "Media")
	proto.RegisterType((*Subscriptions)(nil), "Subscriptions")
	proto.RegisterType((*SocialMedia)(nil), "SocialMedia")
	proto.RegisterType((*Details)(nil), "Details")
	proto.RegisterType((*Experience)(nil), "Experience")
	proto.RegisterType((*Investment)(nil), "Investment")
}

func init() {
	proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0)
}

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 1116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x8e, 0x24, 0x4b, 0xa2, 0x86, 0xb2, 0xad, 0xec, 0x2f, 0xf8, 0x81, 0x71, 0x9b, 0xc6, 0x0e,
	0x52, 0xd4, 0x48, 0x11, 0x09, 0xe8, 0x1f, 0x14, 0xc9, 0x4d, 0xf9, 0x73, 0x48, 0xd1, 0x26, 0x2d,
	0x9d, 0x5c, 0x7a, 0x21, 0x48, 0xee, 0x48, 0xde, 0x86, 0xe4, 0x12, 0xdc, 0xa5, 0x13, 0xf5, 0x9d,
	0xfa, 0x00, 0xc9, 0xa1, 0xaf, 0xd2, 0xa7, 0xe8, 0xa1, 0xe8, 0xa5, 0xd8, 0xd9, 0xa5, 0x2c, 0xab,
	0x05, 0x2c, 0x0b, 0xbd, 0xe4, 0xa6, 0xf9, 0xe6, 0x1b, 0x9a, 0xdc, 0xef, 0x9b, 0x99, 0x35, 0xb0,
	0xb2, 0x92, 0x5a, 0x4e, 0x52, 0x99, 0xe7, 0xb2, 0x18, 0x53, 0x70, 0xf0, 0x70, 0x2e, 0xf4, 0x69,
	0x9d, 0x8c, 0x53, 0x99, 0x4f, 0x44, 0x31, 0x93, 0x49, 0x26, 0xdf, 0xca, 0x12, 0x8b, 0x09, 0xa5,
	0xd3, 0xfb, 0x73, 0x2c, 0xee, 0xcf, 0x65, 0x95, 0x4f, 0x64, 0xa9, 0x85, 0x2c, 0xd4, 0xc4, 0x04,
	0xae, 0xf6, 0x9b, 0x95, 0x5a, 0x42, 0x92, 0x7a, 0x36, 0x51, 0x55, 0x3a, 0x99, 0x4b, 0x39, 0xcf,
	0xf0, 0x1c, 0xd3, 0x22, 0x47, 0xa5, 0xe3, 0xbc, 0xb4, 0x85, 0x77, 0x7e, 0xeb, 0x40, 0x7f, 0xca,
	0x79, 0x85, 0x4a, 0xb1, 0x23, 0x68, 0x0b, 0x1e, 0xb4, 0x0e, 0x5b, 0xc7, 0xdd, 0x47, 0xd7, 0xdf,
	0xbf, 0xbb, 0xb9, 0x0b, 0x3e, 0xeb, 0x8b, 0x42, 0xe3, 0x1c, 0xab, 0xe3, 0x56, 0xd8, 0x16, 0x9c,
	0x7d, 0x0c, 0x83, 0x4c, 0x16, 0x73, 0xa1, 0x6b, 0x8e, 0x41, 0xfb, 0xb0, 0x75, 0x3c, 0x08, 0xcf,
	0x01, 0x76, 0x00, 0x5e, 0x16, 0x6b, 0x9b, 0xec, 0x50, 0x72, 0x19, 0x33, 0x06, 0x3b, 0xa9, 0xd0,
	0x8b, 0x60, 0x87, 0x70, 0xfa, 0xcd, 0x6e, 0x40, 0x57, 0xe9, 0x58, 0x63, 0xd0, 0x25, 0xd0, 0x06,
	0x2c, 0x80, 0x7e, 0x2a, 0xeb, 0x42, 0x57, 0x8b, 0xa0, 0x47, 0x78, 0x13, 0x9a, 0xcc, 0x2f, 0xa2,
	0x7c, 0x2c, 0x39, 0x06, 0x7d, 0x9b, 0x71, 0x21, 0xfb, 0x3f, 0xf4, 0x94, 0xae, 0x10, 0x75, 0xe0,
	0x51, 0xc2, 0x45, 0xec, 0x33, 0xd8, 0x4f, 0x6a, 0x91, 0x71, 0x51, 0xcc, 0xa3, 0xa2, 0xce, 0x13,
	0xac, 0x82, 0x01, 0x11, 0xf6, 0x1a, 0xf8, 0x39, 0xa1, 0xec, 0x01, 0x40, 0x5a, 0x61, 0xac, 0x91,
	0x47, 0xb1, 0x0e, 0xe0, 0xb0, 0x75, 0xec, 0x7f, 0x71, 0x30, 0xb6, 0xa7, 0x37, 0x6e, 0x4e, 0x6f,
	0xfc, 0xb2, 0x39, 0xbd, 0x70, 0xe0, 0xd8, 0x53, 0x6d, 0x4a, 0xeb, 0x92, 0x37, 0xa5, 0xfe, 0xe5,
	0xa5, 0x8e, 0x6d, 0x4b, 0x39, 0x66, 0xe8, 0x4a, 0x87, 0x97, 0x97, 0x3a, 0xf6, 0x54, 0x3f, 0xec,
	0xbd, 0x7f, 0x77, 0xb3, 0xed, 0xb5, 0xee, 0xfc, 0xb1, 0x03, 0x83, 0xa7, 0xbc, 0x4e, 0x63, 0xe3,
	0x89, 0x4d, 0x24, 0x34, 0x47, 0x95, 0x9e, 0x4a, 0x99, 0x39, 0xfd, 0x5c, 0x64, 0x70, 0x8e, 0xf3,
	0x0a, 0x1b, 0xe9, 0x5c, 0xc4, 0xee, 0xc2, 0xde, 0x4c, 0x60, 0xc6, 0x23, 0x39, 0x8b, 0x94, 0xae,
	0x79, 0x23, 0xe1, 0x90, 0xd0, 0x17, 0xb3, 0x13, 0x83, 0x99, 0x2f, 0x51, 0x3a, 0xae, 0x74, 0xc4,
	0x1b, 0x3d, 0x2f, 0xf9, 0x12, 0x62, 0x3f, 0x31, 0x7a, 0x7f, 0x0d, 0x1e, 0x16, 0xdc, 0x16, 0xf6,
	0x2e, 0x2d, 0xec, 0x63, 0xc1, 0xa9, 0x6c, 0x02, 0xff, 0x4b, 0xeb, 0xaa, 0xc2, 0x42, 0x67, 0x8b,
	0x28, 0xd6, 0x1a, 0x0b, 0x23, 0x27, 0x19, 0xc3, 0x0b, 0xd9, 0x32, 0x35, 0x6d, 0x32, 0x6c, 0x04,
	0x9d, 0x79, 0x19, 0x93, 0x41, 0xda, 0xa1, 0xf9, 0xc9, 0x3e, 0x01, 0x88, 0x53, 0x2d, 0xce, 0x84,
	0x16, 0xa8, 0x9c, 0x31, 0x56, 0x10, 0xe3, 0x76, 0x25, 0x53, 0x81, 0x94, 0x06, 0xeb, 0xf6, 0x25,
	0xc0, 0x0e, 0xc1, 0xe7, 0xa8, 0xd2, 0x4a, 0x50, 0x3b, 0x92, 0xf0, 0x83, 0x70, 0x15, 0x5a, 0x33,
	0xd5, 0x70, 0x7b, 0x53, 0xed, 0x6e, 0x6f, 0xaa, 0xbd, 0x2b, 0x98, 0x8a, 0x1d, 0x81, 0x97, 0x23,
	0x17, 0x71, 0x24, 0x78, 0xb0, 0x4f, 0x85, 0xbd, 0xf1, 0xf7, 0x06, 0x08, 0xfb, 0x84, 0x3f, 0xe3,
	0x4b, 0xdf, 0xfd, 0xd5, 0x86, 0x2e, 0xa5, 0x36, 0xf1, 0xdc, 0xa7, 0xb0, 0xc7, 0x65, 0x5a, 0xe7,
	0x58, 0xe8, 0x28, 0x13, 0xc5, 0x6b, 0x15, 0xb4, 0x0f, 0x3b, 0xc7, 0x83, 0x70, 0xb7, 0x41, 0xbf,
	0x33, 0x20, 0xbb, 0x0d, 0x7e, 0x79, 0x2a, 0xb5, 0x74, 0x9c, 0x0e, 0x71, 0x80, 0xa0, 0x25, 0xe1,
	0x4c, 0x70, 0x6c, 0x08, 0x3b, 0x96, 0x40, 0x90, 0x25, 0xdc, 0x37, 0x93, 0x15, 0x15, 0x16, 0x9a,
	0xfa, 0xc1, 0xf1, 0xba, 0xc4, 0xbb, 0xbe, 0x9a, 0xb1, 0xf4, 0x8b, 0x02, 0xf5, 0xae, 0x28, 0xd0,
	0xca, 0x29, 0xf7, 0xaf, 0x72, 0xca, 0x17, 0xb5, 0xf5, 0xae, 0xa0, 0xed, 0xf2, 0xf4, 0xff, 0xec,
	0xc0, 0xee, 0x49, 0x9d, 0x2c, 0xad, 0xb6, 0xd1, 0xf0, 0xfe, 0x1c, 0xae, 0xab, 0x95, 0x9a, 0xa8,
	0x88, 0xf3, 0x66, 0x88, 0x8f, 0x56, 0x13, 0xcf, 0xe3, 0x9c, 0xda, 0xeb, 0x02, 0xd9, 0xcc, 0xe6,
	0x5a, 0xb9, 0xd9, 0xc0, 0x56, 0x53, 0x27, 0x94, 0x59, 0x9b, 0x00, 0x3b, 0xdb, 0x4e, 0x80, 0xee,
	0xe6, 0x13, 0xe0, 0x36, 0xf8, 0x71, 0x9a, 0xa2, 0x52, 0x91, 0x5e, 0x94, 0xe8, 0x96, 0x05, 0x58,
	0xe8, 0xe5, 0xa2, 0x44, 0xf6, 0x11, 0x0c, 0x84, 0x8a, 0xa8, 0xa1, 0xd1, 0x0d, 0x06, 0x4f, 0xa8,
	0x29, 0xc5, 0x6b, 0xda, 0x7b, 0xdb, 0x6b, 0x3f, 0xd8, 0x5e, 0x7b, 0xd8, 0x46, 0xfb, 0x5f, 0x3b,
	0xe0, 0x9f, 0xc8, 0x54, 0xc4, 0xd9, 0xc6, 0xfd, 0x77, 0x0b, 0xc0, 0x5e, 0x10, 0xa2, 0xba, 0xb2,
	0x73, 0x7f, 0x37, 0x1c, 0x58, 0xe4, 0x55, 0x95, 0x99, 0x83, 0x7c, 0x83, 0x89, 0x12, 0x1a, 0x29,
	0x6f, 0x35, 0x06, 0x07, 0x19, 0xc2, 0x11, 0x0c, 0x67, 0x71, 0x8a, 0x89, 0x94, 0xaf, 0x89, 0x61,
	0x37, 0x80, 0xdf, 0x60, 0xee, 0x19, 0xfa, 0x8d, 0xd0, 0x1a, 0x2b, 0x62, 0xd8, 0x8d, 0x0e, 0x0e,
	0x32, 0x84, 0x5b, 0x00, 0xa6, 0x1b, 0x91, 0x53, 0xbe, 0xe7, 0xee, 0x0e, 0x84, 0xb8, 0xfa, 0x85,
	0xac, 0x75, 0x9d, 0xd8, 0x77, 0xb0, 0xfb, 0x1d, 0x1c, 0x64, 0x08, 0x1f, 0xac, 0x5e, 0x3f, 0x42,
	0xff, 0x09, 0xea, 0x58, 0x64, 0xb4, 0x50, 0x9e, 0xfd, 0xf0, 0xc2, 0x36, 0x0c, 0x29, 0x36, 0x08,
	0xcf, 0x01, 0xb3, 0x50, 0x1e, 0xcb, 0xbc, 0x8c, 0x8b, 0x85, 0x71, 0xaf, 0xeb, 0xcc, 0x55, 0x68,
	0xf9, 0xc8, 0xdf, 0x77, 0x00, 0x9e, 0xbe, 0x2d, 0xb1, 0x12, 0x58, 0xa4, 0xb8, 0x89, 0x03, 0x8e,
	0x60, 0x98, 0xda, 0x07, 0xad, 0xb6, 0xbd, 0xef, 0x30, 0xea, 0xf8, 0x1b, 0xd0, 0xd5, 0x42, 0x67,
	0xcd, 0xfe, 0xb7, 0x81, 0xb9, 0x41, 0x61, 0x5e, 0x66, 0x72, 0x41, 0xc3, 0x9b, 0x1a, 0xcd, 0xaa,
	0xbf, 0x77, 0x0e, 0x53, 0xb3, 0x99, 0xcb, 0x9f, 0xb4, 0xd7, 0x10, 0xa7, 0xfe, 0x32, 0x5e, 0x9b,
	0x0d, 0xbd, 0x6d, 0x67, 0x43, 0x7f, 0xf3, 0xd9, 0x70, 0x17, 0xf6, 0x84, 0x8a, 0xdc, 0x2d, 0x20,
	0xfa, 0x59, 0x26, 0xe4, 0x18, 0x2f, 0x1c, 0x0a, 0xf5, 0xd8, 0x82, 0xdf, 0xca, 0xc4, 0xbc, 0xf3,
	0x29, 0xc6, 0x3c, 0x13, 0x05, 0xba, 0xf5, 0xbf, 0x8c, 0xd7, 0xd7, 0x3b, 0xfc, 0x73, 0xbd, 0xaf,
	0x6e, 0x4b, 0xff, 0x5f, 0xb7, 0xe5, 0x07, 0x77, 0x03, 0x58, 0x3a, 0xec, 0x0c, 0xe0, 0x59, 0x71,
	0x86, 0x4a, 0x1b, 0x7d, 0xff, 0x23, 0x83, 0x1d, 0x80, 0x27, 0x0a, 0x5e, 0x2b, 0x73, 0xb3, 0x77,
	0xff, 0x1e, 0x34, 0x71, 0xf3, 0x77, 0xef, 0x3d, 0x00, 0x7f, 0x9a, 0xd2, 0x7d, 0x9f, 0x4c, 0xb5,
	0x0f, 0x7e, 0x88, 0xf3, 0x3a, 0x8b, 0xab, 0x57, 0x0a, 0xab, 0xd1, 0x35, 0xe6, 0x43, 0xff, 0xc4,
	0x78, 0xa3, 0x2e, 0x47, 0x2d, 0x36, 0x04, 0xcf, 0xbe, 0xa4, 0xac, 0x46, 0xed, 0x7b, 0x5f, 0xc1,
	0xfe, 0x4b, 0x8c, 0xf3, 0xb5, 0x72, 0xc7, 0x36, 0x99, 0xd1, 0x35, 0x36, 0x82, 0x61, 0x53, 0x41,
	0x48, 0xeb, 0x51, 0xff, 0xa7, 0x6e, 0x2e, 0x39, 0x66, 0x49, 0x8f, 0x0e, 0xe6, 0xcb, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x8c, 0xd8, 0x40, 0x9b, 0x0d, 0x00, 0x00,
}
