# LoRa-Golang

#### 介绍
用Golang语言实现LoRa数据传输，并对收集信息进行处理

#### 软件架构
```
LoRa-Golang
├─ LICENSE
├─ main.go                     //最终运行程序
├─ module                      //子功能模块
│  ├─ libloragw                //网关所需要的库
│  │  ├─ 99-libftdi.rules      /
│  │  ├─ inc                   //存放头文件
│  │  │  ├─ loragw_aux.h
│  │  │  ├─ loragw_fpga.h
│  │  │  ├─ loragw_gps.h
│  │  │  ├─ loragw_hal.h
│  │  │  ├─ loragw_lbt.h
│  │  │  ├─ loragw_radio.h
│  │  │  ├─ loragw_reg.h
│  │  │  ├─ loragw_spi.h
│  │  │  ├─ loragw_sx125x.h
│  │  │  ├─ loragw_sx1272_fsk.h
│  │  │  ├─ loragw_sx1272_lora.h
│  │  │  ├─ loragw_sx1276_fsk.h
│  │  │  └─ loragw_sx1276_lora.h
│  │  ├─ library.cfg
│  │  ├─ Makefile
│  │  ├─ src
│  │  │  ├─ agc_fw.var
│  │  │  ├─ arb_fw.var
│  │  │  ├─ cal_fw.var
│  │  │  ├─ loragw_aux.c
│  │  │  ├─ loragw_fpga.c
│  │  │  ├─ loragw_gps.c
│  │  │  ├─ loragw_hal.c
│  │  │  ├─ loragw_lbt.c
│  │  │  ├─ loragw_radio.c
│  │  │  ├─ loragw_reg.c
│  │  │  └─ loragw_spi.native.c
│  │  └─ tst
│  │     ├─ test_loragw_cal.c
│  │     ├─ test_loragw_gps.c
│  │     ├─ test_loragw_hal.c
│  │     ├─ test_loragw_reg.c
│  │     └─ test_loragw_spi.c
│  └─ packet_loggerFromNode
│     ├─ global_conf.json
│     ├─ inc
│     │  └─ parson.h
│     ├─ local_conf.json
│     ├─ Makefile
│     ├─ readme.md
│     └─ src
│        └─ parson.c
├─ README.en.md
└─ README.md

```