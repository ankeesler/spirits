# auto tests

You can specify a pair of files for test input and output as follows:

---

#### some-test-input.json:
```
[
  {"name": "a", "health": 2", "power": 1},
  {"name": "b", "health": 1", "power": 1}
]
```

#### some-test-output.txt
```
> summary
  a: 2
  b: 1
> summary
  a: 1
  b: 1
> summary
  a: 1
  b: 0
```

---

That will generate a test that does this:

```
client                                             server
  |                                                  |
  |  ---> battle-start w/ some-test-input.json --->  |
  |                                                  |
  |  <--- battle-stop w/ some-test-output.txt <---   |
  |                                                  |
```

The files must have the same name before the file extension (i.e., `a.json` and
`aa.json` will not match).
