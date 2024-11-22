



def find_max_nums(nums):
    if not nums:
        return 0
    res_list = []
    some_list = []
    list_len = len(nums)
    for i in range(list_len):
        cur_number = nums[i]
        if i < list_len - 1:
            if cur_number > 0 and cur_number == nums[i+1] or cur_number== nums[i-1] and cur_number > 0:
                some_list.append(cur_number)
            else:
                if some_list:
                    res_list.append(some_list)
                    some_list = []
        elif i == list_len -1:
            if cur_number == nums[-1]:
                some_list.append(cur_number)
                res_list.append(some_list)
    return sum(max(res_list)) if res_list else 0              

def find_max_nums2(nums):
    if not nums:
        return 0
    else:
        def convert_list_string(some_nums):
            d = " ".join([str(x) for x in nums if str(x).isdigit()])
            return d 

        string_list=  convert_list_string(nums)
        splited = string_list.split("0")
        result = []
        for split in splited:
            split = [int(x) for x in split if x.isdigit()]
            if split:
                result.append(split)
        
        return sum(max(result)) if result else 0
    

def find_max_nums3(nums):
    best = 0
    current = 0
    for i in nums:
        if i>0: 
            current+= 1
            best = max(best, current)
        else:
            current = 0
    return best

assert find_max_nums3([1,1,0,0,1,1,0]) == find_max_nums([1,1,0,0,1,1, 0])
assert find_max_nums3([1,1,0,0,1,1,1]) == find_max_nums([1,1,0,0,1,1,1])
assert find_max_nums3([1,1,1,1,1,1]) == find_max_nums([1,1,1,1,1,1])
assert find_max_nums3([0,0,0]) == find_max_nums([0,0,0])
assert find_max_nums3([]) == find_max_nums([])