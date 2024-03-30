"""
OBTER O ÚNICO NÚMERO QUE NÃO TEM PAR.

PS: Com valor contante de espaço em memória
"""

nums: list = [2,2,3,3,5,4,4]

def one_num(nums):
    xor = 0
    for num in nums:
        xor ^= num
    return xor

r = one_num(nums)
print(r)
