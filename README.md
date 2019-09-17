# common
load toml config and set log


# use

```

go get github.com/ALEX-yinhao/common

```

```

    //load toml config
    common.LoadConfig("your toml config path")

    //get int config
    common.GetMustIntValue(key, value)

    //get string config
    common.GetMustStringValue(key, value)

    //init log
    common.InitLog(path string, prefix string, suffix string, size int)

    //use log
    common.Log.Info("test")


```