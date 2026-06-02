# Changelog

## 2.3.0 (2026-06-02)

Full Changelog: [v2.2.0...v2.3.0](https://github.com/turbopuffer/turbopuffer-go/compare/v2.2.0...v2.3.0)

### Features

* spec: add SDK support for native embedding ([2147191](https://github.com/turbopuffer/turbopuffer-go/commit/214719174c298eec20447c32dbd9e4dfd3b9f174))


### Bug Fixes

* reject malicious poll locations ([#123](https://github.com/turbopuffer/turbopuffer-go/issues/123)) ([2ae058e](https://github.com/turbopuffer/turbopuffer-go/commit/2ae058eedf07f1cd78743996d017c9226cb0a5a1))


### Chores

* bump apigen ([#124](https://github.com/turbopuffer/turbopuffer-go/issues/124)) ([e30d0fb](https://github.com/turbopuffer/turbopuffer-go/commit/e30d0fb2ee3a44968b5fa9b3532a9ad1b1b1993b))
* fix API docs links ([#121](https://github.com/turbopuffer/turbopuffer-go/issues/121)) ([54effa8](https://github.com/turbopuffer/turbopuffer-go/commit/54effa8e9364b9e381872e157bf9b753ea4ac779))

## 2.2.0 (2026-05-29)

Full Changelog: [v2.1.0...v2.2.0](https://github.com/turbopuffer/turbopuffer-go/compare/v2.1.0...v2.2.0)

### Features

* spec: force generation of FuzzyParams stainless models ([dc75877](https://github.com/turbopuffer/turbopuffer-go/commit/dc75877336d94ac920881299f92157f4b81d5c02))
* transparent async polling ([#119](https://github.com/turbopuffer/turbopuffer-go/issues/119)) ([7dcd309](https://github.com/turbopuffer/turbopuffer-go/commit/7dcd3098ba6bea7c7080047caf8ee6c233563161))

## 2.1.0 (2026-05-17)

Full Changelog: [v2.0.0...v2.1.0](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0...v2.1.0)

### Features

* **client:** optimize json encoder for internal types ([a3a96ea](https://github.com/turbopuffer/turbopuffer-go/commit/a3a96ea12f2b8d296d08895938872ac21ceaef5d))

## 2.0.0 (2026-05-10)

Full Changelog: [v2.0.0-alpha.6...v2.0.0](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.6...v2.0.0)

### Documentation

* document v2.0 breaking changes in UPGRADING.md ([393ee30](https://github.com/turbopuffer/turbopuffer-go/commit/393ee303026a3506c58c683dc4f28f9313558c95))

## 2.0.0-alpha.6 (2026-05-10)

Full Changelog: [v2.0.0-alpha.5...v2.0.0-alpha.6](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.5...v2.0.0-alpha.6)

## 2.0.0-alpha.5 (2026-05-10)

Full Changelog: [v2.0.0-alpha.4...v2.0.0-alpha.5](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.4...v2.0.0-alpha.5)

### Features

* **go:** add QueryParam alias for NamespaceMultiQueryParamsQuery ([46b9fdf](https://github.com/turbopuffer/turbopuffer-go/commit/46b9fdfaea0863110e03f5b1055361c9f7b91f7c))
* spec: rename RankByVector/RankBySparseVector to RankByAnn/RankBySparseKnn ([633b37d](https://github.com/turbopuffer/turbopuffer-go/commit/633b37dcc858aa03bc1eb416fcdb6d0a9f4b50d5))


### Bug Fixes

* **go:** use QueryParam, GroupBy, and RankByAnn in tests and examples ([58b91b8](https://github.com/turbopuffer/turbopuffer-go/commit/58b91b8210764cdf48318d2167a46a0a4565d52d))

## 2.0.0-alpha.4 (2026-05-08)

Full Changelog: [v2.0.0-alpha.3...v2.0.0-alpha.4](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.3...v2.0.0-alpha.4)

### Features

* spec: move update_metadata to v1 ([f2ec0e9](https://github.com/turbopuffer/turbopuffer-go/commit/f2ec0e9aabbd47e3135a13440cbf266fef0a42c4))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([fde3f7a](https://github.com/turbopuffer/turbopuffer-go/commit/fde3f7a22e7b7a2dde49241f3caa2f6b61d58d30))


### Chores

* avoid embedding reflect.Type for dead code elimination ([9567f84](https://github.com/turbopuffer/turbopuffer-go/commit/9567f846900aa1d87a602a03de95188cdc9ed842))
* redact api-key headers in debug logs ([1f0353a](https://github.com/turbopuffer/turbopuffer-go/commit/1f0353a26a71264d6e67e3bf141a20d7ed147451))

## 2.0.0-alpha.3 (2026-04-29)

Full Changelog: [v2.0.0-alpha.2...v2.0.0-alpha.3](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.2...v2.0.0-alpha.3)

### Features

* Spec: Add sparse_distance_metric to the list of models. ([793260f](https://github.com/turbopuffer/turbopuffer-go/commit/793260fb2e4c2ddd3ab60bcee71a56d091255760))

## 2.0.0-alpha.2 (2026-04-29)

Full Changelog: [v2.0.0-alpha.1...v2.0.0-alpha.2](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.1...v2.0.0-alpha.2)

### Bug Fixes

* wrap branch_from/copy_from syntax sugar bodies ([2345916](https://github.com/turbopuffer/turbopuffer-go/commit/234591688f5b0068be0e9826d2637e8f80833871))

## 2.0.0-alpha.1 (2026-04-29)

Full Changelog: [v2.0.0-alpha.1...v2.0.0-alpha.1](https://github.com/turbopuffer/turbopuffer-go/compare/v2.0.0-alpha.1...v2.0.0-alpha.1)

### Features

* [TPUF-1097] tpuf: return failed IDs for conditional writes ([50ba903](https://github.com/turbopuffer/turbopuffer-go/commit/50ba903842e57b2b2558543b6c2cb19f4e9d90e6))
* add ContainsAnyToken and RankByKnn ([ab283b9](https://github.com/turbopuffer/turbopuffer-go/commit/ab283b9406b78ea44b49b6bc94e1abde7037b974))
* Add float, []float and []bool to the list of valid types in the OpenAPI spec. ([9e4d2b0](https://github.com/turbopuffer/turbopuffer-go/commit/9e4d2b05be7ae87ddaa445a4d5a3e5614e0217ad))
* Add ranking-by-attribute to the spec. ([2f3172d](https://github.com/turbopuffer/turbopuffer-go/commit/2f3172dc6eb1262654e23cddf9ff4200e73852d8))
* Add sparse vectors to the spec. ([58682b3](https://github.com/turbopuffer/turbopuffer-go/commit/58682b3a40d4663667d6570d5a1f5a678f7bcfd1))
* add support for `limit` query parameter ([ea97752](https://github.com/turbopuffer/turbopuffer-go/commit/ea977522998bc20a71d190dd73dd51f5a5c03a87))
* add update_metadata method to configure pinning ([9266200](https://github.com/turbopuffer/turbopuffer-go/commit/9266200541f0659ff953d60c7e1489645f0bff14))
* Add vector attribute schema to metadata endpoint ([9fd79eb](https://github.com/turbopuffer/turbopuffer-go/commit/9fd79ebacc251cb11b307e6404c4844b34b7dda4))
* add WithParams variant to BM25 and ContainsAllTokens ([b6ebde5](https://github.com/turbopuffer/turbopuffer-go/commit/b6ebde58ceebc6bbd21682c7f6c489f9bb75dcc9))
* add word_v3 to the spec ([1189307](https://github.com/turbopuffer/turbopuffer-go/commit/11893071516f55828f914e03c430fa6dcee3525e))
* Allow for a CMEK key to be specified in copy_from_namespace ([c59eabb](https://github.com/turbopuffer/turbopuffer-go/commit/c59eabb5b0d1d5d484785a39082c33511ff1cfb9))
* **api:** add a dedicated Namespace resource ([b87fd0e](https://github.com/turbopuffer/turbopuffer-go/commit/b87fd0ed0a6e83bc506231160c8d26b612198782))
* **api:** allow enabling FTS using defaults only ([7c739ea](https://github.com/turbopuffer/turbopuffer-go/commit/7c739ead83e9612af735a43f4f472dbea76dae02))
* **api:** allow enabling FTS using defaults only ([3b1ba4a](https://github.com/turbopuffer/turbopuffer-go/commit/3b1ba4a65f0777320363db4cff259b86d7c627b4))
* **api:** api update ([a8447e2](https://github.com/turbopuffer/turbopuffer-go/commit/a8447e24747ba4a6ae22ba82c7c5d393005facff))
* **api:** api update ([50ccebd](https://github.com/turbopuffer/turbopuffer-go/commit/50ccebde3f013306fe2374cb775b946af6ad2c82))
* **api:** api update ([7f46cea](https://github.com/turbopuffer/turbopuffer-go/commit/7f46cea6eac8c67328b5240cd71b6c231d7df03e))
* Apigen integration ([70d4e07](https://github.com/turbopuffer/turbopuffer-go/commit/70d4e072147aada6126427a2ac011aa22d8b5f2b))
* **api:** introduce dedicated Query model ([531f2df](https://github.com/turbopuffer/turbopuffer-go/commit/531f2df3e602bd785ad7eea5e6e0e6c20b9c72bc))
* **api:** manual updates ([a9c5066](https://github.com/turbopuffer/turbopuffer-go/commit/a9c5066f8c95e98306474c4a6a923716dfc756d0))
* **api:** use float32 instead of float64 when possible ([79fe419](https://github.com/turbopuffer/turbopuffer-go/commit/79fe41966b651a450c05c2c83b75089d0eb29539))
* **client:** add a convenient param.SetJSON helper ([4bc0012](https://github.com/turbopuffer/turbopuffer-go/commit/4bc00126ea1f29f8b1e271a4c7d8e880ea8f1d3b))
* **client:** add escape hatch for null slice & maps ([58ff5bf](https://github.com/turbopuffer/turbopuffer-go/commit/58ff5bf081043089624b3667ff202ba7c27a4d1f))
* **client:** support optional json html escaping ([6796be2](https://github.com/turbopuffer/turbopuffer-go/commit/6796be255e1a2b28015ae01b000ec7f08e91791e))
* ContainsAnyToken last_as_prefix docs ([bfd90aa](https://github.com/turbopuffer/turbopuffer-go/commit/bfd90aa58be1517afe7765db15c3758c3a70fa59))
* disable response compression by default ([#82](https://github.com/turbopuffer/turbopuffer-go/issues/82)) ([17d5293](https://github.com/turbopuffer/turbopuffer-go/commit/17d529389bcf46f14fc0d5d68b13a6e61e42e404))
* docs for glob/regex acceleration ([e527a1d](https://github.com/turbopuffer/turbopuffer-go/commit/e527a1d430d58349d3a527de37176ae47482e8c7))
* Enable Stainless normalizer v2 while keeping Query inlined ([a1db134](https://github.com/turbopuffer/turbopuffer-go/commit/a1db134f535862fb56d4709d79e294abf2e7d4ba))
* Enable Stainless normalizer v2 while keeping Query inlined ([dd7f4f5](https://github.com/turbopuffer/turbopuffer-go/commit/dd7f4f5582fc40975d6bdff4bf406e2fe0bda1d6))
* **encoder:** support bracket encoding form-data object members ([9e8445f](https://github.com/turbopuffer/turbopuffer-go/commit/9e8445f516da05bf9d154af26623206ee1d7b9f7))
* **go:** add default http client with timeout ([7c99444](https://github.com/turbopuffer/turbopuffer-go/commit/7c99444fd9c4e40e8a716528ab5e9c86941b722d))
* **internal:** support comma format in multipart form encoding ([6cc7300](https://github.com/turbopuffer/turbopuffer-go/commit/6cc7300897c7dee5ad9252e21862f863e7de7a1b))
* Make `type` required on `AttributeSchemaConfig` ([7123bd9](https://github.com/turbopuffer/turbopuffer-go/commit/7123bd9efbf76487f5da80ff8105acaf2b816731))
* Make word_v2 the default FTS tokenizer ([dc4bd86](https://github.com/turbopuffer/turbopuffer-go/commit/dc4bd8679a1bad9db8c422ca44203baa52df8d36))
* Metadata endpoint updates (e.g. to track indexing progress) ([1d6d588](https://github.com/turbopuffer/turbopuffer-go/commit/1d6d588f8dd6ac94029aecce3abbd3e75b67417f))
* Multiple Vector Columns Docs ([ea830fa](https://github.com/turbopuffer/turbopuffer-go/commit/ea830fabdcff7ee441864f9efb6cfd6ad849fc8d))
* openapi: Fix stainless warnings ([952ccf5](https://github.com/turbopuffer/turbopuffer-go/commit/952ccf5f2363568acaa1e6591685c653be57a8e1))
* openapi: name variants of `NamespaceMetadata.index` ([0f90deb](https://github.com/turbopuffer/turbopuffer-go/commit/0f90debb7ccfc8eadac4bdca250410ead720d460))
* Promote disable_backpressure to first-class Write property ([8371b8d](https://github.com/turbopuffer/turbopuffer-go/commit/8371b8d0a09c56a152a29b8884f06dd925fa4991))
* Remove `queries` from recall endpoint ([aab437c](https://github.com/turbopuffer/turbopuffer-go/commit/aab437c1db1e48f345d648be4a950f6f59cc8ea8))
* Require `fuzzy: true` in the schema to use the `Fuzzy` filter ([904a15f](https://github.com/turbopuffer/turbopuffer-go/commit/904a15fc7c200f5cb6b74e40a4d4d3f25ae4df87))
* required for patch_by_filter :facepalm: ([1820b2c](https://github.com/turbopuffer/turbopuffer-go/commit/1820b2cb4def88cc6c27c87bdac80b0a63bb0134))
* Response performance object ([792b3b5](https://github.com/turbopuffer/turbopuffer-go/commit/792b3b53ecd5a93f2e67c85422a52dda1d5efdda))
* retract versions before v1.12.2 ([#102](https://github.com/turbopuffer/turbopuffer-go/issues/102)) ([0cd7303](https://github.com/turbopuffer/turbopuffer-go/commit/0cd730381840b434c03eb8e5090ad56e55bbd0eb))
* sdks: add &lt;patch|delete&gt;_by_filter_allow_partial options ([7a484e2](https://github.com/turbopuffer/turbopuffer-go/commit/7a484e25dbf094ca2d8db59d6cbfb90e7b46477a))
* set up branchFrom and copyFrom sugar in the OpenAPI spec ([52262b4](https://github.com/turbopuffer/turbopuffer-go/commit/52262b4b325ba3006603b3172523c76ba220ba3a))
* site: add ascii_folding to docs and SDKs ([02198bf](https://github.com/turbopuffer/turbopuffer-go/commit/02198bf68ab2b96ebce3d492ca4f9062702834ba))
* skip generating Go integer alias for Limit ([4a034e6](https://github.com/turbopuffer/turbopuffer-go/commit/4a034e6b3c5421d867ed7d18667d12bed2f81165))
* spec: add dedicated type for AggregationGroup response ([aec6f7d](https://github.com/turbopuffer/turbopuffer-go/commit/aec6f7d684a479d6b01827ae5050317911d43cba))
* spec: add support for cross-org CFN to SDKs ([73a1be1](https://github.com/turbopuffer/turbopuffer-go/commit/73a1be1c4433a539abf4d5ca01acef31936b76f2))
* spec: clean up limit codegen ([9542d2f](https://github.com/turbopuffer/turbopuffer-go/commit/9542d2fada757be368f655f0262600a1c53c54f7))
* spec: disable CopyFromNamespace syntax sugar in Go ([debcc96](https://github.com/turbopuffer/turbopuffer-go/commit/debcc96c2f97a8804527feb06a716ae053e6e1c1))
* spec: flatten branchFrom/copyFrom sugar endpoint params ([c4a8a79](https://github.com/turbopuffer/turbopuffer-go/commit/c4a8a792239a7183c08d70fc2f1a479befe77007))
* spec: host MCP on stainless ([0274298](https://github.com/turbopuffer/turbopuffer-go/commit/02742984df0958b05503c0d65536a2e3a6ef9977))
* spec: include pinning status in /metadata response ([b09daa4](https://github.com/turbopuffer/turbopuffer-go/commit/b09daa4250b0766476cff41f8a8f04e81c99c796))
* spec: update OpenAPI spec for v2 encryption API ([836e30d](https://github.com/turbopuffer/turbopuffer-go/commit/836e30d92d2881575616f699b3c2e17c8c6769db))
* stainless: add patch_by_filter ([0354bce](https://github.com/turbopuffer/turbopuffer-go/commit/0354bce36a2e3333d31d6605c8fe9909a86328c1))
* support setting headers via env ([bbe8e4f](https://github.com/turbopuffer/turbopuffer-go/commit/bbe8e4f09612f18e5a5ca176188f61bda4e46319))
* **tests:** add a test to protect against regressions of ToParam ([88e7675](https://github.com/turbopuffer/turbopuffer-go/commit/88e7675400c52df645439f4285345557948d29fe))
* tpuf-engine: add verbose form of branch_from_namespace ([55a1b15](https://github.com/turbopuffer/turbopuffer-go/commit/55a1b15e38af0c0b34159d37e6aa5c73bb699757))
* tpuf: add branching to openapi spec ([6c9e02a](https://github.com/turbopuffer/turbopuffer-go/commit/6c9e02a5cdfb8350d2500aea64d840f62325cb25))
* tpuf: add include_ground_truth option to recall endpoint ([899e959](https://github.com/turbopuffer/turbopuffer-go/commit/899e959add69989c9be5cd268a01850ae3d13d98))
* **utils:** add a PrettyPrint utility method ([1ae3ec1](https://github.com/turbopuffer/turbopuffer-go/commit/1ae3ec11a9439807095141718dd17a58668c1e0f))


### Bug Fixes

* add FullTextSearchConfig.MaxTokenLength ([e032f4b](https://github.com/turbopuffer/turbopuffer-go/commit/e032f4b4ff24c47a5a1238de1e0d30e58be11c79))
* add missing Filter prefix to ContainsAllTokensArray ([#39](https://github.com/turbopuffer/turbopuffer-go/issues/39)) ([eb71dfb](https://github.com/turbopuffer/turbopuffer-go/commit/eb71dfb592ee8274c1d5dbbd2bf8802b96d40160))
* add support for `$ref_new` expressions ([#41](https://github.com/turbopuffer/turbopuffer-go/issues/41)) ([e700422](https://github.com/turbopuffer/turbopuffer-go/commit/e70042287ca2482c4da0105e0abfe93fec20333c))
* add support for regex filter ([803486f](https://github.com/turbopuffer/turbopuffer-go/commit/803486fe6c1a64ec6c5a365419374825db04d5e6))
* allow canceling a request while it is waiting to retry ([0f07f48](https://github.com/turbopuffer/turbopuffer-go/commit/0f07f4820355d9a3c44790391457d0d468ea932c))
* **api:** add support for [Not]Contains[Any] operators ([#43](https://github.com/turbopuffer/turbopuffer-go/issues/43)) ([9f6f21e](https://github.com/turbopuffer/turbopuffer-go/commit/9f6f21e957884813afd789925e275a34a3c51c9d))
* **api:** add support for `group_by` query parameter ([744ecc2](https://github.com/turbopuffer/turbopuffer-go/commit/744ecc297214af3308c0466804c269a2a1426f78))
* **api:** api update ([543e4ed](https://github.com/turbopuffer/turbopuffer-go/commit/543e4ed49dba64bffcda96648ec90c6edff542e3))
* **api:** api update ([8b19a44](https://github.com/turbopuffer/turbopuffer-go/commit/8b19a443fbed878093a10eb849d6953b33b59718))
* **api:** api update ([1d231e3](https://github.com/turbopuffer/turbopuffer-go/commit/1d231e339a081b971005a51a1e396fd79e64450b))
* **api:** api update ([3dabdea](https://github.com/turbopuffer/turbopuffer-go/commit/3dabdea5073d780deca395f49f3cf756a6bc74f6))
* **api:** api update ([2e3b385](https://github.com/turbopuffer/turbopuffer-go/commit/2e3b385f0531d2296c11edde76c8ce37f1206216))
* **api:** api update ([2d0167a](https://github.com/turbopuffer/turbopuffer-go/commit/2d0167a30c51a1168c05ceb87e88bbfdc0ab5b45))
* **api:** api update ([37ad78d](https://github.com/turbopuffer/turbopuffer-go/commit/37ad78d23ffda41261cfbf0f5253fcef1f025121))
* **api:** api update ([63b15ac](https://github.com/turbopuffer/turbopuffer-go/commit/63b15ac3b1fdadcb8f098fb64e62680410f8897a))
* **api:** api update ([ea85bf2](https://github.com/turbopuffer/turbopuffer-go/commit/ea85bf21dcc04bf4ac3556731c8cfa8eb7fe0d14))
* **api:** api update ([f18d938](https://github.com/turbopuffer/turbopuffer-go/commit/f18d938ff63124e80b43ec8b3797b793d7d8d5db))
* **api:** api update ([b20defc](https://github.com/turbopuffer/turbopuffer-go/commit/b20defcbe355a85a29ca2ff6b4225b4734e8efe4))
* **api:** don't use floats for numeric attr values ([f504e0a](https://github.com/turbopuffer/turbopuffer-go/commit/f504e0a93a4c708797ebba8fad22099597b9d9f7))
* **api:** fix type for AggregateBy ([5a3140b](https://github.com/turbopuffer/turbopuffer-go/commit/5a3140bfd78ef15a19bbbd456bd2a322507a361c))
* **api:** fix type in improved JSON number type ([5912cc6](https://github.com/turbopuffer/turbopuffer-go/commit/5912cc63e7da3223698310e4b78d5789cd81923d))
* **api:** improved JSON number type ([462e64a](https://github.com/turbopuffer/turbopuffer-go/commit/462e64ac4bf02de69e9064fff89c7369fe01cf0c))
* **api:** use any type for `WriteParams.Delete` too ([78bfa9d](https://github.com/turbopuffer/turbopuffer-go/commit/78bfa9d53c5dc5d995cfc6b6372d037c17b3df02))
* **api:** use float32 rather than float64 in vecs ([dbc472e](https://github.com/turbopuffer/turbopuffer-go/commit/dbc472ea8a12a27af07d92ac65034a272c16d879))
* bugfix for setting JSON keys with special characters ([6ad99b6](https://github.com/turbopuffer/turbopuffer-go/commit/6ad99b6fc385ea6c76b68eadb69200db9283df65))
* **client:** cast to raw message when converting to params ([60a0bdc](https://github.com/turbopuffer/turbopuffer-go/commit/60a0bdc4c9fb239ec12813ed4c6a3731eb967eb1))
* **client:** process custom base url ahead of time ([10634d2](https://github.com/turbopuffer/turbopuffer-go/commit/10634d2e3860f2a7f9f6b7c1594b2582be43f716))
* close body before retrying ([e6d57eb](https://github.com/turbopuffer/turbopuffer-go/commit/e6d57ebee96f4ca4e9a1e6631b6673c83e13fe1c))
* correct name of stainless bot ([#42](https://github.com/turbopuffer/turbopuffer-go/issues/42)) ([0a5c27c](https://github.com/turbopuffer/turbopuffer-go/commit/0a5c27c2c397101f081e93ea727b6b53f94fc962))
* correct some bugs in example code ([#56](https://github.com/turbopuffer/turbopuffer-go/issues/56)) ([84816f6](https://github.com/turbopuffer/turbopuffer-go/commit/84816f679f80b1419eb0d55ae1d72d4cabd5bb68))
* correct unmarshalling of root body params ([d9a9efe](https://github.com/turbopuffer/turbopuffer-go/commit/d9a9efe49289498e807cf810ce5d5e5baf5629bd))
* data race in Client.Namespace caused by slice aliasing ([#88](https://github.com/turbopuffer/turbopuffer-go/issues/88)) ([85a3919](https://github.com/turbopuffer/turbopuffer-go/commit/85a3919ff8f04085b9cff7fb8b5dd59ef83480f1))
* **docs:** add missing pointer prefix to api.md return types ([35991a8](https://github.com/turbopuffer/turbopuffer-go/commit/35991a8bf93a5e984b4c840e3c39485ab5252e75))
* **docs:** fix mcp installation instructions for remote servers ([0cdaa4f](https://github.com/turbopuffer/turbopuffer-go/commit/0cdaa4fbd818725c398d7f6bd3fc1b36a21e84f5))
* don't require region unless baseUrl contains {region} placeholder ([#47](https://github.com/turbopuffer/turbopuffer-go/issues/47)) ([ccecb57](https://github.com/turbopuffer/turbopuffer-go/commit/ccecb57900973a64505e4df68f95749f403c616d))
* don't try to deserialize as json when ResponseBodyInto is []byte ([f1dcf55](https://github.com/turbopuffer/turbopuffer-go/commit/f1dcf55d4ed61907ff159ca479b70a1d3b14388d))
* **encoder:** correctly serialize NullStruct ([874c674](https://github.com/turbopuffer/turbopuffer-go/commit/874c674ec9e667a8f94178f58404e70bd50905ee))
* **examples:** update examples for new names ([0ab065d](https://github.com/turbopuffer/turbopuffer-go/commit/0ab065d80ef5cabfceaa2eac01ae039e1845d2b7))
* **examples:** update for latest API changes ([2bcf32b](https://github.com/turbopuffer/turbopuffer-go/commit/2bcf32be6b9b98f8fc8134b42a7149dec639d350))
* **examples:** update for simpler schema ([d580e60](https://github.com/turbopuffer/turbopuffer-go/commit/d580e601f81ebc89c61eab1c1149e1ab38b8d4ac))
* fix error ([44d1215](https://github.com/turbopuffer/turbopuffer-go/commit/44d1215d7904d35d0d5349826edc93b862d57ea0))
* fix for namespace collisions with client and resource test methods ([3074414](https://github.com/turbopuffer/turbopuffer-go/commit/3074414e0e4fa537f937142bc4cc96d31b1cf5f6))
* fix issue with unmarshaling in some cases ([c9176e8](https://github.com/turbopuffer/turbopuffer-go/commit/c9176e8bcbf6928ead0f32b9a070ed088b55ef80))
* improve tracking of client sharing when client is copied by value ([#96](https://github.com/turbopuffer/turbopuffer-go/issues/96)) ([3b93a54](https://github.com/turbopuffer/turbopuffer-go/commit/3b93a5447d531d5b5234a5bf7269b6731e0cb31a))
* **internal:** adjust formatting of generated code ([09a8e90](https://github.com/turbopuffer/turbopuffer-go/commit/09a8e90b314083da734151a334b4fd73950e54cf))
* **internal:** run gofmt during supplemental codegen ([d9f7a68](https://github.com/turbopuffer/turbopuffer-go/commit/d9f7a683cd80fec5883c220ef8f437ae80e82d5e))
* **internal:** skip tests that depend on mock server ([cf278a8](https://github.com/turbopuffer/turbopuffer-go/commit/cf278a82f34041e174ab4071017804acc89c476f))
* **internal:** unmarshal correctly when there are multiple discriminators ([90adbcf](https://github.com/turbopuffer/turbopuffer-go/commit/90adbcfbfe6d804b39ed2ce18feda3339eabb3b0))
* lint ([f621c3e](https://github.com/turbopuffer/turbopuffer-go/commit/f621c3e2f218880a7db36dd9ca02d592eb6bc22c))
* **mcp:** correct code tool API endpoint ([f4f57c3](https://github.com/turbopuffer/turbopuffer-go/commit/f4f57c3b45f8056e55ddd7cf12dd85e43a0b94aa))
* more precise types for filters that take arrays ([9efd4e7](https://github.com/turbopuffer/turbopuffer-go/commit/9efd4e732bb5fb7723c5fcf86722f45a3e135f75))
* **pagination:** check if page data is empty in GetNextPage ([19a312d](https://github.com/turbopuffer/turbopuffer-go/commit/19a312d298a7825f45fc2f5e4734e6ae8b8fecaa))
* pretty print error messages ([#107](https://github.com/turbopuffer/turbopuffer-go/issues/107)) ([d112cc1](https://github.com/turbopuffer/turbopuffer-go/commit/d112cc11e286f7c430a12e1428e8fb102d2d87c0))
* prevent duplicate ? in query params ([6905729](https://github.com/turbopuffer/turbopuffer-go/commit/6905729c901a4888f5e4c940452b4fcef638e322))
* **README:** improve headline example ([#40](https://github.com/turbopuffer/turbopuffer-go/issues/40)) ([86b1e43](https://github.com/turbopuffer/turbopuffer-go/commit/86b1e434ee16f4737e771b0a2082f71174b1b8d2))
* **README:** remove beta warning ([e323ed5](https://github.com/turbopuffer/turbopuffer-go/commit/e323ed59716f7c9f856e01aa67acce96b20c7f24))
* **README:** update for improved API ([1214410](https://github.com/turbopuffer/turbopuffer-go/commit/12144102ba431e3446f8a8735088276227bf213a))
* remove 200 error code for hint_cache_warm API call (always 202 now) ([27456b8](https://github.com/turbopuffer/turbopuffer-go/commit/27456b8fd37d7df9645ac4127dafc67b68ed6831))
* remove a heap allocation from Client.Namespace ([#97](https://github.com/turbopuffer/turbopuffer-go/issues/97)) ([e563eb2](https://github.com/turbopuffer/turbopuffer-go/commit/e563eb2b3ffb40c8f8d93d31fcfd30b3d43b325f))
* remove null from release please manifest ([f48d2dd](https://github.com/turbopuffer/turbopuffer-go/commit/f48d2dd56a9e07ea5959ed7bb7f79bba80c7b9fd))
* rename param to avoid collision ([4ce5239](https://github.com/turbopuffer/turbopuffer-go/commit/4ce5239899ce4b5b70ea1b522f36bdcdecec0705))
* report incidence of client sharing in user-agent header ([#89](https://github.com/turbopuffer/turbopuffer-go/issues/89)) ([052ddea](https://github.com/turbopuffer/turbopuffer-go/commit/052ddea4aad948e1529610f06c0a7cb1358267b6))
* skip usage tests that don't work with Prism ([02a8479](https://github.com/turbopuffer/turbopuffer-go/commit/02a847971368f58ac4c2176ec69a3685c58c9931))
* strict types for new QueryParam type ([f313206](https://github.com/turbopuffer/turbopuffer-go/commit/f313206d02bf76cc167f4e822c3a1c422bb273b5))
* strict types for write conditions ([#38](https://github.com/turbopuffer/turbopuffer-go/issues/38)) ([054f318](https://github.com/turbopuffer/turbopuffer-go/commit/054f31825a189331f21c31b1dd84afa56b10144d))
* **test:** adjust test for float32 vectors ([7fefd1b](https://github.com/turbopuffer/turbopuffer-go/commit/7fefd1b6a416b5d0106cf806a1f0ec92bfc556fb))
* **tests:** ensure every test uses its own namespace ([f058eb1](https://github.com/turbopuffer/turbopuffer-go/commit/f058eb108f3b60b50885cb511985b3d949cfb705))
* **tests:** update for new JSON decoding behavior ([bca72ed](https://github.com/turbopuffer/turbopuffer-go/commit/bca72ed331b2a67498357d10bd15626b30470b08))
* update README warning for beta release ([508cd84](https://github.com/turbopuffer/turbopuffer-go/commit/508cd844f7b10e1a18c2fc5c5a4db878b9f0daf9))
* update tests for new metadata endpoint ([d339b92](https://github.com/turbopuffer/turbopuffer-go/commit/d339b9205fdd46560787d31fce8a85483615cccc))
* update tests for new update_metadata endpoint ([6a3910f](https://github.com/turbopuffer/turbopuffer-go/commit/6a3910fecdc3da6194071350923e9629880405fe))
* use release please annotations on more places ([7d6af35](https://github.com/turbopuffer/turbopuffer-go/commit/7d6af35fc5ee3e2ecd733b3439b9c58a937a0a24))
* use slices.Concat instead of sometimes modifying r.Options ([e99f3f0](https://github.com/turbopuffer/turbopuffer-go/commit/e99f3f0b39d08b0d042fdcdc68e6d607578d0623))
* use strict types for ExplainQuery ([#62](https://github.com/turbopuffer/turbopuffer-go/issues/62)) ([b2d2d6c](https://github.com/turbopuffer/turbopuffer-go/commit/b2d2d6c67f8d8993fe91d3c605aba14f0b96341c))
* v2 module compatibility ([8d1f803](https://github.com/turbopuffer/turbopuffer-go/commit/8d1f803565910d5bde8960c6dba20af1320bfce1))


### Chores

* add float64 to valid types for RegisterFieldValidator ([8030579](https://github.com/turbopuffer/turbopuffer-go/commit/8030579395df13bd951fedf80511fbf8d4bcd41f))
* bump apigen and regenerate ([a642525](https://github.com/turbopuffer/turbopuffer-go/commit/a642525905b6efb290c52a70484077129feba173))
* bump gjson version ([5092b43](https://github.com/turbopuffer/turbopuffer-go/commit/5092b43158163b99fb80b06e472f31626ac862f8))
* bump minimum go version to 1.22 ([5f446a8](https://github.com/turbopuffer/turbopuffer-go/commit/5f446a8be537c1cb2f576100faec9910aad3ee31))
* **ci:** add build step ([71657ba](https://github.com/turbopuffer/turbopuffer-go/commit/71657bab846006551ca64cd0c11c93b0371d9b7a))
* **ci:** only run for pushes and fork pull requests ([15433df](https://github.com/turbopuffer/turbopuffer-go/commit/15433dff0a50eb902b8db0b74aab8a6a10272438))
* **ci:** skip lint on metadata-only changes ([7d6bab2](https://github.com/turbopuffer/turbopuffer-go/commit/7d6bab26cd9237185560663e0179475b6ac670bd))
* **ci:** skip uploading artifacts on stainless-internal branches ([c79e741](https://github.com/turbopuffer/turbopuffer-go/commit/c79e741899ed6f8dc154cec95b83e3da59622b44))
* **ci:** support opting out of skipping builds on metadata-only commits ([4e298e7](https://github.com/turbopuffer/turbopuffer-go/commit/4e298e74dfabe6be22769815daee56490aa40e90))
* **client:** fix multipart serialisation of Default() fields ([18d402f](https://github.com/turbopuffer/turbopuffer-go/commit/18d402f993bbcd62ccc3acf6397521ee8b7583ba))
* codegen updates ([7d54294](https://github.com/turbopuffer/turbopuffer-go/commit/7d54294bc407bb9aee79592e34db2fdd12b2a3d0))
* **codegen:** bump apigen ([#55](https://github.com/turbopuffer/turbopuffer-go/issues/55)) ([bb9df3f](https://github.com/turbopuffer/turbopuffer-go/commit/bb9df3fe45d4c202354425110f8a3a4a124a0140))
* configure new SDK language ([5038bca](https://github.com/turbopuffer/turbopuffer-go/commit/5038bcaa6f23721f829632d7aaa884aee57832a6))
* do not install brew dependencies in ./scripts/bootstrap by default ([3ad3674](https://github.com/turbopuffer/turbopuffer-go/commit/3ad36744784cf5488ccb7aacae45a6e4b8339b47))
* elide duplicate aliases ([89926ae](https://github.com/turbopuffer/turbopuffer-go/commit/89926ae6a710495adabf2bd1b4495f71a80d75e5))
* fix automatic type regeneration ([632aedd](https://github.com/turbopuffer/turbopuffer-go/commit/632aedd7c48620c80ebc4fb4a21bebc3668b5cdf))
* fix documentation of null map ([17d3136](https://github.com/turbopuffer/turbopuffer-go/commit/17d3136424fc6843f74346e489d76c029e464e8f))
* fix empty interfaces ([c7d93ac](https://github.com/turbopuffer/turbopuffer-go/commit/c7d93ac82d7158f0955c5caa6a9a7a8126ced32b))
* fix type in test ([2add877](https://github.com/turbopuffer/turbopuffer-go/commit/2add877e2d0461333797b8b41cb1c398de426fce))
* fix typesg ([#73](https://github.com/turbopuffer/turbopuffer-go/issues/73)) ([a314531](https://github.com/turbopuffer/turbopuffer-go/commit/a3145310b816bf797d3bb0be7f90fd02489ced52))
* format ([3d5c559](https://github.com/turbopuffer/turbopuffer-go/commit/3d5c5597480dde051f121b2e8075aacbf4ea7d23))
* **gen:** reapply custom code ([b6d999e](https://github.com/turbopuffer/turbopuffer-go/commit/b6d999e289ff7219cf0b50999ff2427e801bb0cd))
* **internal:** codegen related update ([3012f28](https://github.com/turbopuffer/turbopuffer-go/commit/3012f2899bb162443ae7c248c403e2e4648b5adc))
* **internal:** codegen related update ([e2bb4ec](https://github.com/turbopuffer/turbopuffer-go/commit/e2bb4ec0dd1d6e61a1220b290c503934deb67dc2))
* **internal:** codegen related update ([92133ce](https://github.com/turbopuffer/turbopuffer-go/commit/92133ce3cb827191e627fadb261faf00c517e1ec))
* **internal:** codegen related update ([c26d21a](https://github.com/turbopuffer/turbopuffer-go/commit/c26d21af7ed99655dfc0b384b7ee021ac0623759))
* **internal:** codegen related update ([95f9dd8](https://github.com/turbopuffer/turbopuffer-go/commit/95f9dd8bd652fc97382f9d54d4e9ece6765083d5))
* **internal:** fix .stats.yml ([c638e0b](https://github.com/turbopuffer/turbopuffer-go/commit/c638e0b4b5bfa0413aed92895ba7a522be646eae))
* **internal:** fix lint script for tests ([c094793](https://github.com/turbopuffer/turbopuffer-go/commit/c094793a10a65568a992c58a4f755d39af7e2496))
* **internal:** grammar fix (it's -&gt; its) ([6763e84](https://github.com/turbopuffer/turbopuffer-go/commit/6763e8426d292a0a019cd23400bd64b486c690c4))
* **internal:** minor cleanup ([cfa5968](https://github.com/turbopuffer/turbopuffer-go/commit/cfa596835926c954d4c83fed55405221558e7e49))
* **internal:** more robust bootstrap script ([6f4283a](https://github.com/turbopuffer/turbopuffer-go/commit/6f4283aa5c27d6149e06265934673e60bd741a61))
* **internal:** move custom custom `json` tags to `api` ([d7f091b](https://github.com/turbopuffer/turbopuffer-go/commit/d7f091b5890c7d1575e3241390ac15ce08b84aa3))
* **internal:** remove mock server code ([70cf362](https://github.com/turbopuffer/turbopuffer-go/commit/70cf362669a3bb1348daabe1c27114abd16f9455))
* **internal:** support default value struct tag ([dce7928](https://github.com/turbopuffer/turbopuffer-go/commit/dce79285e1e26f17be6607607622d8263e3f61ac))
* **internal:** tweak CI branches ([ce2f8b3](https://github.com/turbopuffer/turbopuffer-go/commit/ce2f8b31e12d1b2cc3fd789054a1627abc659a33))
* **internal:** update `actions/checkout` version ([f602215](https://github.com/turbopuffer/turbopuffer-go/commit/f6022155f644200099932984e4a7412eddb9f294))
* **internal:** update comment in script ([1e37611](https://github.com/turbopuffer/turbopuffer-go/commit/1e3761124a62e7f169f9a1e177d50dda6302d787))
* **internal:** update gitignore ([37e70be](https://github.com/turbopuffer/turbopuffer-go/commit/37e70be9063eae3de4558079a6e5f435524ba158))
* **internal:** use different example values for some enums ([a79c7e8](https://github.com/turbopuffer/turbopuffer-go/commit/a79c7e87bc3b7dcc7858dd090298474c4cb4a0d8))
* **internal:** use explicit returns ([2a8f8b8](https://github.com/turbopuffer/turbopuffer-go/commit/2a8f8b816787f765f53b2e93c6081b9083dfbedd))
* **internal:** use explicit returns in more places ([4b0524d](https://github.com/turbopuffer/turbopuffer-go/commit/4b0524dd3a2d1e194b575028ed62d810efc4f270))
* **internal:** version bump ([e9eee78](https://github.com/turbopuffer/turbopuffer-go/commit/e9eee7878716dadcb6c84975202be57a43d60eef))
* lint tests ([96b8f94](https://github.com/turbopuffer/turbopuffer-go/commit/96b8f94cce13f86df77f39027adcc64649e33832))
* lint tests in subpackages ([a771524](https://github.com/turbopuffer/turbopuffer-go/commit/a771524e5c92a859ac5c84158224e69acfc3f3d1))
* make go mod tidy continue on error ([2a4ab7f](https://github.com/turbopuffer/turbopuffer-go/commit/2a4ab7f733ea69a8cfd2299c1f7b23115f149ad2))
* regenerate custom types ([40f4dca](https://github.com/turbopuffer/turbopuffer-go/commit/40f4dcac0a1b429333a80da614df5df1ffed618f))
* regenerate custom types ([f467c1d](https://github.com/turbopuffer/turbopuffer-go/commit/f467c1d01fd685e663f4f05a5c16ed943674eb5f))
* remove unnecessary error check for url parsing ([055f7d3](https://github.com/turbopuffer/turbopuffer-go/commit/055f7d3f1dbbf51a2f5c9182e3204a52adbe6951))
* **stainless:** unstick codegen ([8274970](https://github.com/turbopuffer/turbopuffer-go/commit/82749708b040548e3f06d255b1f4c96c232fc8c8))
* update @stainless-api/prism-cli to v5.15.0 ([8335db1](https://github.com/turbopuffer/turbopuffer-go/commit/8335db1798a0de6ebe43e88de046bfd9fd84524a))
* update docs for api:"required" ([da158b6](https://github.com/turbopuffer/turbopuffer-go/commit/da158b6be46528b53466201be86383e2d84d8cf8))
* update mock server docs ([0d78267](https://github.com/turbopuffer/turbopuffer-go/commit/0d78267e5d2a87d631bf532b001ceee69d2a0464))
* update more docs for 1.22 ([a5f3ad2](https://github.com/turbopuffer/turbopuffer-go/commit/a5f3ad21f689438ba10a4dcd7b51303cfc3826a1))


### Documentation

* add cross-region copy_from_namespace to write API docs ([6c53240](https://github.com/turbopuffer/turbopuffer-go/commit/6c53240fc3f8d1e293812df9dbadd86ae25c11d3))
* hint_cache_warm also update header and openapi ([86c822b](https://github.com/turbopuffer/turbopuffer-go/commit/86c822b928bd6f797bea7833e0397948c0d0cdec))

## 2.0.0-alpha.1 (2026-04-29)

Full Changelog: [v1.21.0...v2.0.0-alpha.1](https://github.com/turbopuffer/turbopuffer-go/compare/v1.21.0...v2.0.0-alpha.1)

### Features

* Add sparse vectors to the spec. ([58682b3](https://github.com/turbopuffer/turbopuffer-go/commit/58682b3a40d4663667d6570d5a1f5a678f7bcfd1))
* Enable Stainless normalizer v2 while keeping Query inlined ([a1db134](https://github.com/turbopuffer/turbopuffer-go/commit/a1db134f535862fb56d4709d79e294abf2e7d4ba))
* Enable Stainless normalizer v2 while keeping Query inlined ([dd7f4f5](https://github.com/turbopuffer/turbopuffer-go/commit/dd7f4f5582fc40975d6bdff4bf406e2fe0bda1d6))
* **go:** add default http client with timeout ([7c99444](https://github.com/turbopuffer/turbopuffer-go/commit/7c99444fd9c4e40e8a716528ab5e9c86941b722d))
* Require `fuzzy: true` in the schema to use the `Fuzzy` filter ([904a15f](https://github.com/turbopuffer/turbopuffer-go/commit/904a15fc7c200f5cb6b74e40a4d4d3f25ae4df87))
* set up branchFrom and copyFrom sugar in the OpenAPI spec ([52262b4](https://github.com/turbopuffer/turbopuffer-go/commit/52262b4b325ba3006603b3172523c76ba220ba3a))
* spec: flatten branchFrom/copyFrom sugar endpoint params ([c4a8a79](https://github.com/turbopuffer/turbopuffer-go/commit/c4a8a792239a7183c08d70fc2f1a479befe77007))
* spec: update OpenAPI spec for v2 encryption API ([836e30d](https://github.com/turbopuffer/turbopuffer-go/commit/836e30d92d2881575616f699b3c2e17c8c6769db))
* support setting headers via env ([bbe8e4f](https://github.com/turbopuffer/turbopuffer-go/commit/bbe8e4f09612f18e5a5ca176188f61bda4e46319))


### Bug Fixes

* pretty print error messages ([#107](https://github.com/turbopuffer/turbopuffer-go/issues/107)) ([d112cc1](https://github.com/turbopuffer/turbopuffer-go/commit/d112cc11e286f7c430a12e1428e8fb102d2d87c0))


### Chores

* bump apigen and regenerate ([a642525](https://github.com/turbopuffer/turbopuffer-go/commit/a642525905b6efb290c52a70484077129feba173))
* **internal:** more robust bootstrap script ([6f4283a](https://github.com/turbopuffer/turbopuffer-go/commit/6f4283aa5c27d6149e06265934673e60bd741a61))
* regenerate custom types ([40f4dca](https://github.com/turbopuffer/turbopuffer-go/commit/40f4dcac0a1b429333a80da614df5df1ffed618f))

## 1.21.0 (2026-04-06)

Full Changelog: [v1.20.0...v1.21.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.20.0...v1.21.0)

### Features

* add update_metadata method to configure pinning ([de29524](https://github.com/turbopuffer/turbopuffer-go/commit/de29524032d37453a9af5702f7ae5612112d3e70))
* docs for glob/regex acceleration ([a880af7](https://github.com/turbopuffer/turbopuffer-go/commit/a880af7389e23fb1957504d98401849d4c08c6ad))
* **internal:** support comma format in multipart form encoding ([2a346e6](https://github.com/turbopuffer/turbopuffer-go/commit/2a346e6d6b926e595296480da957074be70b7a34))
* Multiple Vector Columns Docs ([8f99599](https://github.com/turbopuffer/turbopuffer-go/commit/8f99599eaf1a9018d2bbcbe820fdd85e10a1c603))
* spec: include pinning status in /metadata response ([5ab6247](https://github.com/turbopuffer/turbopuffer-go/commit/5ab624789d4c5a22164e8de7bc1ec1af4b1fd114))


### Bug Fixes

* fix issue with unmarshaling in some cases ([b8ebac3](https://github.com/turbopuffer/turbopuffer-go/commit/b8ebac3f7926a044cc31d3ade972409497abb5d5))
* prevent duplicate ? in query params ([7580361](https://github.com/turbopuffer/turbopuffer-go/commit/758036150ec2eca94ce1e8327632324fe4d91bd8))


### Chores

* **ci:** skip lint on metadata-only changes ([3c881e6](https://github.com/turbopuffer/turbopuffer-go/commit/3c881e67100d3959cc9bd7de5505f452601bb793))
* **ci:** support opting out of skipping builds on metadata-only commits ([2c2f689](https://github.com/turbopuffer/turbopuffer-go/commit/2c2f6894548a26e5219da3723055fe0c23998ff0))
* **client:** fix multipart serialisation of Default() fields ([f761f6c](https://github.com/turbopuffer/turbopuffer-go/commit/f761f6cca5136185651cf095d159f52701afbd1e))
* **internal:** support default value struct tag ([439a79d](https://github.com/turbopuffer/turbopuffer-go/commit/439a79d67c3dfae0fcc073230f6bbbb028a1d365))
* **internal:** update gitignore ([1d43b62](https://github.com/turbopuffer/turbopuffer-go/commit/1d43b62c9271e1cc537e785136eba2ed2931fc16))
* remove unnecessary error check for url parsing ([dd0abf0](https://github.com/turbopuffer/turbopuffer-go/commit/dd0abf038b4f9cd5d4ade1007ec497a6ddb2fb6a))
* update docs for api:"required" ([0922544](https://github.com/turbopuffer/turbopuffer-go/commit/09225440cf918718677821bf248dc0fe98efadc5))

## 1.20.0 (2026-03-20)

Full Changelog: [v1.19.0...v1.20.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.19.0...v1.20.0)

### Features

* spec: disable CopyFromNamespace syntax sugar in Go ([ebf43a1](https://github.com/turbopuffer/turbopuffer-go/commit/ebf43a187a6a58cd74da5cfaf2a20757f9a625a7))

## 1.19.0 (2026-03-19)

Full Changelog: [v1.18.0...v1.19.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.18.0...v1.19.0)

### Features

* Response performance object ([7331c39](https://github.com/turbopuffer/turbopuffer-go/commit/7331c394370919c6a22a9608845ca2dac9f7f4bb))
* tpuf-engine: add verbose form of branch_from_namespace ([3e7d6fc](https://github.com/turbopuffer/turbopuffer-go/commit/3e7d6fca031d394a4b8969fbf4d3461d9edb53ca))
* tpuf: add branching to openapi spec ([9d35d3c](https://github.com/turbopuffer/turbopuffer-go/commit/9d35d3cf53808069e8e2edb843b40f726c9a40c7))

## 1.18.0 (2026-03-16)

Full Changelog: [v1.17.0...v1.18.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.17.0...v1.18.0)

### Features

* retract versions before v1.12.2 ([#102](https://github.com/turbopuffer/turbopuffer-go/issues/102)) ([b7a7581](https://github.com/turbopuffer/turbopuffer-go/commit/b7a75817fb612d9607c892073985a4f4e434524a))


### Chores

* **internal:** tweak CI branches ([5d65824](https://github.com/turbopuffer/turbopuffer-go/commit/5d65824fef8050d63d2964c0d46f2564816652aa))

## 1.17.0 (2026-03-16)

Full Changelog: [v1.16.0...v1.17.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.16.0...v1.17.0)

### Features

* skip generating Go integer alias for Limit ([d500cb3](https://github.com/turbopuffer/turbopuffer-go/commit/d500cb3feef7a7ca3d6d7da62f84fc248420ac13))

## 1.16.0 (2026-03-15)

Full Changelog: [v1.15.0...v1.16.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.15.0...v1.16.0)

### Features

* spec: host MCP on stainless ([c650252](https://github.com/turbopuffer/turbopuffer-go/commit/c650252a12316c755752ce2c45b72d6a069d5679))


### Bug Fixes

* fix request delays for retrying to be more respectful of high requested delays ([dcc553b](https://github.com/turbopuffer/turbopuffer-go/commit/dcc553b286b8aee9197a96761bdb1574b0da5eb6))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([4e23c73](https://github.com/turbopuffer/turbopuffer-go/commit/4e23c736f4f7b034d879c5d7956925eeeeecc276))
* **internal:** minor cleanup ([1cd5f3a](https://github.com/turbopuffer/turbopuffer-go/commit/1cd5f3a48b9711e96ec79eed6f03600e1a7f003d))
* **internal:** use explicit returns ([5fb8758](https://github.com/turbopuffer/turbopuffer-go/commit/5fb8758b84daafe3676b16a660bcd4b34619b278))
* **internal:** use explicit returns in more places ([c58fa2f](https://github.com/turbopuffer/turbopuffer-go/commit/c58fa2f190cd829f6d765083bab22881ea3fbedd))

## 1.15.0 (2026-03-03)

Full Changelog: [v1.14.1...v1.15.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.14.1...v1.15.0)

### Features

* Remove `queries` from recall endpoint ([8c8c382](https://github.com/turbopuffer/turbopuffer-go/commit/8c8c3828946e47c9a4560e354699979987b78d09))


### Chores

* **ci:** add build step ([9d83d22](https://github.com/turbopuffer/turbopuffer-go/commit/9d83d222d0ef6841bc0bad1d19f03a4c60b67f79))
* **internal:** move custom custom `json` tags to `api` ([7c92765](https://github.com/turbopuffer/turbopuffer-go/commit/7c92765c4d48dcfc8fd2617c7b4a1d17e0b9698c))

## 1.14.1 (2026-02-21)

Full Changelog: [v1.14.0...v1.14.1](https://github.com/turbopuffer/turbopuffer-go/compare/v1.14.0...v1.14.1)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([c8d8de4](https://github.com/turbopuffer/turbopuffer-go/commit/c8d8de4932a7daf9da4f15469c5279363d5c3ddf))
* **internal:** skip tests that depend on mock server ([7212688](https://github.com/turbopuffer/turbopuffer-go/commit/7212688d30966588be71754f379b7d3de7634a3f))


### Chores

* **internal:** remove mock server code ([47e426a](https://github.com/turbopuffer/turbopuffer-go/commit/47e426acbf8f16e2e610fa8df6edd2b97b651947))
* update mock server docs ([afdbcca](https://github.com/turbopuffer/turbopuffer-go/commit/afdbcca2137baa710f6302280f9d656dd0d86a8a))

## 1.14.0 (2026-02-17)

Full Changelog: [v1.12.2...v1.14.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.12.2...v1.14.0)

### Features

* Add ranking-by-attribute to the spec. ([e186f45](https://github.com/turbopuffer/turbopuffer-go/commit/e186f45c8503aff4d901d08b7b7922400fbc60fd))


### Bug Fixes

* improve tracking of client sharing when client is copied by value ([#96](https://github.com/turbopuffer/turbopuffer-go/issues/96)) ([fddc1bb](https://github.com/turbopuffer/turbopuffer-go/commit/fddc1bb9e3a550b3336cf6523e36b58ee8f0072d))
* remove a heap allocation from Client.Namespace ([#97](https://github.com/turbopuffer/turbopuffer-go/issues/97)) ([142c213](https://github.com/turbopuffer/turbopuffer-go/commit/142c213ee807cd49789f84902f907fbdae7127b0))

## 1.12.2 (2026-02-12)

Full Changelog: [v1.12.1...v1.12.2](https://github.com/turbopuffer/turbopuffer-go/compare/v1.12.1...v1.12.2)

### Bug Fixes

* report incidence of client sharing in user-agent header ([#89](https://github.com/turbopuffer/turbopuffer-go/issues/89)) ([88504be](https://github.com/turbopuffer/turbopuffer-go/commit/88504bef482c08d04f974ffc01c367e5e855b1ae))

## 1.12.1 (2026-02-12)

Full Changelog: [v1.12.0...v1.12.1](https://github.com/turbopuffer/turbopuffer-go/compare/v1.12.0...v1.12.1)

### Bug Fixes

* data race in Client.Namespace caused by slice aliasing ([#88](https://github.com/turbopuffer/turbopuffer-go/issues/88)) ([3acb3ef](https://github.com/turbopuffer/turbopuffer-go/commit/3acb3ef9f2051fbacc05b9cd6808ef122611dabb))
* **encoder:** correctly serialize NullStruct ([4023b4e](https://github.com/turbopuffer/turbopuffer-go/commit/4023b4ee28276f56696289b9481863a2bdf8bc44))

## 1.12.0 (2026-02-08)

Full Changelog: [v1.11.0...v1.12.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.11.0...v1.12.0)

### Features

* add support for `limit` query parameter ([6cedf33](https://github.com/turbopuffer/turbopuffer-go/commit/6cedf33af1cbfe9d57b4bf8b3b3db8822c67a303))
* spec: clean up limit codegen ([bc7c81e](https://github.com/turbopuffer/turbopuffer-go/commit/bc7c81e2c71c554ba6ff0717af8316b2b4fd6cb5))

## 1.11.0 (2026-02-03)

Full Changelog: [v1.10.0...v1.11.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.10.0...v1.11.0)

### Features

* ContainsAnyToken last_as_prefix docs ([b83360c](https://github.com/turbopuffer/turbopuffer-go/commit/b83360cfe97a271cf887b30213b6f59c56674c27))


### Chores

* regenerate custom types ([a07c028](https://github.com/turbopuffer/turbopuffer-go/commit/a07c028235f3b1c0dde4c1d826a5f1dd6944c226))

## 1.10.0 (2026-02-01)

Full Changelog: [v1.9.2...v1.10.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.9.2...v1.10.0)

### Features

* [TPUF-1097] tpuf: return failed IDs for conditional writes ([7d3ca6d](https://github.com/turbopuffer/turbopuffer-go/commit/7d3ca6da169ab367c60700ffc9712a09d8d6eb0c))
* add ContainsAnyToken and RankByKnn ([150b93d](https://github.com/turbopuffer/turbopuffer-go/commit/150b93dd0116b017927de8b1b2b95bb83b7b577b))
* Apigen integration ([ea11d4f](https://github.com/turbopuffer/turbopuffer-go/commit/ea11d4f58dac413b4a1d41f8e70a1f3a7d05bf75))
* **client:** add a convenient param.SetJSON helper ([90b3e5c](https://github.com/turbopuffer/turbopuffer-go/commit/90b3e5c94b664be6058c82a279cac46a4d691b1e))
* disable response compression by default ([#82](https://github.com/turbopuffer/turbopuffer-go/issues/82)) ([12c7229](https://github.com/turbopuffer/turbopuffer-go/commit/12c722941503807c9b46365648019a8c900190aa))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([4c9d77e](https://github.com/turbopuffer/turbopuffer-go/commit/4c9d77ed4b1e83b9e77b8d0b9e342ff1cc093205))
* **docs:** fix mcp installation instructions for remote servers ([e5ca2f9](https://github.com/turbopuffer/turbopuffer-go/commit/e5ca2f92c135e4c3ddb75bea6c44d499cd2cb2d0))
* fix for namespace collisions with client and resource test methods ([94df670](https://github.com/turbopuffer/turbopuffer-go/commit/94df670bc96fdf0151d8a2e1ba5cf2158d86746f))


### Chores

* configure new SDK language ([53f0bff](https://github.com/turbopuffer/turbopuffer-go/commit/53f0bff8ce220fb9b5492ea138bef3e7641511f6))
* fix automatic type regeneration ([209063b](https://github.com/turbopuffer/turbopuffer-go/commit/209063bc9b350984968528609521ac69328fd110))
* format ([d3e6135](https://github.com/turbopuffer/turbopuffer-go/commit/d3e613512aff4fa38a3b1fefe42d0108ecba2a9e))
* **internal:** update `actions/checkout` version ([ad49beb](https://github.com/turbopuffer/turbopuffer-go/commit/ad49beb49ec5ff7d1ec5b543b38e98ce377ee0e2))

## 1.9.2 (2026-01-09)

Full Changelog: [v1.9.1...v1.9.2](https://github.com/turbopuffer/turbopuffer-go/compare/v1.9.1...v1.9.2)

### Chores

* **internal:** codegen related update ([a379399](https://github.com/turbopuffer/turbopuffer-go/commit/a3793990b62d7057f8ee3b868fb9e8ce419488cf))
* **internal:** use different example values for some enums ([49e22d7](https://github.com/turbopuffer/turbopuffer-go/commit/49e22d738215613abc9ae650ff99618385ff40cc))

## 1.9.1 (2025-12-20)

Full Changelog: [v1.9.0...v1.9.1](https://github.com/turbopuffer/turbopuffer-go/compare/v1.9.0...v1.9.1)

### Bug Fixes

* skip usage tests that don't work with Prism ([8d223e2](https://github.com/turbopuffer/turbopuffer-go/commit/8d223e28f870012a43c913ffbb55351acdb79f65))


### Chores

* add float64 to valid types for RegisterFieldValidator ([544caf7](https://github.com/turbopuffer/turbopuffer-go/commit/544caf777bf49f49ecb12182779630b06e870c9b))

## 1.9.0 (2025-12-11)

Full Changelog: [v1.8.1...v1.9.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.8.1...v1.9.0)

### Features

* add word_v3 to the spec ([d44424c](https://github.com/turbopuffer/turbopuffer-go/commit/d44424cb44c00235c55135d038a717477ac7051d))
* **encoder:** support bracket encoding form-data object members ([f19598b](https://github.com/turbopuffer/turbopuffer-go/commit/f19598b75e51326c02463994789f7c225f91e791))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([57db013](https://github.com/turbopuffer/turbopuffer-go/commit/57db013f2dc4104ead835246fe079f4746b32f4b))
* rename param to avoid collision ([76aaed3](https://github.com/turbopuffer/turbopuffer-go/commit/76aaed355e99d06af78e2bd937e0f6e1b165803e))


### Chores

* codegen updates ([46fb461](https://github.com/turbopuffer/turbopuffer-go/commit/46fb461b63ac9123550505ef9e037cca68588287))
* elide duplicate aliases ([280bb3d](https://github.com/turbopuffer/turbopuffer-go/commit/280bb3d701917eeafdea4a3fecf9d4c8616fbd72))
* **internal:** codegen related update ([381360f](https://github.com/turbopuffer/turbopuffer-go/commit/381360f0d7a4fc3283b0a2411f9ae0728d6a09d0))

## 1.8.1 (2025-12-02)

Full Changelog: [v1.8.0...v1.8.1](https://github.com/turbopuffer/turbopuffer-go/compare/v1.8.0...v1.8.1)

### Documentation

* add cross-region copy_from_namespace to write API docs ([e6a308a](https://github.com/turbopuffer/turbopuffer-go/commit/e6a308ad645b2e9c3aa0b5a7af61367e5eb3b5b1))

## 1.8.0 (2025-12-02)

Full Changelog: [v1.7.0...v1.8.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.7.0...v1.8.0)

### Features

* sdks: add &lt;patch|delete&gt;_by_filter_allow_partial options ([31d7603](https://github.com/turbopuffer/turbopuffer-go/commit/31d7603c751f661623d879e8045cd10df93239c8))


### Chores

* fix empty interfaces ([00873f0](https://github.com/turbopuffer/turbopuffer-go/commit/00873f0f842927c55c20bd9078c713002d92f7bf))

## 1.7.0 (2025-11-25)

Full Changelog: [v1.6.0...v1.7.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.6.0...v1.7.0)

### Features

* site: add ascii_folding to docs and SDKs ([e194058](https://github.com/turbopuffer/turbopuffer-go/commit/e194058219658aca34e3429049b89dc85a44f8d7))

## 1.6.0 (2025-11-17)

Full Changelog: [v1.5.0...v1.6.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.5.0...v1.6.0)

### Features

* Add vector attribute schema to metadata endpoint ([0982efc](https://github.com/turbopuffer/turbopuffer-go/commit/0982efcce52f5673a5f6568f784c316ff7ed7ba9))
* Allow for a CMEK key to be specified in copy_from_namespace ([399c479](https://github.com/turbopuffer/turbopuffer-go/commit/399c4798d27deb3ba77e4ebb54a1cb7fcadfceac))
* Make `type` required on `AttributeSchemaConfig` ([c14032b](https://github.com/turbopuffer/turbopuffer-go/commit/c14032bf2bf52d2ba755e3563618163e401e45ff))
* openapi: Fix stainless warnings ([9abc646](https://github.com/turbopuffer/turbopuffer-go/commit/9abc646dd97cb3d1d3d0f099a7bc2a02a5abd1f2))
* openapi: name variants of `NamespaceMetadata.index` ([1f60a82](https://github.com/turbopuffer/turbopuffer-go/commit/1f60a82df018bac74a562af87a102577cb328faa))
* spec: add support for cross-org CFN to SDKs ([0793e8f](https://github.com/turbopuffer/turbopuffer-go/commit/0793e8f4c757d18be19374a39978e9b4f32a2287))


### Chores

* bump gjson version ([a19c027](https://github.com/turbopuffer/turbopuffer-go/commit/a19c027b6e58e9e128d91137ab1011fa5db12cd6))
* **internal:** grammar fix (it's -&gt; its) ([efdbc92](https://github.com/turbopuffer/turbopuffer-go/commit/efdbc92fe118ae50a5e07ae17e8563ffbfcc9b73))

## 1.5.0 (2025-10-22)

Full Changelog: [v1.3.0...v1.5.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.3.0...v1.5.0)

### Features

* Metadata endpoint updates (e.g. to track indexing progress) ([e1f2486](https://github.com/turbopuffer/turbopuffer-go/commit/e1f2486e467665a5592cfb0570a170e7be98c368))
* required for patch_by_filter :facepalm: ([edce63a](https://github.com/turbopuffer/turbopuffer-go/commit/edce63af457851832b6565f7576b2b9a79c2eef6))
* stainless: add patch_by_filter ([b351165](https://github.com/turbopuffer/turbopuffer-go/commit/b351165da997cbd251f991ed7861fb65030f20af))


### Chores

* fix typesg ([#73](https://github.com/turbopuffer/turbopuffer-go/issues/73)) ([bc51479](https://github.com/turbopuffer/turbopuffer-go/commit/bc5147925059c78871af4d936b5ae2106e107813))

## 1.3.0 (2025-10-14)

Full Changelog: [v1.2.0...v1.3.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.2.0...v1.3.0)

### Features

* Add float, []float and []bool to the list of valid types in the OpenAPI spec. ([9027106](https://github.com/turbopuffer/turbopuffer-go/commit/902710638c3872b5a0a561a8aeee50ecf87996ff))
* Promote disable_backpressure to first-class Write property ([8a479d4](https://github.com/turbopuffer/turbopuffer-go/commit/8a479d433bd807207ffb493957d1accae8fddbc9))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([cb3e41c](https://github.com/turbopuffer/turbopuffer-go/commit/cb3e41c2baab248cc59e19c6459f7bb0452c58c0))


### Documentation

* hint_cache_warm also update header and openapi ([63c6e68](https://github.com/turbopuffer/turbopuffer-go/commit/63c6e68f4763411a8a4b8efa9c7d1179cccce678))

## 1.2.0 (2025-09-19)

Full Changelog: [v1.1.0...v1.2.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.1.0...v1.2.0)

### Features

* add WithParams variant to BM25 and ContainsAllTokens ([5adf1cd](https://github.com/turbopuffer/turbopuffer-go/commit/5adf1cd11ff1568b9ca7ae35861e6d3099ef9fe0))


### Bug Fixes

* use slices.Concat instead of sometimes modifying r.Options ([bfebe89](https://github.com/turbopuffer/turbopuffer-go/commit/bfebe899beea3a4379f21f88890c4bc8947097de))


### Chores

* bump minimum go version to 1.22 ([c917377](https://github.com/turbopuffer/turbopuffer-go/commit/c917377212159488594d098817a9f004146552d0))
* do not install brew dependencies in ./scripts/bootstrap by default ([895d0d5](https://github.com/turbopuffer/turbopuffer-go/commit/895d0d586761b234839fe2de9924f6d0a627475d))
* update more docs for 1.22 ([9a0c9c0](https://github.com/turbopuffer/turbopuffer-go/commit/9a0c9c0a2d1a8b69d47db39959741df0b9b7a64e))

## 1.1.0 (2025-09-11)

Full Changelog: [v1.0.0...v1.1.0](https://github.com/turbopuffer/turbopuffer-go/compare/v1.0.0...v1.1.0)

### Features

* spec: add dedicated type for AggregationGroup response ([e92ddfe](https://github.com/turbopuffer/turbopuffer-go/commit/e92ddfe7a2bbd10e3fb5cbfea96e0e26a97ec5d2))
* tpuf: add include_ground_truth option to recall endpoint ([a5838ca](https://github.com/turbopuffer/turbopuffer-go/commit/a5838cae384ce595ca82799b9ffb60df068f47cd))


### Bug Fixes

* close body before retrying ([a624a99](https://github.com/turbopuffer/turbopuffer-go/commit/a624a991f40c4255972987720f8e2642ed812b62))
* **internal:** unmarshal correctly when there are multiple discriminators ([1c21ca3](https://github.com/turbopuffer/turbopuffer-go/commit/1c21ca38bb846f39ce58f14e7a4719dbc6bb0bec))
* remove null from release please manifest ([f8ae2f8](https://github.com/turbopuffer/turbopuffer-go/commit/f8ae2f82d103855fb6f69122f545bf32efbaf65b))
* use release please annotations on more places ([a6aa93a](https://github.com/turbopuffer/turbopuffer-go/commit/a6aa93a6d758a2094df1293265df99929beea2d6))

## 1.0.0 (2025-08-28)

Full Changelog: [v0.1.19...v1.0.0](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.19...v1.0.0)

### Features

* Make word_v2 the default FTS tokenizer ([1f6c1c1](https://github.com/turbopuffer/turbopuffer-go/commit/1f6c1c1a8024661195584a099aa8f11dea452530))

## 0.1.19 (2025-08-18)

Full Changelog: [v0.1.18...v0.1.19](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.18...v0.1.19)

### Bug Fixes

* **api:** add support for `group_by` query parameter ([2a06092](https://github.com/turbopuffer/turbopuffer-go/commit/2a06092a3ccb1583c29e0c2636e08ee94f4e3711))

## 0.1.18 (2025-08-12)

Full Changelog: [v0.1.17...v0.1.18](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.17...v0.1.18)

### Bug Fixes

* remove 200 error code for hint_cache_warm API call (always 202 now) ([7662222](https://github.com/turbopuffer/turbopuffer-go/commit/7662222902ae0dc0afecda3067d957e02074e2d2))


### Chores

* **internal:** codegen related update ([7d9d5d6](https://github.com/turbopuffer/turbopuffer-go/commit/7d9d5d63f5f2e48db466866cfc1f3b2ba7189b3b))

## 0.1.17 (2025-08-11)

Full Changelog: [v0.1.16...v0.1.17](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.16...v0.1.17)

### Chores

* **internal:** fix .stats.yml ([06a7a35](https://github.com/turbopuffer/turbopuffer-go/commit/06a7a359126a5785d121a2e4cf59d44b4b72e345))
* **internal:** update comment in script ([8b16e89](https://github.com/turbopuffer/turbopuffer-go/commit/8b16e896be5da2661238d7b2a1758ac8ee1307aa))
* update @stainless-api/prism-cli to v5.15.0 ([26ce790](https://github.com/turbopuffer/turbopuffer-go/commit/26ce790273caff1f42d3afc0792e389312a0de7b))

## 0.1.16 (2025-08-08)

Full Changelog: [v0.1.15...v0.1.16](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.15...v0.1.16)

### Bug Fixes

* use strict types for ExplainQuery ([#62](https://github.com/turbopuffer/turbopuffer-go/issues/62)) ([aa986b8](https://github.com/turbopuffer/turbopuffer-go/commit/aa986b86aebb72c4b63a972bb5fab4d51835beda))

## 0.1.15 (2025-08-08)

Full Changelog: [v0.1.14...v0.1.15](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.14...v0.1.15)

### Features

* **client:** support optional json html escaping ([da5ef74](https://github.com/turbopuffer/turbopuffer-go/commit/da5ef74148df430c91776ea258b47b5844c8fcaf))


### Bug Fixes

* add FullTextSearchConfig.MaxTokenLength ([fbf6e7b](https://github.com/turbopuffer/turbopuffer-go/commit/fbf6e7b662facaa4886b8dc1515510a9b0a52036))
* **api:** api update ([4cf2036](https://github.com/turbopuffer/turbopuffer-go/commit/4cf2036e0fdb46bb89feb087e81715a4c0bdd3de))

## 0.1.14 (2025-07-29)

Full Changelog: [v0.1.13...v0.1.14](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.13...v0.1.14)

### Bug Fixes

* **api:** api update ([8b19a44](https://github.com/turbopuffer/turbopuffer-go/commit/8b19a443fbed878093a10eb849d6953b33b59718))

## 0.1.13 (2025-07-29)

Full Changelog: [v0.1.12...v0.1.13](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.12...v0.1.13)

### Bug Fixes

* **api:** api update ([e26712e](https://github.com/turbopuffer/turbopuffer-go/commit/e26712e0e1ed9863cb272fb3b51665b83efcb0d2))
* **api:** api update ([15f4193](https://github.com/turbopuffer/turbopuffer-go/commit/15f4193cdb352d9b86a8b9f8525f6113310da68e))

## 0.1.12 (2025-07-29)

Full Changelog: [v0.1.11...v0.1.12](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.11...v0.1.12)

### Bug Fixes

* add support for regex filter ([2e53da3](https://github.com/turbopuffer/turbopuffer-go/commit/2e53da3f4f0ddefd09a925574c7a1217108ee684))
* **api:** api update ([1adc270](https://github.com/turbopuffer/turbopuffer-go/commit/1adc270865616d75ba3364c77ddbf2a3eab37e95))
* **api:** api update ([90a773b](https://github.com/turbopuffer/turbopuffer-go/commit/90a773b4843f8794b9c5abeb75a779298d7e9c4d))

## 0.1.11 (2025-07-28)

Full Changelog: [v0.1.10...v0.1.11](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.10...v0.1.11)

### Bug Fixes

* correct some bugs in example code ([#56](https://github.com/turbopuffer/turbopuffer-go/issues/56)) ([9779814](https://github.com/turbopuffer/turbopuffer-go/commit/9779814e770a7e4e700122716995b224245c7fc1))
* more precise types for filters that take arrays ([64ca1b6](https://github.com/turbopuffer/turbopuffer-go/commit/64ca1b6057dc8760ea10da6958f1c7d0de4501f3))


### Chores

* **codegen:** bump apigen ([#55](https://github.com/turbopuffer/turbopuffer-go/issues/55)) ([053ecde](https://github.com/turbopuffer/turbopuffer-go/commit/053ecde20cf2e5d3d5dedd62e4e12a85a8cac51b))

## 0.1.10 (2025-07-23)

Full Changelog: [v0.1.9...v0.1.10](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.9...v0.1.10)

### Bug Fixes

* **api:** api update ([5cb2ff1](https://github.com/turbopuffer/turbopuffer-go/commit/5cb2ff15cf21caa5fd1ee782d1f8c1126f0d7b74))

## 0.1.9 (2025-07-22)

Full Changelog: [v0.1.8...v0.1.9](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.8...v0.1.9)

### Bug Fixes

* **client:** process custom base url ahead of time ([ad1a20a](https://github.com/turbopuffer/turbopuffer-go/commit/ad1a20af43504bbe558ff60afc94da8f527f4c61))

## 0.1.8 (2025-07-18)

Full Changelog: [v0.1.7...v0.1.8](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.7...v0.1.8)

### Chores

* lint tests in subpackages ([c06b027](https://github.com/turbopuffer/turbopuffer-go/commit/c06b0276c389e37a19ed07b8c4be94ca58448866))

## 0.1.7 (2025-07-10)

Full Changelog: [v0.1.6...v0.1.7](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.6...v0.1.7)

### Bug Fixes

* **api:** api update ([756d1b1](https://github.com/turbopuffer/turbopuffer-go/commit/756d1b1f66a9a25504b474eec6649c372d557ee2))
* update tests for new metadata endpoint ([aa44590](https://github.com/turbopuffer/turbopuffer-go/commit/aa4459028acdad58f070a3f256f720f3cf4bb89e))

## 0.1.6 (2025-07-09)

Full Changelog: [v0.1.5...v0.1.6](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.5...v0.1.6)

### Bug Fixes

* **api:** api update ([fb0c2b5](https://github.com/turbopuffer/turbopuffer-go/commit/fb0c2b53a16809d9b67838bb3ced20a201248840))
* don't require region unless baseUrl contains {region} placeholder ([#47](https://github.com/turbopuffer/turbopuffer-go/issues/47)) ([5509f4d](https://github.com/turbopuffer/turbopuffer-go/commit/5509f4d162bf6f22b66a250ae3d5240bdb6cb114))


### Chores

* **internal:** fix lint script for tests ([b3da884](https://github.com/turbopuffer/turbopuffer-go/commit/b3da88451e1870e79844550528617792bf9d469a))
* lint tests ([9f7396a](https://github.com/turbopuffer/turbopuffer-go/commit/9f7396aa76d388a9db8f13561365a9d664413dd0))

## 0.1.5 (2025-07-07)

Full Changelog: [v0.1.4...v0.1.5](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.4...v0.1.5)

### Bug Fixes

* **api:** add support for [Not]Contains[Any] operators ([#43](https://github.com/turbopuffer/turbopuffer-go/issues/43)) ([78e84e0](https://github.com/turbopuffer/turbopuffer-go/commit/78e84e020eb8245e5188d92c71b853b85c60c380))

## 0.1.4 (2025-07-01)

Full Changelog: [v0.1.3...v0.1.4](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.3...v0.1.4)

### Features

* **client:** add escape hatch for null slice & maps ([c89845a](https://github.com/turbopuffer/turbopuffer-go/commit/c89845a9de85e1911b8141f45609559b027b2ea2))


### Bug Fixes

* add missing Filter prefix to ContainsAllTokensArray ([#39](https://github.com/turbopuffer/turbopuffer-go/issues/39)) ([d419e8a](https://github.com/turbopuffer/turbopuffer-go/commit/d419e8a0163096a918b2aaae13562a85d043347b))
* add support for `$ref_new` expressions ([#41](https://github.com/turbopuffer/turbopuffer-go/issues/41)) ([53a4120](https://github.com/turbopuffer/turbopuffer-go/commit/53a4120a248a872337192b3071bae8149073874f))
* **api:** api update ([7466bc1](https://github.com/turbopuffer/turbopuffer-go/commit/7466bc1747bd584a4650d6624bfee1c4bb630cc8))
* **api:** api update ([51942f8](https://github.com/turbopuffer/turbopuffer-go/commit/51942f8e1d88781fd529891a52ce56c61e6c2845))
* correct name of stainless bot ([#42](https://github.com/turbopuffer/turbopuffer-go/issues/42)) ([7ab0e6f](https://github.com/turbopuffer/turbopuffer-go/commit/7ab0e6f4cd4872029b5996968d3a0283544e6a8e))
* don't try to deserialize as json when ResponseBodyInto is []byte ([9f62897](https://github.com/turbopuffer/turbopuffer-go/commit/9f62897af2ba1bd3a1f5f2406c099dcb849ef6ad))
* **pagination:** check if page data is empty in GetNextPage ([1bea120](https://github.com/turbopuffer/turbopuffer-go/commit/1bea12027038f1576052758de6b5cf0a255ac2fc))
* **README:** improve headline example ([#40](https://github.com/turbopuffer/turbopuffer-go/issues/40)) ([17fed10](https://github.com/turbopuffer/turbopuffer-go/commit/17fed108319f629b5e843407e5eb3c7cf94c02ab))
* strict types for write conditions ([#38](https://github.com/turbopuffer/turbopuffer-go/issues/38)) ([ff23fc1](https://github.com/turbopuffer/turbopuffer-go/commit/ff23fc161de06e5cdd21f7343447c1e1c41ff1f5))


### Chores

* **ci:** only run for pushes and fork pull requests ([648ed30](https://github.com/turbopuffer/turbopuffer-go/commit/648ed3042097a0325d9e572ef6780fbaaa3d2bd9))
* fix documentation of null map ([4588f72](https://github.com/turbopuffer/turbopuffer-go/commit/4588f725efffa1eaacd98c23503d1f38a5822265))

## 0.1.3 (2025-06-20)

Full Changelog: [v0.1.2...v0.1.3](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.2...v0.1.3)

### Features

* **api:** introduce dedicated Query model ([531f2df](https://github.com/turbopuffer/turbopuffer-go/commit/531f2df3e602bd785ad7eea5e6e0e6c20b9c72bc))


### Bug Fixes

* strict types for new QueryParam type ([f313206](https://github.com/turbopuffer/turbopuffer-go/commit/f313206d02bf76cc167f4e822c3a1c422bb273b5))

## 0.1.2 (2025-06-19)

Full Changelog: [v0.1.1...v0.1.2](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.1...v0.1.2)

### Features

* **client:** add debug log helper ([d3568a7](https://github.com/turbopuffer/turbopuffer-go/commit/d3568a7fa15a19b59a5be613d5bc5e9bce40808e))


### Bug Fixes

* strict types for async query and multi_query APIs ([#35](https://github.com/turbopuffer/turbopuffer-go/issues/35)) ([1d114bf](https://github.com/turbopuffer/turbopuffer-go/commit/1d114bf90f58a3ea9f06d345ad712d6063ea6a18))


### Chores

* **ci:** enable for pull requests ([00aa5a8](https://github.com/turbopuffer/turbopuffer-go/commit/00aa5a8ade868e49a156bbaa6fca0e0f595e8db6))

## 0.1.1 (2025-06-11)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0...v0.1.1)

### Features
* **api:** Add support for new multi-query api
  
### Bug Fixes

* **README:** align documentation links with other SDKs ([246d40d](https://github.com/turbopuffer/turbopuffer-go/commit/246d40db961fb8d03c2484dfd9f89c581f8015a6))

## 0.1.0 (2025-06-09)

Full Changelog: [v0.1.0-beta.1...v0.1.0](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-beta.1...v0.1.0)

### Bug Fixes

* **README:** remove beta warning ([0c4f8d1](https://github.com/turbopuffer/turbopuffer-go/commit/0c4f8d1dbf784e06924f64e3475984cf3b30e7cb))
* **tests:** ensure every test uses its own namespace ([354e598](https://github.com/turbopuffer/turbopuffer-go/commit/354e5981d76d95b1379d390dae1da0f56b25dd7f))
* update README warning for beta release ([508cd84](https://github.com/turbopuffer/turbopuffer-go/commit/508cd844f7b10e1a18c2fc5c5a4db878b9f0daf9))

## 0.1.0-beta.1 (2025-06-05)

Full Changelog: [v0.1.0-alpha.25...v0.1.0-beta.1](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.25...v0.1.0-beta.1)

### Chores

* prepare for beta release ([1e98a55](https://github.com/turbopuffer/turbopuffer-go/commit/1e98a55425c11033c1397d8872e70f99561a63e5))

## 0.1.0-alpha.25 (2025-06-05)

Full Changelog: [v0.1.0-alpha.24...v0.1.0-alpha.25](https://github.com/turbopuffer/turbopuffer-go/compare/v0.1.0-alpha.24...v0.1.0-alpha.25)

### Bug Fixes

* **api:** fix type in improved JSON number type ([75d62e6](https://github.com/turbopuffer/turbopuffer-go/commit/75d62e6cc5a5a66a10cf5030af10e0a7bf8b68ff))
* **api:** improved JSON number type ([600c85d](https://github.com/turbopuffer/turbopuffer-go/commit/600c85ddedfd173524d83bb3b96fd6a42735c884))
* **README:** update for improved API ([0ce58ac](https://github.com/turbopuffer/turbopuffer-go/commit/0ce58ac14e587262ce4590e5b73fa013820c7457))

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
