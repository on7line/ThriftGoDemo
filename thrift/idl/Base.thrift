namespace java com.qiknow.proto

/**
 * 请求头信息
 */
struct ReqHeader{
	/**
	 * Token
	 */
	1: string token
}

/**
 * 用户信息
 */
struct UserInfo{
	/**
	 * 头像
	 */
	1: string headImgPath
	/**
	 * 用户姓名
	 */
	2: string name
	/**
	 * 昵称
	 */
	3: string nickname
	/**
	 * 性别(1男，2女，0未知)
	 */
	4: i32 gender
	/**
	 * 电话
	 */
	5: string phoneNo
	/**
	 * 用户的标签
	 */
	6: list<string> tags
	/**
	 * 用户ID
	 */
	7: i64 userId
	/**
	 * 用户关系。1：其他关系，2：我的关注者，3：关注我的人，4：互相关注，5：好友关系，
	 */
	8: i32 relationship
}
