一个实现通过wasm实现函数集的解法


base go 1.12

#### install and run
* cd wasm-server && go mod tidy && build.sh // 生成wasm 并启动server
* npm install && npm run dev
* http://localhost:8081/ // 观察控制台，能看到计算结果由wasm产生
