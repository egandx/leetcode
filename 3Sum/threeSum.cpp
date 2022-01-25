class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        int n = nums.size();
        sort(nums.begin(),nums.end());
        vector<vector<int>> ans;
        if(n < 3) return ans;
        for(int i = 0; i < n; ++i){
            if(nums[i] > 0) return ans;
            if(i > 0 && nums[i] == nums[i-1])
                continue;
            int L = i+1;
            int R = n-1;
            while(L<R){
                if(nums[i] + nums[L] +nums[R] == 0){
                    ans.push_back({nums[i],nums[L],nums[R]});
                    while(L<R and nums[L]==nums[L+1]){
                        L=L+1;
                    }
                    while(L<R and nums[R]==nums[R-1]){
                        R=R-1;
                    }
                    L=L+1;
                    R=R-1;
                }
                else if(nums[i]+nums[L]+nums[R]>0){
                    R=R-1;
                }
                else{
                    L=L+1;
                }
            }
        }
        return ans;
    }
};
