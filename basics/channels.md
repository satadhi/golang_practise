
### 🔒 Default Behavior of Channels:

- **Receiver blocks** if there's no data in the channel.
- **Sender blocks** if the channel is full (in the case of buffered channels) or there's no receiver ready (in unbuffered channels).

### ⚡ `select` to the Rescue:

You can use the `select` statement to **avoid blocking**. It allows you to listen to multiple channels and **proceed with whatever is ready**.

### ✅ Non-blocking Example with `select`:

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No data, moving on...")
}
```

In this case, if `ch` has no data, it hits the `default` case instead of blocking. So yeah, `select` gives you more control and makes things non-blocking when needed.

Let me know if you want the same logic with `send`, timeouts, or multiple channels too!
