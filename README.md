# Event Handler for GoLang

This can be used to create an event system in your go application

## Example

### Register event handler

```
event.Events().On("init-router", func(input ...interface{}) {
  router, ok := input[0].(*gin.Engine)
	if !ok {
		log.Fatal("Can't convert input to gin.Engine.")
	}

	plugins.Registry().RegisterRoutes(router)
})
```

### Emit event and run event handlers

```
event.Events().Emit("init-router", router)
```

## License

This stuff is licensed under LGPL 3.0.
