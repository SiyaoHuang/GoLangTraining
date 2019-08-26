import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';
import { connect, sendMsg } from "./api";

class App extends Component {
  constructor(props){
    super(props)
    connect()
  }
  
  send(){
    console.log("hello")
    sendMsg("hello")
  }
  render() {
    return (

      <div className="App">
        <body>
          The content of the document......
        </body>
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
