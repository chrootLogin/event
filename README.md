# Event Handler for GoLang
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fchrootlogin%2Fevent.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fchrootlogin%2Fevent?ref=badge_shield)


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


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fchrootlogin%2Fevent.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fchrootlogin%2Fevent?ref=badge_large)