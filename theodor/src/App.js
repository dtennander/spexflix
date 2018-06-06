import React, { Component } from 'react';
import Komplimanger from "./Komplimanger";
const AppStyle = {
    textAlign: "center",
};

const HeaderStyle = {
    backgroundColor: "#f1eb00",
    height: "50px",
    padding: "10px",
    color: "#061599",
    boxShadow: "0 6px 8px 0 rgba(0,0,0,0.12), 0 9px 25px 0 rgba(0,0,0,0.09)",
};

const TitleStyle = {
    fontSize: "2em",
    marginTop: "5px",
    marginBottom: "0px",
    fontFamily: "'Playfair Display', serif",
};

const Header = (props) => {
    return (<header style={HeaderStyle}>
        <h1 style={TitleStyle}>Stort Grattis!</h1>
    </header>
    );
};



class App extends Component {
  render() {
    return (
      <div style={AppStyle}>
        <Header/>
        <Komplimanger/>
      </div>
    );
  }
}

export default App;
