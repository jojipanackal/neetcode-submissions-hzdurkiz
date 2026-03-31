class Solution:

    def encode(self, strs: List[str]) -> str:
        # s = ["neet","code","love","you"]
        # result = "4#neet4#code4#love3#you"
        result = ""
        for s in strs:
            size = str(len(s))
            result += size + "#" + s
        return result

    def decode(self, s: str) -> List[str]:
        # s = "4#neet4#code4#love3#you"
        #      012345678
        # ans = ["neet","code","love","you"]
        ans = []
        i = 0
        while i < len(s):
            size = ""
            while s[i] != "#":
                size += s[i]
                i += 1
            size = int(size)
            i += 1
            ans.append(s[i:size+i])
            i += size
        return ans


