# Changelog

## 0.0.1 (2025-04-08)


### Features

* add location ID to routes ([f881707](https://github.com/majori/wrc-laptimer/commit/f881707cbe476c127c548fd90f4e1ab0bbe132cf))
* handle system signals properly (ctrl-c) ([93d1f15](https://github.com/majori/wrc-laptimer/commit/93d1f15521d31bfc36129b12ea347e28cbbab357))
* link telemetry to a session ([a3e33ae](https://github.com/majori/wrc-laptimer/commit/a3e33ae08a1d2e59bf12327063546b8444ee1190))
* save default telemetry packet to DuckDB ([0bbd7d9](https://github.com/majori/wrc-laptimer/commit/0bbd7d98d8f87051a72107d888182a22842085ed))
* serve static web files ([37612da](https://github.com/majori/wrc-laptimer/commit/37612da3b74c9d51dc96221b83d4cef93e15daac))
* split telemetry to separate channels ([3ae0582](https://github.com/majori/wrc-laptimer/commit/3ae05825cab53ab44858b6d085220f995ea4cdbb))
* store sessions to database ([a936bcd](https://github.com/majori/wrc-laptimer/commit/a936bcd155017993e4304d1a7ad82e0bfae91c88))
* **web:** initialize alpine project ([#2](https://github.com/majori/wrc-laptimer/issues/2)) ([ce59fd8](https://github.com/majori/wrc-laptimer/commit/ce59fd84f77eee9d8a5acf642eff17303fd162f9))


### Bug Fixes

* close db properly ([b85e7d0](https://github.com/majori/wrc-laptimer/commit/b85e7d02569bac83cd62c1a84ce526ddd2c43087))
* ignore telemetry if no active session ([c53a460](https://github.com/majori/wrc-laptimer/commit/c53a46023f337974c6fc9fb65bd3844a6a4c2cea))
* insert login if user exists ([8d0aa0f](https://github.com/majori/wrc-laptimer/commit/8d0aa0ffd594ef9a6527a7db5be9ca06f0e37825))
* LastInsertId() does not work on DuckDB ([b797596](https://github.com/majori/wrc-laptimer/commit/b7975969a62fc6fa1e5a8f0f9b03d1ee1b4ef086))
* remove foreign key constraint ([5561f50](https://github.com/majori/wrc-laptimer/commit/5561f50804b1c9090c38a1ff699ac57a16949aa8))
* set session ID correctly ([e13316f](https://github.com/majori/wrc-laptimer/commit/e13316f0ae8fb4efd126bc26b181673ab09f4103))
