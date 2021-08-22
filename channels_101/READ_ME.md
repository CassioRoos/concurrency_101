# THIS ARTICLE CAN BE READ IN FULL [HERE](https://go101.org/article/channel-use-cases.html)

#### var bidirectionalChan _chan string_  
> can read from, write to and close()
#### var receiveOnlyChan _<-chan string_  
>can read from, but cannot write to or close()
#### var sendOnlyChan _chan<- string_  
>cannot read from, but can write to and close()


##Ways to block the current goroutine ‎forever by using the channel mechanism
Without importing any package, we can use the following ways to make the current goroutine ‎enter (and stay in) blocking state forever:

1. send a value to a channel which no ones will receive values from

``` go
make(chan struct{}) <- struct{}{}
```
 or
``` go
make(chan<- struct{}) <- struct{}{}
```
2. receive a value from a never-closed channel which no values have been and will be 
sent to
``` go
<-make(chan struct{})
```
 or
``` go
<-make(<-chan struct{})
```
 or
``` go
for range make(<-chan struct{}) {}
```
3. receive a value from (or send a value to) a nil channel
``` go
chan struct{}(nil) <- struct{}{}
```
or
``` go
<-chan struct{}(nil)
```
 or
``` go
for range chan struct{}(nil) {}
```
4. use a bare select block
``` go
select{}
```