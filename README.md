# utils [![Go](https://github.com/flow-lab/utils/actions/workflows/go.yml/badge.svg)](https://github.com/flow-lab/utils/actions/workflows/go.yml)

Go utils. This package contains some useful functions. It is used in other packages of the flow-lab org. 

## Installation

```shell
# Install the package
go get github.com/flow-lab/utils

# Import in the project
import utils "github.com/flow-lab/utils"

# Use it
envVar := utils.EnvOrDefault("ENV_VAR", "default_value")
```

License [MIT](LICENSE)
