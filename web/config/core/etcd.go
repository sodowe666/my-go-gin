package core

//func initEtcd() {
//	cfg := client.Config{
//		Endpoints:               []string{"http://127.0.0.1:2379"},
//		Transport:               client.DefaultTransport,
//		HeaderTimeoutPerRequest: time.Second,
//	}
//	c, err := client.New(cfg)
//	if err != nil {
//		panic(err.Error())
//	}
//	kapi := client.NewKeysAPI(c)
//	watcher := kapi.Watcher("/kkk", &client.WatcherOptions{
//		AfterIndex: 0,
//	})
//}
