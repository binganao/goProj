class a():
    #b=1
    @property
    def b(self):
        return 2

d=a()
print(d.b)