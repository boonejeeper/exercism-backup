def find(search_list, value):
    result = find_rec(search_list, value, 0, len(search_list)-1, len(search_list)//2)
    print(f'result = {result}')
    return result

def find_rec(search_list, value, low, high, index):
    if low > high:
        raise ValueError('value not in array')


    print(f'low:{low} high:{high} index:{index} value:{value}')
    
    if int(search_list[index]) == value:
        print(f'returning {index}')
        return index

    if low == high:
        raise ValueError('value not in array')
    elif int(search_list[index]) < value:
        return find_rec(search_list, value, index+1, high, (index+1+high)//2)
    else:
        return find_rec(search_list, value, low, index, (low+index)//2)
