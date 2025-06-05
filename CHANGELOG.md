# Changelog

## 0.1.0-alpha.24 (2025-06-05)

Full Changelog: [v0.1.0-alpha.23...v0.1.0-alpha.24](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.23...v0.1.0-alpha.24)

### Bug Fixes

* **api:** use any type for `WriteParams.Delete` too ([f4dc15a](https://github.com/turbopuffer/turbopuffer-go/commit/f4dc15a1abb9fb55b782c6564a3dcf8e9dafedb1))

## 0.1.0-alpha.23 (2025-06-04)

Full Changelog: [v0.1.0-alpha.22...v0.1.0-alpha.23](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.22...v0.1.0-alpha.23)

### Bug Fixes

* **api:** don't use floats for numeric attr values ([754d14b](https://github.com/turbopuffer/turbopuffer-go/commit/754d14b0d652554c69cf554bafe2968b0a10a4da))
* **tests:** update for new JSON decoding behavior ([aa02f73](https://github.com/turbopuffer/turbopuffer-go/commit/aa02f73f06455fcd73ffed27e8af99a305678002))

## 0.1.0-alpha.22 (2025-06-04)

Full Changelog: [v0.1.0-alpha.21...v0.1.0-alpha.22](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.21...v0.1.0-alpha.22)

### Bug Fixes

* **api:** simplify type of Row and Columns ([faf595c](https://github.com/turbopuffer/turbopuffer-go/commit/faf595cdf0c075941e4f34d3c3f7027c4fbc1685))
* **tests:** update for simplified Row and Columns ([25710ef](https://github.com/turbopuffer/turbopuffer-go/commit/25710eff67de538ff89890dd4e3fa41b0d0b2eaa))
* **tests:** update for simplified Row and Columns ([c267e76](https://github.com/turbopuffer/turbopuffer-go/commit/c267e769c813b7694e5d1f3b4fb9ea2a8c577914))

## 0.1.0-alpha.21 (2025-06-04)

Full Changelog: [v0.1.0-alpha.20...v0.1.0-alpha.21](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.20...v0.1.0-alpha.21)

### Features

* **api:** add a dedicated Namespace resource ([dbd639d](https://github.com/turbopuffer/turbopuffer-go/commit/dbd639dd4575e525fa5cc641c4124effb2eeac7d))

## 0.1.0-alpha.20 (2025-06-03)

Full Changelog: [v0.1.0-alpha.19...v0.1.0-alpha.20](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.19...v0.1.0-alpha.20)

### Features

* **api:** manual updates ([e3eb2a0](https://github.com/turbopuffer/turbopuffer-go/commit/e3eb2a03e8db74e85ea0fa2934c91121e88ec721))
* **api:** use float32 instead of float64 when possible ([1139c83](https://github.com/turbopuffer/turbopuffer-go/commit/1139c83e4b7f7cc00154a8a30a988bc6bafe29c5))
* **tests:** add a test to protect against regressions of ToParam ([ebe656f](https://github.com/turbopuffer/turbopuffer-go/commit/ebe656f0a4a50b1c77c891f0aae191d6f45307e0))


### Bug Fixes

* **api:** use float32 rather than float64 in vecs ([4e45fad](https://github.com/turbopuffer/turbopuffer-go/commit/4e45fad61ec2dd2924b68a6f068f0de87aeecd72))
* **client:** cast to raw message when converting to params ([96cfe2e](https://github.com/turbopuffer/turbopuffer-go/commit/96cfe2ecbae487f0a95eea1cc5adfbbe23f47bc2))
* **internal:** adjust formatting of generated code ([f26c503](https://github.com/turbopuffer/turbopuffer-go/commit/f26c5038e7af47fbe7d443786dc96e9936eab175))
* **internal:** run gofmt during supplemental codegen ([13263a2](https://github.com/turbopuffer/turbopuffer-go/commit/13263a2cfbaa5487fbfc0509d4ca9ad9d716762e))
* **test:** adjust test for float32 vectors ([667b516](https://github.com/turbopuffer/turbopuffer-go/commit/667b5160d231622e7ae2887878c4eafc1e69b15a))

## 0.1.0-alpha.19 (2025-06-03)

Full Changelog: [v0.1.0-alpha.18...v0.1.0-alpha.19](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.18...v0.1.0-alpha.19)

### Features

* **client:** allow overriding unions ([dbea310](https://github.com/turbopuffer/turbopuffer-go/commit/dbea3106c3fdbb0afe9c7c40150c1b4d2638f3da))

## 0.1.0-alpha.18 (2025-06-03)

Full Changelog: [v0.1.0-alpha.17...v0.1.0-alpha.18](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.17...v0.1.0-alpha.18)

### Features

* **api:** allow enabling FTS using defaults only ([dc14ac2](https://github.com/turbopuffer/turbopuffer-go/commit/dc14ac229210418d11fbbad694174011d2f905a6))
* **api:** allow enabling FTS using defaults only ([a4b3522](https://github.com/turbopuffer/turbopuffer-go/commit/a4b3522b7543ed0b54efa9b444f13ca867e3f0ee))

## 0.1.0-alpha.17 (2025-06-03)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **api:** api update ([20a68c8](https://github.com/turbopuffer/turbopuffer-go/commit/20a68c8dbb12d4402b36aa18202d4f63b0034748))
* **utils:** add a PrettyPrint utility method ([4a5ffe5](https://github.com/turbopuffer/turbopuffer-go/commit/4a5ffe5d4ed78a7971eda73f9afd247e8a3396fd))


### Bug Fixes

* **examples:** update for simpler schema ([9acf5bd](https://github.com/turbopuffer/turbopuffer-go/commit/9acf5bdedd473fd67525cd2123db6d03854d7d49))

## 0.1.0-alpha.16 (2025-06-03)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Bug Fixes

* **api:** fix type for AggregateBy ([691bfbd](https://github.com/turbopuffer/turbopuffer-go/commit/691bfbd599ab5165eb92703e0e3f58fd56e88242))

## 0.1.0-alpha.15 (2025-06-03)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Features

* **api:** api update ([0a1c12b](https://github.com/turbopuffer/turbopuffer-go/commit/0a1c12b8a2495039802aac1593081115839ccff7))


### Chores

* make go mod tidy continue on error ([e5ebb69](https://github.com/turbopuffer/turbopuffer-go/commit/e5ebb69bf4b5542fc3d5b1d658f719ed28f8da65))

## 0.1.0-alpha.14 (2025-05-30)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Chores

* **internal:** codegen related update ([c26d21a](https://github.com/turbopuffer/turbopuffer-go/commit/c26d21af7ed99655dfc0b384b7ee021ac0623759))
* **internal:** version bump ([e9eee78](https://github.com/turbopuffer/turbopuffer-go/commit/e9eee7878716dadcb6c84975202be57a43d60eef))

## 0.1.0-alpha.13 (2025-05-30)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Bug Fixes

* correct unmarshalling of root body params ([d9a9efe](https://github.com/turbopuffer/turbopuffer-go/commit/d9a9efe49289498e807cf810ce5d5e5baf5629bd))

## 0.1.0-alpha.12 (2025-05-30)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** api update ([7f46cea](https://github.com/turbopuffer/turbopuffer-go/commit/7f46cea6eac8c67328b5240cd71b6c231d7df03e))

## 0.1.0-alpha.11 (2025-05-30)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Bug Fixes

* fix error ([44d1215](https://github.com/turbopuffer/turbopuffer-go/commit/44d1215d7904d35d0d5349826edc93b862d57ea0))


### Chores

* **internal:** codegen related update ([95f9dd8](https://github.com/turbopuffer/turbopuffer-go/commit/95f9dd8bd652fc97382f9d54d4e9ece6765083d5))
* **stainless:** unstick codegen ([8274970](https://github.com/turbopuffer/turbopuffer-go/commit/82749708b040548e3f06d255b1f4c96c232fc8c8))

## 0.1.0-alpha.10 (2025-05-29)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** api update ([f732735](https://github.com/turbopuffer/turbopuffer-go/commit/f732735eb07d96c4e8415b5f790d018e254be6dc))

## 0.1.0-alpha.9 (2025-05-29)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** api update ([bd82c5b](https://github.com/turbopuffer/turbopuffer-go/commit/bd82c5b06db3bf8b4f4105083c408cf162a27365))

## 0.1.0-alpha.8 (2025-05-29)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** api update ([0680880](https://github.com/turbopuffer/turbopuffer-go/commit/068088083fc668effb4b0e1c13f40d39c7064617))
* **api:** api update ([1f7b0ee](https://github.com/turbopuffer/turbopuffer-go/commit/1f7b0ee17d36ee63a4f8a6ef4306aa24418da181))

## 0.1.0-alpha.7 (2025-05-29)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** api update ([75f7317](https://github.com/turbopuffer/turbopuffer-go/commit/75f73170701793c5c34a2e0c84da2db4120b1418))
* **api:** api update ([f697124](https://github.com/turbopuffer/turbopuffer-go/commit/f69712441c0ea357fc00943e6465f7cc82f3e12c))
* **api:** api update ([7321103](https://github.com/turbopuffer/turbopuffer-go/commit/732110375feeb50c45ae563f47c49579f0b0b631))
* **api:** api update ([5db66fe](https://github.com/turbopuffer/turbopuffer-go/commit/5db66fe4f7d993df87910487ff3ffb6e97e48da6))
* **api:** api update ([1729afc](https://github.com/turbopuffer/turbopuffer-go/commit/1729afcc1b8d4b9cf7254fbe8d967eb4b41d2e35))
* **api:** api update ([814d9a0](https://github.com/turbopuffer/turbopuffer-go/commit/814d9a0d76432df2b5823652e055071e370b5247))
* **api:** api update ([bffb7fc](https://github.com/turbopuffer/turbopuffer-go/commit/bffb7fc3963a458cbf9306ac96fd92b153e4d35f))
* **api:** api update ([b42f578](https://github.com/turbopuffer/turbopuffer-go/commit/b42f578ccc119fcb7f8d638d954567a6242c4c47))
* **api:** api update ([f1d9aba](https://github.com/turbopuffer/turbopuffer-go/commit/f1d9abad327a5e565db4c442997c412eaf6cc67b))
* **api:** api update ([4d865b6](https://github.com/turbopuffer/turbopuffer-go/commit/4d865b62b7af57756fb9d3df422a0cea2cce0ed2))


### Chores

* fix union naming in readme ([f3ab2f2](https://github.com/turbopuffer/turbopuffer-go/commit/f3ab2f23671357575673794ed8f28597dace0fe2))

## 0.1.0-alpha.6 (2025-05-28)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** api update ([f817adb](https://github.com/turbopuffer/turbopuffer-go/commit/f817adba52f5e88984dc9e57d742eb3d3d6f9b5b))
* **api:** api update ([eab169d](https://github.com/turbopuffer/turbopuffer-go/commit/eab169d1d2d5b8d916c430b8f3a36f222f2494cb))

## 0.1.0-alpha.5 (2025-05-28)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** api update ([6705a45](https://github.com/turbopuffer/turbopuffer-go/commit/6705a4536b871f78bfcb35c2403fb00b154cc070))

## 0.1.0-alpha.4 (2025-05-28)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** api update ([e290c95](https://github.com/turbopuffer/turbopuffer-go/commit/e290c95fff148224f6372b0d20ce0b502dc28167))

## 0.1.0-alpha.3 (2025-05-28)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([63183a1](https://github.com/turbopuffer/turbopuffer-go/commit/63183a1497a670fc17c4593a0c533b36ad1b386f))


### Bug Fixes

* **readme:** correct capitalization ([6449434](https://github.com/turbopuffer/turbopuffer-go/commit/6449434a8883a3fff60a24f58b44e51d818949b7))
* **readme:** fix puffy image ([62b7f8d](https://github.com/turbopuffer/turbopuffer-go/commit/62b7f8d3d3e41389824a0299f6b0deb574853a1b))

## 0.1.0-alpha.2 (2025-05-28)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** api update ([e43993f](https://github.com/turbopuffer/turbopuffer-go/commit/e43993f9d730f170d9434fd72b86c76fdbcea3bd))
* **api:** api update ([cc1114e](https://github.com/turbopuffer/turbopuffer-go/commit/cc1114eac9c052e16be62fd88689773df244069a))
* **api:** api update ([4111c5e](https://github.com/turbopuffer/turbopuffer-go/commit/4111c5ea17dba916a2bfd37d223e8d859461ed40))
* **README:** add alpha warning ([187bfa2](https://github.com/turbopuffer/turbopuffer-go/commit/187bfa209a0a16c7193bb831b1157e22180ff56b))


### Bug Fixes

* **README:** update code examples ([235b34d](https://github.com/turbopuffer/turbopuffer-go/commit/235b34dc822fce37c25b85434c32a270be43097b))

## 0.1.0-alpha.1 (2025-05-22)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/turbopuffer/turbopuffer-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** api update ([2d9643f](https://github.com/turbopuffer/turbopuffer-go/commit/2d9643fe72c0392c4d856cc34c9975188dcbae5b))
* **api:** api update ([ab8df0c](https://github.com/turbopuffer/turbopuffer-go/commit/ab8df0c473ba246eaf68b80cdc20121206c69cd8))
* **api:** api update ([14899d0](https://github.com/turbopuffer/turbopuffer-go/commit/14899d07ef23bd927e9fca9076fa654f678b0789))
* **api:** api update ([3209cd1](https://github.com/turbopuffer/turbopuffer-go/commit/3209cd1e134b53582f86a26890e8604b65e44414))
* **api:** api update ([4a910ac](https://github.com/turbopuffer/turbopuffer-go/commit/4a910ac668b9717f0f96641940175af155f340cc))
* **api:** api update ([fd23ffb](https://github.com/turbopuffer/turbopuffer-go/commit/fd23ffbb97b060133d9464f01c47fbaec91e3410))
* **api:** api update ([63842ee](https://github.com/turbopuffer/turbopuffer-go/commit/63842ee011a8e22650e7c16ae3a143fb9a114d80))
* **api:** api update ([999a5b9](https://github.com/turbopuffer/turbopuffer-go/commit/999a5b9d5d701815d7800252d6f7f7908255886a))
* **api:** api update ([a547842](https://github.com/turbopuffer/turbopuffer-go/commit/a547842d735c9f6ddeb71e6baed3fd3637b95297))
* **api:** api update ([ab3eecc](https://github.com/turbopuffer/turbopuffer-go/commit/ab3eecc2da94caf850b32a80267004008491f528))
* **api:** api update ([a39a420](https://github.com/turbopuffer/turbopuffer-go/commit/a39a420bed82ef39e14378b5b43b6f6e82e147b3))
* **api:** api update ([b9be695](https://github.com/turbopuffer/turbopuffer-go/commit/b9be6956c4fba1a0a2f3b15395437bf32ee37f17))
* **api:** api update ([b86572f](https://github.com/turbopuffer/turbopuffer-go/commit/b86572fbbc3340b62177d80e766fc8de606af5a6))
* **api:** api update ([4032aba](https://github.com/turbopuffer/turbopuffer-go/commit/4032ababb4133258cda0ad46b73dfb038c6d9d38))
* **api:** api update ([8ca198b](https://github.com/turbopuffer/turbopuffer-go/commit/8ca198b8855a9671f41079637017b8e4774d3179))
* **api:** api update ([58e3ac4](https://github.com/turbopuffer/turbopuffer-go/commit/58e3ac401f7c540cf3cf4c1d92290874dec09ac4))
* **api:** api update ([fb2f6e7](https://github.com/turbopuffer/turbopuffer-go/commit/fb2f6e7384d5006b7f4981418a907f2a45e0f41e))
* **api:** api update ([80238f0](https://github.com/turbopuffer/turbopuffer-go/commit/80238f0a151621a001dc4ef8fc51a9a6dc99e3b5))
* **api:** api update ([5e1cfd3](https://github.com/turbopuffer/turbopuffer-go/commit/5e1cfd3e4fe63278525fa2579e02ff09e71cd194))
* **api:** api update ([53340e6](https://github.com/turbopuffer/turbopuffer-go/commit/53340e6ccdc2c78501ea13987ab030f6caea4368))
* **api:** api update ([999bc21](https://github.com/turbopuffer/turbopuffer-go/commit/999bc21e7107f8a14adc7c4a9e5c693cbdf154d2))
* **api:** api update ([027186c](https://github.com/turbopuffer/turbopuffer-go/commit/027186cf59e146e74f9dea22752a4cf3aa7fcdfb))
* **api:** api update ([dcf6f71](https://github.com/turbopuffer/turbopuffer-go/commit/dcf6f71c910dea739b5f003bbe8b1d34e3ec0a87))
* **api:** api update ([e5d3186](https://github.com/turbopuffer/turbopuffer-go/commit/e5d31862ec6fcc7bd014a99eafa57e6cd56c809e))
* **api:** api update ([d598e50](https://github.com/turbopuffer/turbopuffer-go/commit/d598e50a1f92adc8757beefa3a86ff696484fd97))
* **api:** manual updates ([1281f28](https://github.com/turbopuffer/turbopuffer-go/commit/1281f281a80d8ec5a14e9c2dafcbc2dc4454463f))
* **client:** add support for endpoint-specific base URLs in python ([3faacbf](https://github.com/turbopuffer/turbopuffer-go/commit/3faacbfcd05819208a4c4f96966ac634bbd088f1))


### Bug Fixes

* **client:** correctly set stream key for multipart ([8d9be31](https://github.com/turbopuffer/turbopuffer-go/commit/8d9be31fc93722c36aadc2add73557409f7b2c73))
* **client:** don't panic on marshal with extra null field ([814038d](https://github.com/turbopuffer/turbopuffer-go/commit/814038dea3f59b3d9da21b0772b5dd2de53d8923))


### Chores

* configure new SDK language ([d705ddf](https://github.com/turbopuffer/turbopuffer-go/commit/d705ddf1b6222b4c4c3133347bd6c7344d82a5fe))
* configure new SDK language ([7a8f813](https://github.com/turbopuffer/turbopuffer-go/commit/7a8f8139ab65e9ca4b9f1ac402ad126ea73c4546))
* configure new SDK language ([bd2db13](https://github.com/turbopuffer/turbopuffer-go/commit/bd2db1340c530ee0934b3547b5988e0ab2863215))
* **docs:** grammar improvements ([d9d6aed](https://github.com/turbopuffer/turbopuffer-go/commit/d9d6aedd7ef299212d4c9c08e314147849b40e21))
* improve devcontainer setup ([09186b3](https://github.com/turbopuffer/turbopuffer-go/commit/09186b3d66e85721e5c4df22cc5f5fc5e759668f))
* **internal:** codegen related update ([ee8e653](https://github.com/turbopuffer/turbopuffer-go/commit/ee8e6537e90549683e035a307d53d0e2f4b39cd5))
* update SDK settings ([88b1029](https://github.com/turbopuffer/turbopuffer-go/commit/88b1029351fcd23a3be012439f9286c260b6143c))
* update SDK settings ([2d5e57c](https://github.com/turbopuffer/turbopuffer-go/commit/2d5e57c6a8f76e2b67b25d92846ad09ad9bbfc9d))
