### json-response

JSON wants returned values to be in key-value format. This package removes a couple lines of code I got tired of copypasta-ing between projects.

### Dependencies
- [Echo](https://labstack.com/echo)

### Usage
```
func SampleFunction(c echo.Context) error {
	err := functionThatReturnsAnError()
	if err != nil {
		return jsonresp.Create(c, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}
}
```
