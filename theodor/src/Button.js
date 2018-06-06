import React, { Component } from 'react';
import Interactive from "react-interactive";


const buttonStyle = {
    color: "#061599",
    backgroundColor: "#f1eb00",
    borderRadius: "15px",
    border: "none",
    fontSize: "0.8em",
    fontWeight: "bold",
    margin: "5px",
    height: "50px",
    boxShadow: "0 6px 8px 0 rgba(0,0,0,0.12), 0 9px 25px 0 rgba(0,0,0,0.09)",
};
const buttonActiveStyle = {
    transitionDuration: "0.1s",
    boxShadow: "0 0px 0px 0 rgba(0,0,0,0.12), 0 0px 0px 0 rgba(0,0,0,0.09)",
};

class Button extends Component {
    render() {
        return (
            <Interactive
                as="button"
                type="button"
                style={{...buttonStyle, ...this.props.style}}
                active={buttonActiveStyle}
                onClick={this.props.onClick}>
                    {this.props.text}
            </Interactive>
        )
    }
}

export default Button