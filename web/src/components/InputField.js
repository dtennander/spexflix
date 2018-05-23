import React, { Component } from 'react';

const inputStyle = {
    width: '100%',
    padding: "6px 10px",
    margin: "1px 0",
    display: "inline-block",
    border: "1px solid #ccc",
    borderRadius: "4px",
    boxSizing: "border-box",
};

/**
 * Input field for forms.
 * @param type
 * @param name
 * @param placeholder
 */
class InputField  extends Component {
    constructor(props) {
        super(props);
        this.state = {
            inputValue: ''
        }
    }

    getStyle() {
        return {...inputStyle, ...this.props.style}
    }
    render() {
        return (
            <input
                style={this.getStyle()}
                type={this.props.type}
                name={this.props.name}
                value={this.state.inputValue}
                placeholder={this.props.placeholder}
                onChange={(evt) => this.updateInput(evt.target.value)}/>
        )
    }

    updateInput(newInput) {
        this.setState({
            inputValue: newInput
        })
    }

    getInput() {
        return this.state.inputValue
    }
}

export default InputField