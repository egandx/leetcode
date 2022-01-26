/*
 * @lc app=leetcode.cn id=1 lang=cpp
 *
 * [1] 两数之和
 */

// @lc code=start

//枚举数组中的每一个数 x，寻找数组中是否存在 target - x

//暴力枚举
// class Solution {
// public:
//     vector<int> twoSum(vector<int>& nums, int target) {
//         int n = nums.size();
//         for (int i = 0; i < n; ++i) {
//             for (int j = i + 1; j < n; ++j) {
//                 if (nums[i] + nums[j] == target) {
//                     return {i, j};
//                 }
//             }
//         }
//         return {};
//     }
// };

//哈希表
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> hashtable;
        for (int i = 0; i < nums.size(); ++i) {
            if (hashtable.count(target - nums[i]) == 1) {
                return {hashtable[target - nums[i]], i};
            }
            hashtable[nums[i]] = i;
        }
        return {};
    }
};
// @lc code=end

