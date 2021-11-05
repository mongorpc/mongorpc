package mongorpc

import (
	"github.com/fsnotify/fsnotify"
)

func WatchFile(filePath string, callback func(err error, event *fsnotify.Event)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		callback(err, nil)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					callback(nil, &event)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				callback(err, nil)
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		callback(err, nil)
	}
	<-done
}
