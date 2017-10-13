include "Req.thrift"
include "Resp.thrift"

namespace java com.qiknow.proto

/**<pre>
描述:账号服务，包括登录的、登出、获取验证码、修改密码、用户信息等相关功能接口
</pre> */	
service UserService{

	/**<pre>
	描述：申请加好友
	参数：申请加好友参数
	结果：返回申请结果
	认证方式:Token
	</pre> */
	Resp.AddFriendResp AddFriend(1: Req.AddFriendReq req)
	
}
