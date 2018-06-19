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
        this.updateInput = this.updateInput.bind(this);
        this.state = {
            inputValue: ''
        }
    }

    getStyle() {
        return this.props.style || inputStyle;
    }

    render() {
        return (
            <input
                style={this.getStyle()}
                type={this.props.type}
                name={this.props.name}
                value={this.state.inputValue}
                placeholder={this.props.placeholder}
                onKeyDownCapture={this.updateInput}
                onChange={this.updateInput}/>
        )
    }

    updateInput(event) {
        if (event.key === "Enter" && typeof this.props.onEnter !== 'undefined') {
            this.props.onEnter();
            return
        }
        if (event.key === "Escape" && typeof this.props.onEscape !== 'undefined') {
            this.props.onEscape();
            return
        }
        this.setState({
            inputValue: event.target.value
        })
    }

    getInput() {
        return this.state.inputValue
    }

    clearField() {
        this.setState({inputValue: ""})
    }
}

export default InputField