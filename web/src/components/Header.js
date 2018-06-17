import React, { Component } from 'react';
import logo from '../images/luva.svg'
import LoginForm from "./LoginForm";
import Button from "./Button";
import {Link} from "react-router-dom";

const mainStyle = {
    textAlign: "left",
};

const headerStyle = {
    backgroundColor: "#962528",
    height: "40px",
    padding: "5px",
    color: "#f1eb00",
    boxShadow: "0 6px 8px 0 rgba(0,0,0,0.12), 0 9px 25px 0 rgba(0,0,0,0.09)",
};

const titleStyle = {
    fontSize: "2em",
    margin: "0px",
    marginLeft: "8px",
    marginTop: "5px"
};

const logoStyle = {
    marginRight: "5px",
    height: "30px",
};

/**
 * @param isLoggedIn
 * @param onSuccessfullLogIn
 */
class Header extends Component {

    constructor(props) {
        super(props);
        this.state = {width: 0, height: 0 };
        this.updateWindowDimensions = this.updateWindowDimensions.bind(this);
    }

    componentDidMount() {
        this.updateWindowDimensions();
        window.addEventListener('resize', this.updateWindowDimensions);
    }

    componentWillUnmount() {
        window.removeEventListener('resize', this.updateWindowDimensions);
    }

    updateWindowDimensions() {
        this.setState({ width: window.innerWidth, height: window.innerHeight });
    }

    getCenterIfSmallElse(e) {
        let result = {};
        if (this.isCentered()) {
            result = {
                textAlign: "center",
                color: "#f1eb00",
                textDecoration: "none",
            };
        } else {
            result = {
                float: e,
                color: "#f1eb00",
                textDecoration: "none",
            };
        }
        return result;
    }

    isCentered() {
        return this.state.width < 494;
    }

    getHeaderStyle() {
        if (this.isCentered()) {
            return {...headerStyle, ...{height: "80px"}}
        }  else {
            return headerStyle
        }
    }

    render() {
        let LeftDiv;
        if (!this.props.isLoggedIn) {
            LeftDiv = this.getLoginDiv();
        } else {
            LeftDiv = this.getUserDiv();
        }
        return (
            <div style={mainStyle}>
                <header style={this.getHeaderStyle()}>
                    <Link style={this.getCenterIfSmallElse("left")} to="/">
                        <h1 style={titleStyle}>
                            <img style={logoStyle} src={logo} alt="logo" />
                            Spexflix
                        </h1>
                    </Link>
                    <LeftDiv/>
                </header>
            </div>
        )
    }

    getLoginDiv() {
        return () => (
            <div style={this.getCenterIfSmallElse("right")}>
                <LoginForm onSuccessfulLogIn={this.props.onSuccessfulLogIn}/>
            </div>
        )
    }

    getUserDiv() {
        return () => (
            <div style={this.getCenterIfSmallElse("right")}>
                <Button onClick={this.props.onSuccessfulLogout} text="Logga ut"/>
            </div>
        )
    }
}

export default Header