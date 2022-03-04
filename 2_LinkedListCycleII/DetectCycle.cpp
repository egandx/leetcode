/**
 * Leetcode 142 medium
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        // First go to determine whether there is a ring
        ListNode *fast = head;
        ListNode *slow = head;

        while(fast && fast->next != NULL && fast->next->next != NULL){
            fast = fast->next->next;
            slow = slow->next;
            if(fast == slow ){
                break;
            }
        }
        
        //hasn't cycle
        if(fast == NULL||fast->next == NULL || fast->next->next ==NULL ){
            return NULL;
        }

        //has cycle
        ListNode *h = head;

        while(h != slow){
            h = h->next;
            slow = slow->next;
        }
        return h;
    }
};
