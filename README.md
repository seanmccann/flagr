我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

<p align="center">
    <img src="./browser/flagr-ui/src/assets/logo.png" width="150">
</p>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/checkr/flagr" target="_blank">
        <img src="https://goreportcard.com/badge/github.com/checkr/flagr">
    </a>
    <a href="https://circleci.com/gh/checkr/flagr" target="_blank">
        <img src="https://circleci.com/gh/checkr/flagr.svg?style=shield">
    </a>
    <a href="https://gitter.im/checkr-flagr/Lobby" target="_blank">
        <img src="https://img.shields.io/gitter/room/nwjs/nw.js.svg">
    </a>
    <a href="https://raw.githubusercontent.com/checkr/flagr/master/LICENSE" target="_blank">
        <img src="http://img.shields.io/badge/license-Apache%20v2-green.svg">
    </a>
</p>

## Introduction

Flagr is an open source Go service that delivers the right experience to the right entity and monitors the impact. It provides feature flags, experimentation (A/B testing), and dynamic configuration. It has clear swagger REST APIs for flags management and flag evaluation.

## Documentation
- [Flagr Documentation - https://checkr.github.io/flagr](https://checkr.github.io/flagr/)

## Quick demo

```sh
# Get the flagr-mini image for the demo
docker run -it -p 18000:18000 checkr/flagr

# Open the flagr homepage
open localhost:18000
```

## Flagr Evaluation Performance

Tested with `vegeta`. For more details, see [benchmarks](./benchmark).

```
Requests      [total, rate]            56521, 2000.04
Duration      [total, attack, wait]    28.2603654s, 28.259999871s, 365.529µs
Latencies     [mean, 50, 95, 99, max]  371.632µs, 327.991µs, 614.918µs, 1.385568ms, 12.50012ms
Bytes In      [total, mean]            23250552, 411.36
Bytes Out     [total, mean]            8308587, 147.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:56521
Error Set:
```

## Flagr UI

<p align="center">
    <img src="./docs/images/demo_readme.png" width="900">
</p>

