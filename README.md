### json-response

JSON wants returned values to be in key-value format. This package removes a couple lines of code I got tired of copypasta-ing between projects.

### Dependencies
- [Echo](https://labstack.com/echo)

### Usage
```
import "github.com/jessemillar/jsonresp"
```
```
func SampleFunction(context echo.Context) error {
	err := functionThatReturnsAnError()
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}
}
```
