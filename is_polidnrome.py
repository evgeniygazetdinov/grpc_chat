def is_polindrome(s1):
    l=0
    r = len(s1) -1

    while l <r:
        if s1[l] != s1[r]:
            return False
        
        l+=1
        r-= 1
    return True


print(is_polindrome("дед"))
print(is_polindrome("доход"))