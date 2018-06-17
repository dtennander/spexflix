import React, { Component } from 'react';
import Header from '../components/Header'
import {Redirect} from "react-router";
import MovieView from "../components/MovieView";

class Year extends Component {

    constructor(props) {
        super(props);
        const jwtToken = getJwtToken();
        this.state = {
            jwtToken: jwtToken,
            isLoggedIn : jwtToken != null
        };
        console.log(props);
        this.onSuccessfulLogout = this.onSuccessfulLogout.bind(this);
        this.onSuccessfulLogIn = this.onSuccessfulLogIn.bind(this);
    }

    componentDidMount() {
        document.title = document.title + " | " + this.props.match.params.year;
    }

    render() {
        if (!this.state.isLoggedIn) {
            return <Redirect to="/"/>;
        } else {
            return (
                <div>
                    <Header
                        isLoggedIn={true}
                        onSuccessfulLogout={this.onSuccessfulLogout}/>
                    <MovieView year={this.props.match.params.year} token={this.state.jwtToken}/>
                </div>
            );
        }
    }

    onSuccessfulLogout() {
        localStorage.removeItem("jwtToken");
        this.setState({
            isLoggedIn: false,
            jwtToken: null,
        });
    }

    onSuccessfulLogIn(token) {
        localStorage.setItem("jwtToken", token);
        this.setState({
            isLoggedIn: true,
            jwtToken: token,
        })
    }
}

function getJwtToken() {
    return localStorage.getItem("jwtToken")
}

export default Year;
