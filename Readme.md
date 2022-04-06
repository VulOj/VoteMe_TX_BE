沈博宇2022腾讯春招笔试题

概述

设计一个后端投票系统VoteMe，该后端系统提供一些接口（见后文），客户端可以通过这些接口给指定的用户名投票，同时客户端也可以通过接口查询指定用户名的当前票数。

1. 该程序需要每 2 秒中生成一个随机的 ticket， ticket 由服务端产生，并提供相应的接口获取该 ticket
2. 每个 ticket 的有效时间是从服务端生成到下一次生成新的 ticket 为止，且每个 ticket 在投票请求中存在一个使用上限
3. 客户端需要每次传递 ticket才能投票，同时服务端还需要验证该 ticket 的有效性
4.在 ticket 的有效时间内，支持给任意用户投票，但不能超过ticket使用次数上限


接口设计
1. 投票
输入1: 指定的用户名（支持输入多个用户名）
输入2：ticket 行为：检查 ticket 是否合法，如果合法，则将指定的（多个）用户的票数+1
输出：投票是否成功
2. 查询票数
输入: 指定的用户名
输出：指定用户的当前票数
3. 获取票据 ticket
输入：空
输出：当前合法的ticket

要求
1. Golang语言实现，需要使用 GraphQL 设计相应的 API
2. 投票数据需要持久化
3. 请尽可能地提升程序性能，注意程序可能存在的条件竞争、扩展性等问题