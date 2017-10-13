package handler

import (
	"thrift/proto/req"
	"thrift/proto/resp"
)
type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}


// <pre>
// 描述：申请加好友
// 参数：申请加好友参数
// 结果：返回申请结果
// 认证方式:Token
// </pre>
//
// Parameters:
//  - Req
func (p *UserHandler) AddFriend(req *req.AddFriendReq) (r *resp.AddFriendResp, err error) {
	return &resp.AddFriendResp{},nil
}

