Here we have two examples.

First run `docker-compose up`

and then for python:

```
python -m venv venv
source venv/bin/activate
pip install gremlinpython
python main.py
```

for go:

1. github.com/qasaur/gremgo can not be even installed using "go get"
2. https://github.com/go-gremlin/gremlin is used in gremlin.go and gives https://github.com/go-gremlin/gremlin/issues/18
3. github.com/northwesternmutual/grammes WORKS!! See implementation in main.go

