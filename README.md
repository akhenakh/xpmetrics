# XPMetrics

Library to read X-Plane metrics sent over UDP.

Instantiate a `NewListener` and run `Start()` in its own goroutine
```
l, err := xpmetrics.NewListener("127.0.0.1:49012")
go l.Start()
```

X-Plane send 6 float32 per metrics type (XPMsg)
```
v, ok := l.Data.Query(xpmetrics.LatLonAltMsg)
lat := v[0]
```

Or use some helper functions
```
hdg, ok := l.Data.Compass()
```

To activate UDP Data Output in X-Plane go to the settings | Data Output tabs | Check a line as Network via UDP.  

On the same page on the right check send network data output and set the IP Address and Port of the listener.

## Moving Map
A simple moving map demo is located in `cmd/xpmap`.
