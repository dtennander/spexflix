import React, {Component} from 'react';
import Api from "../api";
import jwtDecode from 'jwt-decode'
import MovieList from "./MovieList";

const headerStyle = {
    margin: "20px",
};

class HomeView extends Component{

    constructor(props) {
        super(props);
        this.state = {
            user: {},
            tokenData: {},
        }
    }

    componentDidMount() {
        try {
            const tokenData = jwtDecode(this.props.token);
            this.setState({tokenDate: tokenData});
            Api.GetUser(tokenData.id, this.props.token)
                .then(user => this.setState({user: user}))
                .catch(error => console.log("Could not load user. ", error))
        } catch(error) {
            console.log(error)
        }
    }

    render() {
        return (
            <div style={headerStyle}>
                <h1>Välkommen {this.state.user.name}!</h1>
                <MovieList token={this.props.token}/>
            </div>
        )
    }
}

export default HomeView