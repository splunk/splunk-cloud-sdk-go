import inspect

# Answers if the given value is a dict.
def isdict(value):
    return isinstance(value, dict)


# Answers if the given value is a callable.
def isfunction(value):
    return inspect.isfunction(value)


# Answers if the given value is *really* a list (and not a string).
def islist(value):
    return isinstance(value, list) and not isinstance(value, basestring)


def isset(value):
    return isinstance(value, set)


# Answers if the given value is a string.
def isstring(value):
    return isinstance(value, basestring)


# Answers if the given value is a tuple.
def istuple(value):
    return isinstance(value, tuple)
