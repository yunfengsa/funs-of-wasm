
/**
 * wasm加载
 */
let go_instance;
let isOnGoing = false;
// 实例过程初始化
const loadWasm = (resolve, reject) => {
  isOnGoing = true;
  WebAssembly.instantiateStreaming(fetch("http://127.0.0.1:3000/wasm"), go.importObject).then(async (result) => {
    go.run(result.instance)
    console.log(result.instance)
    go_instance = result.instance
    setTimeout(() => {
      resolve && resolve()
      isOnGoing = false;
    }, 0);
  }).catch(err => {
    isOnGoing = false;
    reject && reject()
  });
};
loadWasm();

// 获得工具函数
const getUtils = () => {
  return new Promise((resolve, reject) => {
    if (go_instance) {
      resolve();
    } else {
      if (isOnGoing) {
        const delay = (mResolve, mReject) => setTimeout(() => {
          if (isOnGoing) {
            delay(mResolve, mReject)
          } else {
            if (go_instance) {
              mResolve()
            } else {
              mReject()
            }
          }
        }, 100);
        delay(resolve, reject)
      } else {
        loadWasm(resolve, reject);
      }
    }
  })
}

export default getUtils
