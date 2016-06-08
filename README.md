Gohint is an alternative linter for Go source code. It is based on
[golint](https://github.com/golang/lint/golint) and inspired by [jshint](http://www.jshint.com/).
That is, basically gohint is the same golint plus configuration (plus some extra checks,
plus report generators, plus... :) )

To install, run
```
  go get github.com/hellosuper/gohint
```

Any other documentation you can check at [golint's README](https://github.com/golang/lint/blob/master/README).
At the moment there is not so much difference, but I'm not sure what will happen in future :)

Why you may need it? As `golint` says:

```
The suggestions made by golint are exactly that: suggestions.
Golint is not perfect, and has both false positives and false negatives.
```

But sometimes you want those suggestions to work, so you start to use the tool
in your continuous integration. And after that you begin to realize that many of
those suggestions do not fit your project and just add noise to health report.

That's why you can try `gohint` - you can configure it, so it will scream only
in case of problems that really hurt **your** project. Don't try to filter the noise,
define your rules and do not produce such noise! :)

Also it can produce reports in various formats, so integration to your CI cycle becomes even easier.
At the moment it supports only 2 formats: plain text and [Checkstyle XML](http://checkstyle.sourceforge.net/).

# Running gohint

`gohint` supports 2 options that can be passed:

| opt        | description                                                                                                                                                                                     |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `config`   | path to JSON file with configuration. See above how to prepare config file                                                                                                                      |
| `reporter` | name of reporter to use for output. Supported ones: `plain` and `checkstyle`.  `plain` outputs report in plain text (like `golint`) and is used by default. `checkstyle` outputs Checkstyle XML |

Example:

```
gohint -config="/path/to/config.json" -reporter=plain
```


# Configuration options

JSON config can contain following options:

TBD...

| opt                | type    | description                                                                       |
|--------------------|---------|-----------------------------------------------------------------------------------|
| **package**        | *bool*  | analyze package comments, definitions, etc                                        |
| **imports**        | *bool*  | check imports for `.`                                                             |
| **names**          | *bool*  | check names of functions, variables, etc                                          |
| **exported**       | *bool*  | check exported types and vars for correct comments and other                      |
| **var-decls**      | *bool*  | analyze variable declaration for correctness                                      |
| **elses**          | *bool*  | check `if..else` statements for redundant `else`                                  |
| **make-slice**     | *bool*  | check making slices for using short syntax                                        |
| **error-return**   | *bool*  | check list of function's return values for position of `error`, it should be last |
| **ignored-return** | *bool*  | check if there any function call which returned result is ignored                 |
| **min-confidence** | *float* | minimal confidence value to reduce output, `[0..1]`                               |
