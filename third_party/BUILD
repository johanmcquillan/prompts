package(default_visibility = ["PUBLIC"])

go_module(
    name = "testify",
    install = [
        "assert",
        "require",
    ],
    module = "github.com/stretchr/testify",
    version = "v1.6.1",
    deps = [
        ":difflib",
        ":spew",
        ":yaml",
    ],
)

go_module(
    name = "spew",
    install = ["spew"],
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_module(
    name = "yaml",
    licences = ["Apache 2.0"],
    module = "gopkg.in/yaml.v3",
    version = "496545a6307b2a7d7a710fd516e5e16e8ab62dbc",
)

go_module(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_module(
    name = "flags",
    module = "github.com/jessevdk/go-flags",
    version = "8766aaeab5ca44d0eceae49614346e28ab11c439",
    deps = [
        "//third_party/x:sys",
    ],
)
