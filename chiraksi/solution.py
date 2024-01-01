class ExceptionProxy(Exception):
    def __init__(self, msg, function):
        self.msg = msg
        self.function = function

def transform_exceptions(functions):
    exceptions = []

    for func in functions:
        try:
            func()
            exceptions.append(ExceptionProxy("ok!", func))
        except Exception as e:
            exceptions.append(ExceptionProxy(str(e), func))

    return exceptions
