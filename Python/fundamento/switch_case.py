
number:int = 5

def check_number(num:int):
    match num:
        case 1:
            print("Num is one")
        case 2:
            print("Num is two")
        case 3:
            print("Num is three")
        case 4:
            print("Num is four")
        case 5:
            print("Num is five")
        case 6:
            print("Num is six")
        case _:
            print("Any")

check_number(number)
