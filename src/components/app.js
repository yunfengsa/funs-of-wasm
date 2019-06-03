import React from 'react'
import getUtils from './utils';

export default class App extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      isLoading: true
    }
  }

  componentDidMount() {
    getUtils().then(() => {
      console.log('add', wasm_func1(1, 2))
      console.log('minus', wasm_func2(1, 3))
      const now = Date.now();
      console.log('sqrt', wasm_func3())
      console.log('diff time', Date.now() - now)
      const jsNow = Date.now();
      console.log('jssqrt', this.sqrt())
      console.log('js diff time', Date.now() - jsNow)
    }, () => {
      console.log('加载失败')
    }).catch(err => {
      console.log(err)
    });
  }

  sqrt() {
    let a = 700000000;
    for (let i = 0; i < 100; i++){
      a = Math.sqrt(a);
    }
    return a;
  }


  render() {
    return (
      <div style={{ height: '100%', display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
        { 
          this.state.isLoading ? 
            <div>123</div> : <div>456</div>
        }
      </div>
    )
  }
}
