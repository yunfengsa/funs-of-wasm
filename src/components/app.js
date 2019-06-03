import React from 'react'

export default class App extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      isLoading: true
    }
  }

  componentDidMount() {
		WebAssembly.instantiateStreaming(fetch("http://127.0.0.1:3000/wasm"), go.importObject).then(async (result) => {
      go.run(result.instance)
      this.setState({ isLoading: false })
      console.log(checkGameState(1,2))
    });
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
