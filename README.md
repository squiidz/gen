# Gen

Gen is a Tiny static site generator, base on a json file.

#### Create file which content :

``` json
{
	"name": "test",
	"version": "0.0.1",
	"routes": [{
		"url": "/index",
		"method": "GET",
		"template": "index.html"
	},
	{
		"url": "/home",
		"method": "GET",
		"template": "home.html"
	},
	{
		"url": "/admin",
		"method": "GET",
		"template": "index.html"
	}]
}

```
#### And create a tiny server from it:

``` go

package main

import (
	"github.com/squiidz/gen"
)

func main() {
	g := gen.NewGen("application/app.json")
	g.Run(":8080")
}

```

## TODO
This project is only for fun.

## License

MIT