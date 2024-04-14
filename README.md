# Morris-go
Implementing Three Types of Morris Traverse with Go

## morris遍历的特点  时间o(n),空间o(1)
## 实现前中后序遍历，基本上是在**morris序的基础上调节**打印的顺序

- 先序遍历：在第一次遇到节点时打印
- 中序遍历：对于能来到自己两次的，在第二次打印；只能来到自己一次的，直接打印
- 后序遍历：打印时机是在第二次回到自己时，不过打印的是其左孩子的右边界，并且逆序打印






对于二叉树的非递归遍历，前序和后序使用迭代法比较方便记忆和书写，中序则是采用morris方便
- 前序迭代逻辑（中左右）：加入中节点，出栈，将右节点、左节点入栈，栈非空执行循环
- 后续迭代逻辑（左右中）：调整前序的左右节点入站顺序从右、左变为左、右，之后reverse数组，即可得到结果
- 中序morris遍历逻辑：记录mostright,如果mostright指向空，说明第一次来到，指回cur;如果mostright指向cur,第二次来到，加入res,并把指针指回空。记录时机：对与来到自己两次的，在第二次打印


```
// 前序
class Solution {
public:

    //前序，中左右
    //非递归遍历，加入中,出栈，加入res,同时按照右左顺序将节点加入栈
    
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> sk;
        if (!root) return res;
        sk.push(root);
        while(!sk.empty()) {
            TreeNode* x = sk.top();
            sk.pop();
            res.push_back(x->val);
            if (x->right) sk.push(x->right);
            if (x->left) sk.push(x->left);
        }
        return res;
    }
};
```

```
// 后序

class Solution {
public:
    // 左右中
    // 其实是变化前序的逆序列
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> sk;
        if (!root) return res;
        sk.push(root);
        while (!sk.empty()) {
            TreeNode* x = sk.top();
            sk.pop();
            res.push_back(x->val);
            if (x->left) sk.push(x->left);
            if (x->right) sk.push(x->right);
        }
        reverse(res.begin(),res.end());
        return res;
    }
};
```



```
// 中序
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        if (!root) return res;
        TreeNode* cur = root;
        while (cur) {
            TreeNode* mostRight = cur->left;
            if (!mostRight) {
                res.push_back(cur->val);
                cur = cur->right;
            }else {
                while (mostRight->right && mostRight->right != cur) {
                    mostRight = mostRight->right;
                }
                if (mostRight->right == nullptr) {
                    mostRight->right = cur;
                    cur = cur->left;
                }else {
                    mostRight->right = nullptr;
                    res.push_back(cur->val);
                    cur = cur->right;
                }
            }
        }
        return res;
    }
};

```
