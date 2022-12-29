class Solution:
    def levelOrder(self, root):
        r = []
        current = root
        d = deque([root])
        while len(d):
            c = d.popleft()
            r.append(c.data)
            if c.left:
                d.append(c.left)
            if c.right:
                d.append(c.right)
        return r

