# Changelog

## [0.1.0](https://github.com/loom-go/sig/compare/v0.1.0...v0.1.0) (2026-04-07)


### ⚠ BREAKING CHANGES

* new public api

### Features

* add owner.OnDispose ([6e6ccc2](https://github.com/loom-go/sig/commit/6e6ccc24248411e32d3ac20bb4cf7ee3af1d3171))
* contexts ([a6d47b9](https://github.com/loom-go/sig/commit/a6d47b95ba1dd1c944f5c16f706a08c8b2b241bb))
* owner.run error return ([b711818](https://github.com/loom-go/sig/commit/b7118182682affa814fab816e5b1b70190026cbf))
* render effects & on settled ([6df905a](https://github.com/loom-go/sig/commit/6df905a6eb3610c68d3821a31caf542ba43d56f7))
* **signal:** custom predicate ([1725101](https://github.com/loom-go/sig/commit/17251018b48a6ee3bd1d88cf056c0ef97cd4f966))


### Bug Fixes

* avoid uneeded propagations ([94438dd](https://github.com/loom-go/sig/commit/94438dd24c5cf091541a5018434089843545e403))
* batch from another runtime ([a827198](https://github.com/loom-go/sig/commit/a8271984b06c3e31266ddbf093f986bbf59b2fd2))
* children using wrong owner in goroutine ([4a04aaa](https://github.com/loom-go/sig/commit/4a04aaa7dbf2e009e67289941a97d8f965ff940c))
* disposing computed with same deps ([bd0ef55](https://github.com/loom-go/sig/commit/bd0ef55fecd479611670b3e95394430b6f42e63b))
* effect panics not cought ([a094447](https://github.com/loom-go/sig/commit/a094447f6e37399e3c904667c6905e05a809a1e1))
* effects not diposed in owner ([754cf2a](https://github.com/loom-go/sig/commit/754cf2ac5849349ee8f8dc0c41397ba2960bace6))
* interleaving contexts ([337d392](https://github.com/loom-go/sig/commit/337d3928be14070307e79fa87def953e123727df))
* nested owners tracking ([adc498f](https://github.com/loom-go/sig/commit/adc498f647193b0b4ad0f5d13ecf1298a9f543af))
* **owner:** r/w races ([#1](https://github.com/loom-go/sig/issues/1)) ([b5f1542](https://github.com/loom-go/sig/commit/b5f15421cf6c23a8ce04b349802c3aa95da22404))
* racondition & goroutine isolation ([e491448](https://github.com/loom-go/sig/commit/e491448b1e3adcc3a26da7177e5f9dacb9fb161a))
* read error on nil zero values ([8da7a28](https://github.com/loom-go/sig/commit/8da7a2838a55207c79479779c589fcf20fe243d0))
* rw race condition ([ac5439d](https://github.com/loom-go/sig/commit/ac5439d2fd6366dad3ff72a6c501aba5dc7e1938))
* **wasm:** goroutine isolation ([592667e](https://github.com/loom-go/sig/commit/592667e4e4460347fe6d7c8fe3e704c9b3b54d19))


### Miscellaneous Chores

* release 0.1.0 ([022f5c7](https://github.com/loom-go/sig/commit/022f5c7e0880f458d2a00da4c2181d5801675c54))


### Code Refactoring

* new public api ([1b702a8](https://github.com/loom-go/sig/commit/1b702a8dc0c6c6a633d0e06ba7147c341288087e))
