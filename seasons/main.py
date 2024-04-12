def getSeason(month):
    if month.lower() in ['december', 'january', 'february']:
        return 'summer'
    elif month.lower() in ['march', 'april', 'may']:
        return 'autumn'
    elif month.lower() in ['june', 'july', 'august']:
        return 'winter'
    elif month.lower() in ['september', 'october', 'november']:
        return 'spring'


inp = input()
print(getSeason(inp))