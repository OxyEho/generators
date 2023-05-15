from math import sqrt
from decimal import getcontext, Decimal

def my_sqrt(num: int):
        with open("sqrt.log", "x", encoding='utf-8') as f:
                getcontext().prec = 70000000
                n = Decimal(num)
                x = n
                offset = 0
                prev_str = ""
                is_first = True
                while 1:
                        x = (x + n / x) / Decimal(2)
                        cur_str = str(x)
                        for i in range(offset, min(len(prev_str), len(cur_str))):
                                offset += 1
                                if prev_str[i] != cur_str[i]:
                                        offset = i
                                        f.write(f"посчитано знаков: {offset}\n")
                                        print(f"посчитано знаков: {offset}")
                                        break
                        else:
                                if is_first:
                                        is_first = False
                                else:
                                        f.write(f"посчитано знаков: {len(cur_str)-1}")
                                        print(f"посчитано знаков: {len(cur_str)-1}")
                                        return
                        prev_str = cur_str

def main():
        print(sqrt(11))
        my_sqrt(11)

if __name__ == "__main__":
        main()