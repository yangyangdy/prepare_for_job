package main
//给你一个可装载重量为W的背包和N个物品，每个物品有重量和价值两个属性。
//其中第i个物品的重量为wt[i]，价值为val[i]，现在让你用这个背包装物品，最多能装的价值是多少？
/*
N = 3, W = 4
wt = [2, 1, 3]
val = [4, 2, 3]
 */
/*
动归框架
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 择优(选择1，选择2...)
 */
/*
dp[i][w]的定义如下：对于前i个物品，当前背包的容量为w，这种情况下可以装的最大价值是dp[i][w]。
比如说，如果 dp[3][5] = 6，其含义为：对于给定的一系列物品中，若只对前 3 个物品进行选择，当背包容量为 5 时，最多可以装下的价值为 6。

根据这个定义，我们想求的最终答案就是dp[N][W]。base case 就是dp[0][..] = dp[..][0] = 0，因为没有物品或者背包没有空间的时候，能装的最大价值就是 0。

细化上面的框架：
int dp[N+1][W+1]
dp[0][..] = 0
dp[..][0] = 0

for i in [1..N]:
    for w in [1..W]://背包容量变化的意义是什么
        dp[i][w] = max(
            把物品 i 装进背包, //那么状态转移为dp[i-1][w-wt[i-1]] + val[i-1]
            不把物品 i 装进背包 //那么状态转移为dp[i-1][w]
        )
return dp[N][W]

//动归的状态转移方程中的每个状态就应该是问题的解，问题只是其中最后状态的解
 */
func main() {

}

