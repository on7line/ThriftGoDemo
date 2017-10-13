// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package base

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// 请求头信息
// 
// Attributes:
//  - Token: Token
type ReqHeader struct {
  Token string `thrift:"token,1" db:"token" json:"token"`
}

func NewReqHeader() *ReqHeader {
  return &ReqHeader{}
}


func (p *ReqHeader) GetToken() string {
  return p.Token
}
func (p *ReqHeader) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ReqHeader)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Token = v
}
  return nil
}

func (p *ReqHeader) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ReqHeader"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ReqHeader) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err) }
  if err := oprot.WriteString(string(p.Token)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err) }
  return err
}

func (p *ReqHeader) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ReqHeader(%+v)", *p)
}

// 用户信息
// 
// Attributes:
//  - HeadImgPath: 头像
//  - Name: 用户姓名
//  - Nickname: 昵称
//  - Gender: 性别(1男，2女，0未知)
//  - PhoneNo: 电话
//  - Tags: 用户的标签
//  - UserId: 用户ID
//  - Relationship: 用户关系。1：其他关系，2：我的关注者，3：关注我的人，4：互相关注，5：好友关系，
type UserInfo struct {
  HeadImgPath string `thrift:"headImgPath,1" db:"headImgPath" json:"headImgPath"`
  Name string `thrift:"name,2" db:"name" json:"name"`
  Nickname string `thrift:"nickname,3" db:"nickname" json:"nickname"`
  Gender int32 `thrift:"gender,4" db:"gender" json:"gender"`
  PhoneNo string `thrift:"phoneNo,5" db:"phoneNo" json:"phoneNo"`
  Tags []string `thrift:"tags,6" db:"tags" json:"tags"`
  UserId int64 `thrift:"userId,7" db:"userId" json:"userId"`
  Relationship int32 `thrift:"relationship,8" db:"relationship" json:"relationship"`
}

func NewUserInfo() *UserInfo {
  return &UserInfo{}
}


func (p *UserInfo) GetHeadImgPath() string {
  return p.HeadImgPath
}

func (p *UserInfo) GetName() string {
  return p.Name
}

func (p *UserInfo) GetNickname() string {
  return p.Nickname
}

func (p *UserInfo) GetGender() int32 {
  return p.Gender
}

func (p *UserInfo) GetPhoneNo() string {
  return p.PhoneNo
}

func (p *UserInfo) GetTags() []string {
  return p.Tags
}

func (p *UserInfo) GetUserId() int64 {
  return p.UserId
}

func (p *UserInfo) GetRelationship() int32 {
  return p.Relationship
}
func (p *UserInfo) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 6:
      if err := p.ReadField6(iprot); err != nil {
        return err
      }
    case 7:
      if err := p.ReadField7(iprot); err != nil {
        return err
      }
    case 8:
      if err := p.ReadField8(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *UserInfo)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.HeadImgPath = v
}
  return nil
}

func (p *UserInfo)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *UserInfo)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Nickname = v
}
  return nil
}

func (p *UserInfo)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Gender = v
}
  return nil
}

func (p *UserInfo)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.PhoneNo = v
}
  return nil
}

func (p *UserInfo)  ReadField6(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.Tags =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.Tags = append(p.Tags, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *UserInfo)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  p.UserId = v
}
  return nil
}

func (p *UserInfo)  ReadField8(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 8: ", err)
} else {
  p.Relationship = v
}
  return nil
}

func (p *UserInfo) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("UserInfo"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField7(oprot); err != nil { return err }
    if err := p.writeField8(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *UserInfo) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("headImgPath", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:headImgPath: ", p), err) }
  if err := oprot.WriteString(string(p.HeadImgPath)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.headImgPath (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:headImgPath: ", p), err) }
  return err
}

func (p *UserInfo) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err) }
  return err
}

func (p *UserInfo) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("nickname", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:nickname: ", p), err) }
  if err := oprot.WriteString(string(p.Nickname)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.nickname (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:nickname: ", p), err) }
  return err
}

func (p *UserInfo) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("gender", thrift.I32, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:gender: ", p), err) }
  if err := oprot.WriteI32(int32(p.Gender)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.gender (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:gender: ", p), err) }
  return err
}

func (p *UserInfo) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("phoneNo", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:phoneNo: ", p), err) }
  if err := oprot.WriteString(string(p.PhoneNo)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.phoneNo (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:phoneNo: ", p), err) }
  return err
}

func (p *UserInfo) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("tags", thrift.LIST, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:tags: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRING, len(p.Tags)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Tags {
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:tags: ", p), err) }
  return err
}

func (p *UserInfo) writeField7(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("userId", thrift.I64, 7); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:userId: ", p), err) }
  if err := oprot.WriteI64(int64(p.UserId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.userId (7) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 7:userId: ", p), err) }
  return err
}

func (p *UserInfo) writeField8(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("relationship", thrift.I32, 8); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:relationship: ", p), err) }
  if err := oprot.WriteI32(int32(p.Relationship)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.relationship (8) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 8:relationship: ", p), err) }
  return err
}

func (p *UserInfo) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("UserInfo(%+v)", *p)
}
