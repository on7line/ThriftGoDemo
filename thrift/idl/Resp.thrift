include "Base.thrift"
namespace java com.qiknow.proto

struct AddFriendResp {
	/**
	 * 状态码
	 */
	1: i32 status
	/**
	 * 返回结果描述
	 */
	2: string msg
}
