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