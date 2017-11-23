# Data Management

## Store

The store interface gives access to useful rapid key|value stores such as redis. The store uses a consistent interface defined as such:

```go

type Store interface {
	ChangeExpiration(time.Duration) Store
	Set(string, string) error
	SetEx(string, time.Duration, string) error
	Get(string) (StoreValue, error)
	Del(string) error
	Keys(string) (StoreValue, error)
}

```

Usage is simple. 
```go

func SomeHandler(app *butter.App) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.Store("my-key", "my-value")
        
        val, _:= app.Store.Get("my-key")
        b := val.Value() 

		w.WriteHeader(200)
	})
}
```

## Database
