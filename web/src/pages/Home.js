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

    componentDidMount() {
        document.title = "Spexflix"
    }

    render() {
        return (
            <div>
                <Header isLoggedIn={this.state.isLoggedIn}
                        onSuccessfulLogout={this.onSuccessfulLogout}
                        onSuccessfulLogIn={this.onSuccessfulLogIn}/>
                {this.state.isLoggedIn
                    ? <HomeView token={this.state.jwtToken}/>
                    : <CreateAccount/> }
            </div>);
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

export default Home;
