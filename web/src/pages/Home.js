import React, { Component } from 'react';
import Header from '../components/Header'
import HomeView from '../components/HomeView';
import CreateAccount from "../components/CreateAccount";

class Home extends Component {

    constructor(props) {
        super(props);
        const jwtToken = getJwtToken();
        this.state = {
            jwtToken: jwtToken,
            isLoggedIn : jwtToken != null
        };
        this.onSuccessfulLogout = this.onSuccessfulLogout.bind(this);
        this.onSuccessfulLogIn = this.onSuccessfulLogIn.bind(this);
    }

    render() {
        if (!this.state.isLoggedIn) {
            return (<div>
                <Header
                    isLoggedIn={false}
                    onSuccessfulLogIn={this.onSuccessfulLogIn}/>
                <CreateAccount/>
            </div>);
        } else {
            return (
                <div>
                    <Header
                        isLoggedIn={true}
                        onSuccessfulLogout={this.onSuccessfulLogout}/>
                    <HomeView token={this.state.jwtToken}/>
                </div>
            );
        }
    }

    onSuccessfulLogout() {
        localStorage.removeItem("jwtToken");
        this.setState({isLoggedIn: false});
    }

    onSuccessfulLogIn(token) {
        localStorage.setItem("jwtToken", token);
        this.setState({isLoggedIn: true})
    }
}

function getJwtToken() {
    return localStorage.getItem("jwtToken")
}

export default Home;
