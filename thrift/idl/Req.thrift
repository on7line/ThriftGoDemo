include "Base.thrift"
namespace java com.qiknow.proto

/**
 * 加好友请求参数
 */
struct AddFriendReq{
	1: Base.ReqHeader header
	/**
     * 申请者用户ID
     */
    2: i64 requesterId
	/**
     * 被申请者用户ID
     */
    3: i64 approverId
	/**
     * 加好友备注信息
     */
    4: string comments
}
