import React, { Component } from 'react';
import Interactive from "react-interactive";


const buttonStyle = {
    color: "#71141b",
    backgroundColor: "#f1eb00",
    borderRadius: "3px",
    border: "none",
    fontSize: "0.8em",
    fontWeight: "bold",
    transitionDuration: "0.4s",
    margin: "5px",
    height: "25px",
};

const buttonHoverStyle = {
    boxShadow: "0 6px 8px 0 rgba(0,0,0,0.12), 0 9px 25px 0 rgba(0,0,0,0.09)",
};

const buttonActiveStyle = {
    transitionDuration: "0.1s"
};

class Button extends Component {
    render() {
        return (
            <Interactive
                as="button"
                type="button"
                style={{...buttonStyle, ...this.props.style}}
                hover={buttonHoverStyle}
                active={buttonActiveStyle}
                onClick={this.props.onClick}>
                    {this.props.text}
            </Interactive>
        )
    }
}

export default Button