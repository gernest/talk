So, I decided to try out gopherjs and the vecty framework. The results are implressinve, the following are notes I have taken that misght be useful.

### Channels
Go channels and goroutine works fine. But, don't send data over un initalizedchannel. The compiler will build and your app will run without complining but it won't work, so `make(chan )` before sending anything, in go the app will cra
