import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header'
import ChatHistory from './components/ChatHistory/ChatHistory'
import ChatInput from './components/ChatInput/ChatInput'

class App extends Component {
  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }
  
  constructor(props) {
    super(props)
    this.state = {
        chatHistory: []
    }
}
  
  send(event){
    if(event.keyCode === 13) {
      sendMsg(event.target.value)
      event.target.value = ""
    }
    // console.log("hello")
    // sendMsg("hello")
  }
  render() {
    return (
      <div className="App">
        <Header />
        <ChatInput send={this.send}/>
        <button onClick={this.send}>Hit</button>
        <ChatHistory chatHistory={this.state.chatHistory} />
      </div>
    );
  }
}

export default App;
